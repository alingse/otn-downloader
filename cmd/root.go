package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "otn-downloader",
	Short: "single direction optical transport network downloader",
	Long:  `The data is encode as video by the OTN downloader, then the other devices can extract data from the display device or terminal or from some video recorders`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
