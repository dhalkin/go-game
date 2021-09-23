package cmd

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"image"
	"image/color"
	"math"
)

// drawPngCmd represents the drawPng command
var drawPngCmd = &cobra.Command{
	Use:   "drawPng",
	Short: "Create a png file with graphics",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		drawPng()
	},
}

func init() {
	rootCmd.AddCommand(drawPngCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// drawPngCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// drawPngCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func drawPng() {
	fmt.Println("drawPng called")
	path := viper.Get("path.sandbox")

	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 600, 200))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0xff, 0xff})
	gc.SetStrokeColor(color.RGBA{0xff, 0x44, 0x44, 0xff})
	gc.SetLineWidth(3)

	// Draw a closed shape
	//gc.BeginPath() // Initialize a new path
	//gc.MoveTo(10, 10) // Move to a position to start the new path
	//gc.LineTo(100, 50)
	//gc.QuadCurveTo(100, 10, 10, 10)
	//gc.Close()
	//gc.FillStroke()

	xStart := 20.0
	yStart := 50.0
	shift := 20.0

	// draw sin continuously from point to point
	gc.MoveTo(xStart, yStart)
	for i := 0.00; i < shift; i += 0.1 {
		x, y := gc.LastPoint()
		gc.MoveTo(x, y)

		x1 := xStart + i*14
		y1 := yStart + math.Sin(i)*40
		gc.LineTo(x1, y1)
		gc.Stroke()
	}

	gc.SetLineWidth(1.00)
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	xStart2 := xStart + 300
	// draw sin by separate lines from x
	for i := 0.00; i < shift; i += 0.5 {

		gc.MoveTo(xStart2+i*10, yStart)

		y1 := yStart + math.Sin(i)*50
		gc.LineTo(xStart2+i*10, y1)
		gc.Stroke()
	}

	// Save to file
	draw2dimg.SaveToPngFile(fmt.Sprintf("%v", path)+"hello.png", dest)
}
