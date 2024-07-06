// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/lures/l3d"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240707"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 420)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau  = blmath.Tau
	pi   = math.Pi
	blur = true
)

var (
	shape *wire.Shape
	bg    *cairo.Surface
)

func init() {
	bg, _ = cairo.NewSurfaceFromPNG("woods.png")
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.ClearClear()
	context.SetSourceWhite()
	wire.InitWorld(context, 200, 200, 1000)

	pickover := l3d.NewPickover()
	pickover.Params.A = blmath.LoopSin(percent, -0.5, -0.9)
	shape = wire.ShapeFromLure(pickover, 20000)
	s := shape.Clone()

	s.Rotate(percent*tau, percent*2*tau, 0)

	s.RenderPoints(1)

	if blur {
		context.GaussianBlur(20)
		s.RenderPoints(0.5)
	}
	context.DrawSurfaceUnder(bg, 0, 0)
}
