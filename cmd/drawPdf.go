package cmd

import (
	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dpdf"
	"github.com/spf13/viper"
	"math"

	//"github.com/llgcode/draw2d/draw2dkit"
	"github.com/spf13/cobra"
	"image/color"
)

// drawPdfCmd represents the drawPdf command
var drawPdfCmd = &cobra.Command{
	Use:   "drawPdf",
	Short: "Create a pdf doc with graphics",
	Long:  `A longer descriptions`,
	Run: func(cmd *cobra.Command, args []string) {
		drawPdf()
	},
}

func init() {
	rootCmd.AddCommand(drawPdfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// drawPdfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// drawPdfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func drawPdf() {
	fmt.Println("drawPdf called")
	path := viper.Get("path.sandbox")

	// Initialize the graphic context on an RGBA image
	dest := draw2dpdf.NewPdf("P", "mm", "A4")
	gc := draw2dpdf.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(2)

	// Draw a closed shape
	//gc.MoveTo(10, 10) // should always be called first for a new path
	//gc.LineTo(100, 50)
	//gc.QuadCurveTo(100, 10, 10, 10)
	//gc.Close()
	//gc.FillStroke()

	// Main draws vertically spaced lines and returns the filename.
	//gc.SetFillRule(draw2d.FillRuleWinding)
	//gc.Clear()
	// Draw the line
	//for x := 5.0; x < 297; x += 10 {
	//	Draw(gc, x, 0, x, 210)
	//}
	//gc.ClearRect(100, 75, 197, 135)
	//draw2dkit.Ellipse(gc, 148.5, 105, 35, 25)
	//gc.SetFillColor(color.RGBA{0xff, 0xff, 0x44, 0xff})
	//gc.FillStroke()

	//gc.SetLineWidth(1)

	xStart := 20.0
	yStart := 50.0
	shift := 17.0

	// draw sin continuously from point to point
	gc.MoveTo(xStart, yStart)
	for i := 0.00; i < shift; i += 0.1 {
		x, y := gc.LastPoint()
		gc.MoveTo(x, y)

		x1 := xStart + i*4
		y1 := yStart + math.Sin(i)*12
		gc.LineTo(x1, y1)
		gc.Stroke()
	}

	gc.SetLineWidth(0.7)
	xStart2 := xStart + 100
	// draw sin by separate lines from x
	for i := 0.00; i < shift; i += 0.5 {

		gc.MoveTo(xStart2+i*4, yStart)

		y1 := yStart + math.Sin(i)*12
		gc.LineTo(xStart2+i*4, y1)
		gc.Stroke()
	}

	// Save to file
	draw2dpdf.SaveToPdfFile(fmt.Sprintf("%v", path)+"hello.pdf", dest)
}

// Draw vertically spaced lines
func Draw(gc draw2d.GraphicContext, x0, y0, x1, y1 float64) {
	// Draw a line
	gc.MoveTo(x0, y0)
	gc.LineTo(x1, y1)
	gc.Stroke()
}
