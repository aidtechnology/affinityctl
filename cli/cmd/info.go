package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Display information on a given identifier",
	Aliases: []string{"inspect", "details", "describe"},
	Example: "affinityctl info [DID_REFERENCE_NAME]",
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("you must specify a DID name")
		}
		rec, err := details(args[0])
		if err != nil {
			return err
		}
		js, err := json.MarshalIndent(rec.Doc, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("DID: %s\n", rec.ID)
		fmt.Printf("\n%s\n", js)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
