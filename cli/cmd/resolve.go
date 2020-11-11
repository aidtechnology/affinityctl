package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:     "resolve",
	Short:   "Retrieve the DID document for a given identifier",
	Aliases: []string{"get"},
	Example: "affinityctl resolve [DID]",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("you must specify a DID name")
		}

		// Get SDK client
		sdk, err := sdkClient()
		if err != nil {
			return err
		}

		// Send request and print results
		doc, err := sdk.DID.Resolve(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", doc)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(resolveCmd)
}
