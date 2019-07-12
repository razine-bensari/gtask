package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(push)
}

var push = &cobra.Command{
	Use:   "push",
	Short: "Pushes your local google task to your google account",
	Long:  "Pushes all your google tasks info in your computer to the google task api if internet connection is present and valid",
	Run: func(cmd *cobra.Command, args []string) {
		//do something
	},
}
