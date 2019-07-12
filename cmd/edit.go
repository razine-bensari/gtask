package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(edit)
}

var edit = &cobra.Command{
	Use:   "edit",
	Short: "Edits a google task",
	Long:  "Edits a google task given the correct parameter and updates it in your google task.",
	Run: func(cmd *cobra.Command, args []string) {
		//do something
	},
}
