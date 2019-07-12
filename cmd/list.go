package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(list)
}

var list = &cobra.Command{
	Use:   "list",
	Short: "Lists all google task",
	Long:  "Lists all google tasks from google task account or computer if no internet is found.",
	Run: func(cmd *cobra.Command, args []string) {
		//do something
	},
}
