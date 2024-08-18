package cmd

import (
	"strconv"

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
			Slices:    parseInts(*slices),
		}
		encode.EncodToQRCode(*filename, cfg)
	},
}

var fps *int
var chunkSize *int
var filename *string
var loop *int
var slices *[]string

func init() {
	rootCmd.AddCommand(encodeCmd)
	fps = encodeCmd.Flags().Int("fps", 10, "the data encode fps")
	loop = encodeCmd.Flags().Int("loop", 3, "the number of times process")
	chunkSize = encodeCmd.Flags().IntP("chunk-size", "c", 60, "the chunk size of the input file")
	filename = encodeCmd.Flags().StringP("input-file", "f", "", "the source files")
	slices = encodeCmd.Flags().StringSliceP("slices", "s", []string{}, "this miss slice of the chunks")
}

func parseInts(strValues []string) map[int]bool {
	var result = map[int]bool{}
	for _, str := range strValues {
		if val, err := strconv.Atoi(str); err == nil {
			result[val] = true
		}
	}
	return result
}
