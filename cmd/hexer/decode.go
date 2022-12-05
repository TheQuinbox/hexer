package hexer

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thequinbox/hexer/pkg/hextools"
)

var decodeCommand = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"d"},
	Short:   "Decodes a hex string back to plain text",
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		input := args[0]
		result := hextools.Decode(input)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(decodeCommand)
}
