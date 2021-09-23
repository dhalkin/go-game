package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// drawPngToPdfCmd represents the drawPngToPdf command
var drawPngToPdfCmd = &cobra.Command{
	Use:   "drawPngToPdf",
	Short: "Content from png to pdf",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		pngToPdf()
	},
}

func init() {
	rootCmd.AddCommand(drawPngToPdfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// drawPngToPdfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// drawPngToPdfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func pngToPdf() {
	fmt.Println("drawPngToPdf called")
	// prepare ./source/hello.png
	drawPng()

}
