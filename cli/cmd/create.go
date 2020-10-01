package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aidtechnology/affinityctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.bryk.io/x/cli"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new DID",
	Aliases: []string{"new", "enroll"},
	Example: "affinityctl create -n my-id -e your@email.com",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get user parameters
		name := viper.GetString("create.name")
		if strings.TrimSpace(name) == "" {
			return errors.New("--name is required")
		}
		email := viper.GetString("create.email")
		if strings.TrimSpace(email) == "" {
			return errors.New("--email is required")
		}

		// Get SDK client
		sdk, err := client.New(nil)
		if err != nil {
			return err
		}

		// Get PIN
		pin := viper.GetString("create.pin")
		if strings.TrimSpace(pin) == "" {
			pin, err = getPIN(sdk)
			if err != nil {
				return err
			}
		}

		// Create new DID
		did, err := sdk.DID.Create(pin, email)
		if err != nil {
			return err
		}

		// Obtain and store did document
		doc, err := sdk.DID.Resolve(did)
		if err != nil {
			return err
		}
		if err = store(did, name, email, doc); err != nil {
			return err
		}

		// All good!
		fmt.Printf("DID generated: %s\n", did)
		return nil
	},
}

func init() {
	params := []cli.Param{
		{
			Name:      "name",
			Usage:     "reference name for the new DID",
			FlagKey:   "create.name",
			ByDefault: "",
			Short:     "n",
		},
		{
			Name:      "email",
			Usage:     "email address associated with the DID",
			FlagKey:   "create.email",
			ByDefault: "",
			Short:     "e",
		},
		{
			Name:      "pin",
			Usage:     "to be used for authentication",
			FlagKey:   "create.pin",
			ByDefault: "",
			Short:     "p",
		},
	}
	if err := cli.SetupCommandParams(createCmd, params); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(createCmd)
}
