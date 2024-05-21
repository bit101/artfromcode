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
	fileName := "20240525"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 420)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.RandomSurfaceSphere(300, 10000)
	txt := wire.NewString("noise noise noise", wire.FontArcade).AsCylinder(330, 10)
	txt.Subdivide(10)
	shape.AddShape(txt)
	shape.Points.Noisify(wire.NewPoint(0, 0, 0), 0.005, blmath.LoopSin(percent, 0, 0.5))
	shape.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(4)
	shape.Stroke()
	shape.RenderPoints(2)

	context.GaussianBlur(20)
	context.SetLineWidth(2)
	shape.Stroke()
	shape.RenderPoints(1)
}
