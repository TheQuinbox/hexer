package cmd

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var decodeCommand = &cobra.Command{
	Use: "decode",
	Aliases: []string{"d"},
	Short: "Decodes a hex string back to plain text",
	Args: cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		input := args[0]
		result, err := hex.DecodeString(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(result))
	},
}

func init() {
	rootCmd.AddCommand(decodeCommand)
}
