package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(login)
}

var login = &cobra.Command{
	Use:   "login",
	Short: "Logs in to your google account",
	Long:  "Logs in to your google account using OAuth 2.0.",
	Run: func(cmd *cobra.Command, args []string) {
		//do something
	},
}
