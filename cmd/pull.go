package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(pull)
}

var pull = &cobra.Command{
	Use:   "pull",
	Short: "Pulls your remote google tasks to your local computer",
	Long:  "Pulls all your google tasks to your local computer if not present.",
	Run: func(cmd *cobra.Command, args []string) {
		//do something
	},
}
