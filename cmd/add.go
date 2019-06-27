package cmd

import (
	"github.com/spf13/cobra"
	"gtask/invoker"
)

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use: "add",
	short: "Adds a google task",
	Long: "Adds a google task to your google task account. If no internet connection, it persists it on your computer.",
	Run: func(cmd *cobra.Command, args []String){ 
		var invoker 
		invoker.executeCommands(addCommandd)
	}
}