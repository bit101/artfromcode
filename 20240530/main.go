// Package main renders an image or video
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

//revive:disable:unused-parameter
const (
	tau = blmath.Tau
	pi  = math.Pi
)

func main() {
	renderTarget := target.Video
	fileName := "20240530"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 420)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

var shape *wire.Shape

func init() {
	shape = wire.NewShape()

	x := 0.1
	y := 0.1
	z := 0.1

	for range 40000 {
		x, y, z = l3d.FourWing(x, y, z, 0.05)
		shape.AddXYZ(x, y, z)
	}
	shape.UniScale(160)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	lure := shape.Rotated(percent*tau, percent*2*tau, 0)

	a := (percent + 0.0625) * 2 * tau
	p := wire.NewPoint(math.Cos(a)*200, math.Sin(a)*300, 0)
	lure.Points.Push(p, 160)

	s := wire.Sphere(100, 12, 24, true, true)
	s.Rotate(percent*tau, percent*2*tau, 0)
	s.Translate(p.X, p.Y, p.Z)

	context.SetLineWidth(2)
	lure.RenderPoints(2)
	s.Stroke()

	context.GaussianBlur(20)
	lure.RenderPoints(1)
	context.SetLineWidth(1)
	s.Stroke()
}
