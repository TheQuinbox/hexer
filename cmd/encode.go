package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var encodeCommand = &cobra.Command{
	Use: "encode",
	Aliases: []string{"e"},
	Short: "Encodes a string to hex",
	Args: cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		input := args[0]
		result := hex.EncodeToString([]byte(input))
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(encodeCommand)
}
