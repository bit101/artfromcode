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

	switch renderTarget {
	case target.Image:
		render.Image(800, 800, "out/out.png", scene1, 0.0)
		render.ViewImage("out/out.png")
		break

	case target.Video:
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 360)
		program.AddSceneWithFrames(scene2, 360)
		program.AddSceneWithFrames(scene3, 360)
		program.RenderVideo("out/frames", "out/out.mp4")
		render.PlayVideo("out/out.mp4")
		break
	}
}

//revive:disable:unused-parameter
const (
	tau = blmath.Tau
	pi  = math.Pi
)

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	a0 := l3d.NewAizawa()
	a1 := l3d.NewFourWings()
	x0, y0, z0 := a0.InitVals3d()
	x1, y1, z1 := a1.InitVals3d()

	s := wire.NewShape()
	for range 10000 {
		x0, y0, z0 = a0.Iterate(x0, y0, z0)
		x1, y1, z1 = a1.Iterate(x1, y1, z1)
		x := blmath.Lerp(percent, x0, x1)
		y := blmath.Lerp(percent, y0, y1)
		z := blmath.Lerp(percent, z0, z1)
		s.AddXYZ(x, y, z)

	}

	s.UniScale(blmath.Lerp(percent, a0.Scale, a1.Scale))
	s.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	s.RenderPoints(2)

	context.GaussianBlur(20)
	s.RenderPoints(1)
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	a0 := l3d.NewFourWings()
	a1 := l3d.NewRossler()
	x0, y0, z0 := a0.InitVals3d()
	x1, y1, z1 := a1.InitVals3d()

	s := wire.NewShape()
	for range 10000 {
		x0, y0, z0 = a0.Iterate(x0, y0, z0)
		x1, y1, z1 = a1.Iterate(x1, y1, z1)
		x := blmath.Lerp(percent, x0, x1)
		y := blmath.Lerp(percent, y0, y1)
		z := blmath.Lerp(percent, z0, z1)
		s.AddXYZ(x, y, z)

	}

	s.UniScale(blmath.Lerp(percent, a0.Scale, a1.Scale))
	s.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	s.RenderPoints(2)

	context.GaussianBlur(20)
	s.RenderPoints(1)
}

func scene3(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	a0 := l3d.NewRossler()
	a1 := l3d.NewAizawa()
	x0, y0, z0 := a0.InitVals3d()
	x1, y1, z1 := a1.InitVals3d()

	s := wire.NewShape()
	for range 10000 {
		x0, y0, z0 = a0.Iterate(x0, y0, z0)
		x1, y1, z1 = a1.Iterate(x1, y1, z1)
		x := blmath.Lerp(percent, x0, x1)
		y := blmath.Lerp(percent, y0, y1)
		z := blmath.Lerp(percent, z0, z1)
		s.AddXYZ(x, y, z)

	}

	s.UniScale(blmath.Lerp(percent, a0.Scale, a1.Scale))
	s.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	s.RenderPoints(2)

	context.GaussianBlur(20)
	s.RenderPoints(1)
}
