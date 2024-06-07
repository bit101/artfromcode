// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/lures/l3d"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240610"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 180)
		program.AddSceneWithFrames(scene2, 180)
		program.AddSceneWithFrames(scene3, 180)
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
	sphere = wire.Sphere(300, 64, 128, true, true)
	gp     = wire.GridPlane(800, 800, 20, 20)
	a0     = wire.NewShape()
	a1     = wire.NewShape()
)

func init() {
	sphere.Randomize(5)
	lure := l3d.NewHalvorsen()

	x, y, z := lure.InitVals3d()
	for range 20000 {
		a0.AddXYZ(x, y, z)
		x, y, z = lure.Iterate(x, y, z)
	}
	a0.Translate(lure.Center3d())
	a0.UniScale(lure.Scale)
	a0.TranslateY(-50)

	x, y, z = lure.InitVals3d()
	for range 5000 {
		a1.AddXYZ(x, y, z)
		x, y, z = lure.Iterate(x, y, z)
	}
	a1.Translate(lure.Center3d())
	a1.UniScale(lure.Scale)
	a1.Randomize(20)
	a1.TranslateY(-50)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	y := easing.CubicEaseInOut(percent, -320, 320)

	s0 := sphere.Rotated(0.2, percent*tau, 0.3)
	s0.Cull(func(p *wire.Point) bool {
		return p.Y > y
	})

	a := a1.Rotated(0.2, percent*tau, 0.3)
	a.Cull(func(p *wire.Point) bool {
		return p.Y < y
	})

	g := gp.TranslatedY(y)
	g.Subdivide(5)
	g.RotateY(easing.CubicEaseInOut(percent, 0, -pi/2))

	s0.RenderPoints(2)
	a.RenderPoints(2)
	g.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s0.RenderPoints(1)
		a.RenderPoints(1)
		g.Stroke(1)
	}
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	y := easing.CubicEaseInOut(percent, 320, -320)

	s0 := a1.Rotated(0.2, percent*tau, 0.3)
	s0.Cull(func(p *wire.Point) bool {
		return p.Y < y
	})

	a := a0.Rotated(0.2, percent*tau, 0.3)
	a.Cull(func(p *wire.Point) bool {
		return p.Y > y
	})

	g := gp.TranslatedY(y)
	g.Subdivide(5)
	g.RotateY(easing.CubicEaseInOut(percent, 0, -pi/2))

	s0.RenderPoints(2)
	a.RenderPoints(2)
	g.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s0.RenderPoints(1)
		a.RenderPoints(1)
		g.Stroke(1)
	}
}

func scene3(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	x := 1000.0

	s0 := sphere.Rotated(0.2, percent*tau, 0.3)
	s0.TranslateX(easing.CubicEaseInOut(percent, -x, 0))

	a := a0.Rotated(0.2, percent*tau, 0.3)
	a.TranslateX(easing.CubicEaseInOut(percent, 0, x))

	g := gp.TranslatedY(-320)
	g.Subdivide(5)

	s0.RenderPoints(2)
	a.RenderPoints(2)
	g.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s0.RenderPoints(1)
		a.RenderPoints(1)
		g.Stroke(1)
	}
}
