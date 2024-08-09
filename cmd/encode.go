package cmd

import (
	"fmt"

	"github.com/alignse/otn-downloader/encode"
	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Convert data to the output",
	Long:  `It only support qrcode to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encode called")
		encode.EncodToQrCode()
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
