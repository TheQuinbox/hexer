package hexer

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hexer",
	Short: "Encode or decode hex from your terminal",
	Long:  `Hexer allows you to convert to and from hex strings directly within your terminal.`,
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
