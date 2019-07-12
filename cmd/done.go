package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(done)
}

var done = &cobra.Command{
	Use:   "done",
	Short: "Marks a google task as complete (done status)",
	Long:  "Marks the corresponding google task in the 'done' category.",
	Run: func(cmd *cobra.Command, args []string) {
		//do something
	},
}
