package cmd

import (
	"github.com/alignse/otn-downloader/encode"
	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "encode data to the output",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := encode.Config{
			Fps:       *fps,
			ChunkSize: *chunkSize,
			Loop:      *loop,
		}
		encode.EncodToQRCode(*filename, cfg)
	},
}

var fps *int
var chunkSize *int
var filename *string
var loop *int

func init() {
	rootCmd.AddCommand(encodeCmd)
	fps = encodeCmd.Flags().Int("fps", 10, "the data encode fps")
	loop = encodeCmd.Flags().Int("loop", 3, "the number of times process")
	chunkSize = encodeCmd.Flags().IntP("chunk-size", "c", 60, "the chunk size of the input file")
	filename = encodeCmd.Flags().StringP("input-file", "f", "", "the source files")
}
