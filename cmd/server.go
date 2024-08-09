package cmd

import (
	"fmt"

	"github.com/alignse/otn-downloader/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "serve the decode js and can merge the files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
