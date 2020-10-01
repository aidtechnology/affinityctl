package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/aidtechnology/affinityctl/client"
	"github.com/spf13/cobra"
	"go.bryk.io/x/cli"
)

var vcVerifyCmd = &cobra.Command{
	Use:     "verify",
	Short:   "Verify an existing credential",
	Aliases: []string{"check"},
	Example: "affinityctl vc verify path/to/vc.json",
	RunE: func(cmd *cobra.Command, args []string) error {
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
		sdk, err := client.New(nil)
		if err != nil {
			return err
		}

		// Run verification
		result, err := sdk.VC.Verify(vc)
		if err != nil {
			return err
		}
		if !result {
			return errors.New("invalid credential")
		}
		fmt.Println("credential is valid")
		return nil
	},
}

func init() {
	vcCmd.AddCommand(vcVerifyCmd)
}
