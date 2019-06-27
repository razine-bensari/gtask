package cmd


import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(version)
}

var version = &cobra.Command{
	Use: "version",
	short: "Shows the version of the command line",
	Long: "Outputs the version of the command line in the console.",
	Run: func(cmd *cobra.Command, args []String){
			var invoker
			invoker.executeCommand( vesioncommnand)
	}
}