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
	fileName := "20240611"

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
	a0 = wire.NewShape()
)

func init() {
	lure := l3d.NewHalvorsen()

	x, y, z := lure.InitVals3d()
	for range 20000 {
		a0.AddXYZ(x, y, z)
		x, y, z = lure.Iterate(x, y, z)
	}
	a0.Translate(lure.Center3d())
	a0.UniScale(lure.Scale)
	a0.TranslateY(-50)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	y := easing.CubicEaseInOut(percent, -50, 550)

	a := a0.Rotated(0.2, percent*tau, 0.3)
	a.TranslateY(y)
	b := a.Culled(func(p *wire.Point) bool {
		return p.Y > 250
	})
	a.Cull(func(p *wire.Point) bool {
		return p.Y < 250
	})

	wire.SetRGB(0.5, 0, 0)
	b.RenderPoints(1)

	wire.SetRGB(1, 1, 1)
	a.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		wire.SetRGB(1, 1, 1)
		a.RenderPoints(1)
	}
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	y := easing.CubicEaseInOut(percent, 550, -50)

	a := a0.Rotated(0.2, percent*tau, 0.3)
	a.TranslateY(y)
	b := a.Culled(func(p *wire.Point) bool {
		return p.Y > 250
	})
	a.Cull(func(p *wire.Point) bool {
		return p.Y < 250
	})

	wire.SetRGB(0.5, 0, 0)
	b.RenderPoints(1)

	wire.SetRGB(1, 0, 0)
	a.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		a.RenderPoints(1)
	}
	wire.SetRGB(1, 1, 1)
	a.Translate(1, -1, 0)
	a.RenderPoints(0.5)
}

func scene3(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	x := 1000.0

	a := a0.Rotated(0.2, percent*tau, 0.3)
	a.TranslateY(-50)
	a.TranslateX(easing.CubicEaseInOut(percent, 0, x))
	b := a.TranslatedX(-x)

	wire.SetRGB(1, 0, 0)
	a.RenderPoints(2)
	wire.SetRGB(1, 1, 1)
	b.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		wire.SetRGB(1, 0, 0)
		a.RenderPoints(1)
		wire.SetRGB(1, 1, 1)
		b.RenderPoints(1)
	}
	wire.SetRGB(1, 1, 1)
	a.Translate(1, -1, 0)
	a.RenderPoints(0.5)
}
