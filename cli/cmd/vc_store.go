package cmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.bryk.io/x/cli"
)

var vcStoreCmd = &cobra.Command{
	Use:     "store",
	Short:   "Store an existing VC in the user's wallet",
	Aliases: []string{"add", "save"},
	Example: "affinityctl vc store -u [DID] -p [PIN] path/to/vc.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get user parameters
		did := viper.GetString("store.did")
		if strings.TrimSpace(did) == "" {
			return errors.New("--did is required")
		}
		pin := viper.GetString("store.pin")
		if strings.TrimSpace(pin) == "" {
			p, err := cli.ReadSecure("Enter PIN: ")
			if err != nil {
				return err
			}
			pin = string(p)
		}

		// Read credential from file or standard input
		var contents []byte
		var err error
		if len(args) > 0 {
			contents, err = ioutil.ReadFile(filepath.Clean(args[0]))
		} else {
			contents, err = cli.ReadPipedInput(1 << 20) // 1MB
		}
		if err != nil {
			return err
		}

		// Decode credential
		vc := make(map[string]interface{})
		if err = json.Unmarshal(contents, &vc); err != nil {
			return err
		}

		// Get SDK client
		sdk, err := sdkClient()
		if err != nil {
			return err
		}

		// Run store operation
		return sdk.VC.Store(did, pin, vc)
	},
}

func init() {
	params := []cli.Param{
		{
			Name:      "did",
			Usage:     "user's identifier",
			FlagKey:   "store.did",
			ByDefault: "",
			Short:     "u",
		},
		{
			Name:      "pin",
			Usage:     "user's active PIN",
			FlagKey:   "store.pin",
			ByDefault: "",
			Short:     "p",
		},
	}
	if err := cli.SetupCommandParams(vcStoreCmd, params); err != nil {
		panic(err)
	}
	vcCmd.AddCommand(vcStoreCmd)
}
