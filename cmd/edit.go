package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(edit)
}

var edit = &cobra.Command{
	Use: "edit",
	short: "Edits a google task",
	Long: "Edits a google task given the correct parameter and updates it in your google task",
	Run: func(cmd *cobra.Command, args []String){
		//do something
	}
}