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
	fileName := "20240524"

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

	shape := wire.RandomSurfaceBox(500, 500, 500, 15000)
	shape.Cull(func(p *wire.Point) bool {
		radius := 200.0
		return (wire.NewPoint(250, 0, 0).Distance(p) > radius) &&
			(wire.NewPoint(-250, 0, 0).Distance(p) > radius) &&
			(wire.NewPoint(0, 250, 0).Distance(p) > radius) &&
			(wire.NewPoint(0, -250, 0).Distance(p) > radius) &&
			(wire.NewPoint(0, 0, 250).Distance(p) > radius) &&
			(wire.NewPoint(0, 0, -250).Distance(p) > radius)
	})

	x := percent * tau
	y := percent * tau * 2
	z := 0.0

	shape.AddShape(wire.Box(100, 100, 100).Rotated(x, y, z).TranslatedX(400))
	shape.AddShape(wire.Box(100, 100, 100).Rotated(x, z, y).TranslatedX(-400))
	shape.AddShape(wire.Box(100, 100, 100).Rotated(y, x, z).TranslatedY(400))
	shape.AddShape(wire.Box(100, 100, 100).Rotated(y, z, x).TranslatedY(-400))
	shape.AddShape(wire.Box(100, 100, 100).Rotated(z, x, y).TranslatedZ(400))
	shape.AddShape(wire.Box(100, 100, 100).Rotated(z, y, x).TranslatedZ(-400))

	shape.AddShape(wire.RandomInnerSphere(50, 250).Rotated(x, y, z).TranslatedX(400))
	shape.AddShape(wire.RandomInnerSphere(50, 250).Rotated(x, z, y).TranslatedX(-400))
	shape.AddShape(wire.RandomInnerSphere(50, 250).Rotated(y, x, z).TranslatedY(400))
	shape.AddShape(wire.RandomInnerSphere(50, 250).Rotated(y, z, x).TranslatedY(-400))
	shape.AddShape(wire.RandomInnerSphere(50, 250).Rotated(z, x, y).TranslatedZ(400))
	shape.AddShape(wire.RandomInnerSphere(50, 250).Rotated(z, y, x).TranslatedZ(-400))

	shape.AddShape(wire.Box(506, 506, 506))

	shape.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	shape.RenderPoints(2)
	shape.Stroke()

	context.GaussianBlur(12)
	context.SetLineWidth(0.5)
	shape.RenderPoints(1)
	shape.Stroke()

}
