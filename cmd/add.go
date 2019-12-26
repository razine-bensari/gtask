package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/api/tasks/v1"
)

func init() {
	RootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Adds a google task",
	Long:  "Adds a google task to your google task account. If no internet connection, it persists it on your computer.",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := tasks.New(googleTaskClient)

	},
}
