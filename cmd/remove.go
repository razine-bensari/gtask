package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(remove)
}

var remove = &cobra.Command{
	Use:   "remove",
	Short: "Removes the gTask from your tasks",
	Long:  "Removes the specified gtask from your computer and also from google tasks",
	Run: func(cmd *cobra.Command, args []string) {
		//do something
	},
}
