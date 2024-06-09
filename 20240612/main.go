// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/lures"
	"github.com/bit101/lures/l3d"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240612"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 180)
		program.AddSceneWithFrames(scene2, 180)
		program.AddSceneWithFrames(scene3, 180)
		program.AddSceneWithFrames(scene4, 180)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau   = blmath.Tau
	pi    = math.Pi
	blur  = true
	count = 20000
)

var (
	attr lures.Lure3d
	s    *wire.Shape
	p    wire.PointList
)

func init() {
	attr = l3d.NewChenLee()
	s = wire.ShapeFromLure(attr, count)
	p = s.Points.Normalized()
	p.UniScale(300)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	renderScene(context, percent, 0)
}

func scene2(context *cairo.Context, width, height, percent float64) {
	renderScene(context, percent, percent)
}

func scene3(context *cairo.Context, width, height, percent float64) {
	renderScene(context, percent, 1)
}

func scene4(context *cairo.Context, width, height, percent float64) {
	renderScene(context, percent, 1-percent)
}

func renderScene(context *cairo.Context, percent, t float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	s1 := s.Clone()
	s1.Points.Lerp(t, p)
	s1.Rotate(percent*tau, percent*tau*2, 0)
	s1.RenderPoints(3)

	if blur {
		context.GaussianBlur(20)
		s1.RenderPoints(1.5)
	}
}
