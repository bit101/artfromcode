// Package main renders an image or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
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
	fileName := "20240522"

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
	wire.InitWorld(context, 200, 200, 800)
	wire.SetClipping(50, 10000)

	text := wire.NewString("tunnel", wire.FontArcade).AsLine(20)
	text.TranslateY(-600)
	text.RotateY(-pi / 2)

	torus := wire.Torus(200, 100, tau, 32, 20, true, true)
	torus.RotateX(pi / 2)
	shape := wire.NewShape()
	shape.AddShape(torus.TranslatedZ(-400).RotatedZ(percent * pi))
	shape.AddShape(torus.UniScaled(1.5).RotatedZ(percent * -pi))
	shape.AddShape(torus.TranslatedZ(400).RotatedZ(percent * pi))

	shape.RotateY(blmath.LoopSin(percent+0.25, 0, pi))
	shape.TranslateZ(blmath.LoopSin(percent, -800, 800))

	text.RotateY(blmath.LoopSin(percent+0.25, 0, pi))
	text.TranslateZ(blmath.LoopSin(percent, -800, 800))

	context.SetLineWidth(2)
	shape.Stroke()
	context.SetLineWidth(4)
	text.Stroke()

	context.GaussianBlur(20)
	context.SetLineWidth(0.5)
	shape.Stroke()
	context.SetLineWidth(2)
	text.Stroke()
}
