// Package main renders an image or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

//revive:disable:unused-parameter
const (
	tau = blmath.Tau
	pi  = math.Pi
)

func main() {
	renderTarget := target.Video
	fileName := "20240527"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 400)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	line1 := wire.NewString("without", wire.FontAsteroid).AsLine(20)
	line2 := wire.NewString("connection", wire.FontAsteroid).AsLine(20)
	line3 := wire.NewString("there is", wire.FontAsteroid).AsLine(20)
	line4 := wire.NewString("nothing.", wire.FontAsteroid).AsLine(20)
	text := wire.NewShape()
	text.AddShape(line1.TranslatedY(-300))
	text.AddShape(line2.TranslatedY(-100))
	text.AddShape(line3.TranslatedY(100))
	text.AddShape(line4.TranslatedY(300))
	text.Rotate(-0.3, percent*tau, 0)
	text.TranslateZ(800)

	shape := wire.RandomSurfaceSphere(400, 500)
	shape.Rotate(percent*tau, percent*2*tau, 0)

	maxSize := blmath.LoopSin(percent, 1, 200)
	for i := 0; i < len(shape.Points)-1; i++ {
		for j := i + 1; j < len(shape.Points); j++ {
			size := shape.Points[i].Distance(shape.Points[j])
			if size < maxSize {
				shape.AddSegmentByIndex(i, j)
			}
		}
	}
	shape.Points.Project()
	for _, seg := range shape.Segments {
		size := seg.Length()
		context.SetLineWidth((1 - size/maxSize) * 2)
		seg.Stroke()
	}
	context.SetLineWidth(blmath.LoopSin(percent, 4, 0))
	text.Stroke()

	context.GaussianBlur(12)
	shape.Points.Project()
	for _, seg := range shape.Segments {
		size := seg.Length()
		context.SetLineWidth((1 - size/maxSize) * 1)
		seg.Stroke()
	}
	context.SetLineWidth(blmath.LoopSin(percent, 1, 0))
	text.Stroke()

	context.Starfield(100, 0.5)
}
