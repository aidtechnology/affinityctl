package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List identifiers generated",
	Aliases: []string{"ls"},
	RunE: func(cmd *cobra.Command, args []string) error {
		list := list()
		if len(list) == 0 {
			fmt.Println("No identifiers registered")
			return nil
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.DiscardEmptyColumns)
		_, _ = fmt.Fprintln(w, "Name  \tEmail  \tDID")
		for _, r := range list {
			_, _ = fmt.Fprintf(w, "%s\t%s\t%s\n", r.Name, r.Email, r.ID)
		}
		return w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
