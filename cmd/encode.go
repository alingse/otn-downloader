package cmd

import (
	"fmt"

	"github.com/alignse/otn-downloader/encode"
	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "encode data to the output",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encode called")
		encode.EncodToQRCode("example.txt", 5, 3)
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
