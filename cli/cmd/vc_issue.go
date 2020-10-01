package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/aidtechnology/affinityctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.bryk.io/x/cli"
)

var vcIssueCmd = &cobra.Command{
	Use:     "issue",
	Short:   "Issue a new verifiable credential",
	Aliases: []string{"create", "new", "generate"},
	Example: "affinityctl vc issue -i [ISSUER DID] -p [ISSUER PIN] -u [SUBJECT DID] payload.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get user parameters
		issuerDID := viper.GetString("issuer.did")
		if strings.TrimSpace(issuerDID) == "" {
			return errors.New("--issuer-did is required")
		}
		subject := viper.GetString("vc.subject")
		if strings.TrimSpace(subject) == "" {
			return errors.New("--subject is required")
		}
		issuerPIN := viper.GetString("issuer.pin")
		if strings.TrimSpace(issuerPIN) == "" {
			p, err := cli.ReadSecure("Enter PIN: ")
			if err != nil {
				return err
			}
			issuerPIN = string(p)
		}

		// Get issuer instance
		issuer := &client.Issuer{
			DID: issuerDID,
			PIN: issuerPIN,
		}

		// Read payload from file or standard input
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
		payload := make(map[string]interface{})
		if err = json.Unmarshal(contents, &payload); err != nil {
			return err
		}

		// Get SDK client
		sdk, err := client.New(nil)
		if err != nil {
			return err
		}

		// Run issue operation
		vc, err := sdk.VC.Issue(issuer, subject, payload)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", vc)
		return nil
	},
}

func init() {
	params := []cli.Param{
		{
			Name:      "issuer-did",
			Usage:     "issuer's identifier",
			FlagKey:   "issuer.did",
			ByDefault: "",
			Short:     "i",
		},
		{
			Name:      "issuer-pin",
			Usage:     "user's active PIN",
			FlagKey:   "issuer.pin",
			ByDefault: "",
			Short:     "p",
		},
		{
			Name:      "subject",
			Usage:     "user's identifier",
			FlagKey:   "vc.subject",
			ByDefault: "",
			Short:     "u",
		},
	}
	if err := cli.SetupCommandParams(vcIssueCmd, params); err != nil {
		panic(err)
	}
	vcCmd.AddCommand(vcIssueCmd)
}
