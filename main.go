package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ajstarks/svgo"
)

func main() {

	var (
		generator       = flag.String("generator", "simple", "Choose your generator (simple, centered, grid, sinusoidal).")
		height          = flag.Int("height", 720, "The height of the canvas.")
		width           = flag.Int("width", 1280, "The width of the canvas.")
		radius          = flag.Int("radius", 5, "The dots radius.")
		n               = flag.Int("n", 1000, "The number of dots to draw.")
		backgroundColor = flag.String("bc", "#234d51", "The background color.")
		dotsColors      = flag.String("dc", "#ff513f,#ffce00", "The dots colors.")
	)

	flag.Parse()

	settings := NewDotGeneratorSettings(*width, *height, *radius)

	var g DotGenerator

	switch *generator {
	case "simple":
		g = NewSimpleGenerator(settings)
	case "centered":
		g = NewCenteredGenerator(settings)
	case "grid":
		g = NewGridGenerator(settings)
	case "sinusoidal":
		g = NewSinusoidalGenerator(settings)
	default:
		fmt.Fprintf(os.Stderr, "Unknown %q generator\n", *generator)
		flag.Usage()
		os.Exit(1)
	}

	canvas := svg.New(os.Stdout)
	canvas.Start(*width, *height)
	canvas.Rect(0, 0, *width, *height, Fill(HexToColor(*backgroundColor)))

	c := NewColorGenerator(HexesToColors(*dotsColors))

	for i := 0; i < *n; i++ {
		x, y := g.Generate()
		color := c.Generate()
		canvas.Circle(x, y, *radius, Fill(color))
	}

	canvas.End()
}
