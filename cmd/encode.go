package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thequinbox/hexer/hextools"
)

var encodeCommand = &cobra.Command{
	Use: "encode",
	Aliases: []string{"e"},
	Short: "Encodes a string to hex",
	Args: cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		input := args[0]
		result := hextools.Encode(input)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(encodeCommand)
}
