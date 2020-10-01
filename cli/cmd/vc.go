package cmd

import (
	"github.com/spf13/cobra"
)

var vcCmd = &cobra.Command{
	Use:   "vc",
	Short: "Verifiable credential operations",
}

func init() {
	rootCmd.AddCommand(vcCmd)
}
