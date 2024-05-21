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
	fileName := "20240523"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 360)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.RandomSurfaceBox(800, 400, 200, 5000)
	txt := wire.NewString("boxed", wire.FontAsteroid).AsLine(20)
	txt.TranslateY(25)
	shape.AddShape(txt)

	shape.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	shape.RenderPoints(2)
	shape.Stroke()

	context.GaussianBlur(20)
	context.SetLineWidth(1)
	shape.RenderPoints(1)
	shape.Stroke()
}

func pull(shape *wire.Shape, point *wire.Point, force float64) {
	for _, p := range shape.Points {
		dist := p.Distance(point)
		dist = dist * dist
		f := math.Min(force/dist, 1)
		p.X += (point.X - p.X) * f
		p.Y += (point.Y - p.Y) * f
		p.Z += (point.Z - p.Z) * f
	}
}
