package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.bryk.io/x/cli"
)

var vcListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Authenticate and list stored credentials",
	Aliases: []string{"ls"},
	Example: "affinityctl vc list -u [DID] -p [PIN]",
	RunE: func(_ *cobra.Command, _ []string) error {
		// Get user parameters
		did := viper.GetString("list.did")
		if strings.TrimSpace(did) == "" {
			return errors.New("--did is required")
		}
		pin := viper.GetString("list.pin")
		if strings.TrimSpace(pin) == "" {
			p, err := cli.ReadSecure("Enter PIN: ")
			if err != nil {
				return err
			}
			pin = string(p)
		}

		// Get SDK client
		sdk, err := sdkClient()
		if err != nil {
			return err
		}

		// Run authentication request and print results
		auth, vcs, err := sdk.DID.Authenticate(did, pin)
		if err != nil {
			return err
		}
		if !auth {
			return errors.New("invalid credentials")
		}
		fmt.Printf("\n%s\n", vcs)
		return nil
	},
}

func init() {
	params := []cli.Param{
		{
			Name:      "did",
			Usage:     "user's identifier",
			FlagKey:   "list.did",
			ByDefault: "",
			Short:     "u",
		},
		{
			Name:      "pin",
			Usage:     "user's active PIN",
			FlagKey:   "list.pin",
			ByDefault: "",
			Short:     "p",
		},
	}
	if err := cli.SetupCommandParams(vcListCmd, params); err != nil {
		panic(err)
	}
	vcCmd.AddCommand(vcListCmd)
}
