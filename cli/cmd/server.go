package cmd

import (
	"context"
	"fmt"
	"os"
	"syscall"

	protov1 "github.com/aidtechnology/affinityctl/proto/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.bryk.io/x/cli"
	"go.bryk.io/x/net/loader"
	"go.bryk.io/x/net/rpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var helper *loader.Helper

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start a gateway instance",
	Aliases: []string{"gateway", "gw", "node"},
	RunE:    runServer,
}

func init() {
	segments := []string{
		loader.SegmentRPC,
		loader.SegmentObservability,
	}
	helper = loader.New()
	if err := cli.SetupCommandParams(serverCmd, helper.Params(segments...)); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(serverCmd)
}

func runServer(cmd *cobra.Command, args []string) error {
	// Load configuration settings.
	// Manually set some default values.
	viper.Set("rpc.network_interface", rpc.NetworkInterfaceAll)
	viper.Set("rpc.http_gateway.enabled", true)
	if err := viper.Unmarshal(helper.Data); err != nil {
		return err
	}

	// Create HTTP gateway
	gwOpts := helper.HTTPGateway()
	gw, _ := rpc.NewHTTPGateway(gwOpts...)

	// Create server instance
	serverOptions := append(helper.ServerRPC(),
		rpc.WithHTTPGateway(gw),
		rpc.WithServiceProvider(&sampleServiceHandler{}),
	)
	server, _ := rpc.NewServer(serverOptions...)

	// Start server
	ready := make(chan bool)
	go func() {
		_ = server.Start(ready)
	}()
	<-ready

	// wait for system signals
	fmt.Printf("waiting for requests on port: %d\n", helper.Data.RPC.Port)
	<-cli.SignalsHandler([]os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	})

	// Close server
	fmt.Println("closing server")
	_ = server.Stop(true)
	return nil
}

// Sample service handler
// For real applications a service handler implementation is
// usually located on a separate package and handles all the
// business logic and functional requirements.

type sampleServiceHandler struct {
	protov1.UnimplementedGatewayAPIServer
}

func (s *sampleServiceHandler) Ping(_ context.Context, _ *emptypb.Empty) (*protov1.PingResponse, error) {
	return &protov1.PingResponse{Ok: true}, nil
}

func (s *sampleServiceHandler) ServerSetup(server *grpc.Server) {
	protov1.RegisterGatewayAPIServer(server, s)
}

func (s *sampleServiceHandler) GatewaySetup() rpc.GatewayRegister {
	return protov1.RegisterGatewayAPIHandlerFromEndpoint
}
