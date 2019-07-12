package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(version)
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of the command line",
	Long:  "Outputs the version of the command line in the console.",
	Run: func(cmd *cobra.Command, args []string) {
		//logic
	},
}
