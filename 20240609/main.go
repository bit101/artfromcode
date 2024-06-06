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
	fileName := "20240609"

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
	tau = blmath.Tau
	pi  = math.Pi
)

var (
	model *wire.Shape
)

func init() {
	model = wire.ShapeFromXYZ("skull.xyz")
	model.ThinPoints(1, 6)
	model.UniScale(4)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := model.Clone()
	top := s.Culled(func(p *wire.Point) bool {
		return p.Y < -150
	})
	top.TranslateY(blmath.LoopSin(percent+0.25, 0, -300))

	s.Cull(func(p *wire.Point) bool {
		return p.Y > -150
	})
	s.AddShape(top)
	s.RotateX(-0.2)
	s.RotateY(percent * tau)

	a := wire.NewShape()
	lure := l3d.NewThomas()
	x, y, z := lure.InitVals3d()
	for range 10000 {
		a.AddXYZ(x, y, z)
		x, y, z = lure.Iterate(x, y, z)
	}
	a.Translate(lure.Center3d())
	a.UniScale(lure.Scale * 0.6)
	a.Rotate(percent*tau, percent*tau*2, 0)
	a.TranslateY(blmath.LoopSin(percent+0.25, -150, -320))

	a.RenderPoints(1.5)
	s.RenderPoints(1)

	context.GaussianBlur(20)
	a.RenderPoints(0.75)
	s.RenderPoints(0.5)
}
