package cmd


import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(remove)
}

var remove = &cobra.Command{
	Use: "remove",
	short: "Removes the gTask from your tasks",
	Long: "Removes the specified gtask from your computer and also from google tasks",
	Run: func(cmd *cobra.Command, args []String){
		//do something
	}
}