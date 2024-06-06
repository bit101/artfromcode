// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/lures/l3d"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240608"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 180)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau   = blmath.Tau
	pi    = math.Pi
	count = 20000
)

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 220, 150, 800)

	s0 := wire.NewShape()
	lure0 := l3d.NewPickover()
	x, y, z := lure0.InitVals3d()
	s0.AddXYZ(x, y, z)

	random.Seed(0)
	for range count {
		x, y, z = lure0.Iterate(x, y, z)
		s0.AddXYZ(x, y, z)
	}
	s0.Translate(lure0.Center3d())
	s0.UniScale(lure0.Scale * 1.5)

	s0.Rotate(blmath.LoopSin(percent, -0.3, 0.3), 0, blmath.LoopSin(percent+0.25, -0.3, 0.3))
	s0.TranslateY(200 + blmath.LoopSin(percent, -40, 40))
	s1 := s0.Clone()

	s1.Cull(func(p *wire.Point) bool {
		return p.Y > 200
	})
	s0.Cull(func(p *wire.Point) bool {
		return p.Y < 200
	})

	s0.RenderPoints(2)
	s1.RenderPoints(0.5)

	context.GaussianBlur(20)
	s0.RenderPoints(1)
	s1.RenderPoints(0.25)
}
