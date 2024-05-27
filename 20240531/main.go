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
	fileName := "20240531"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 600)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

var shape *wire.Shape

func init() {
	shape = wire.RandomInnerBox(3000, 400, 400, 100000)

}
func scene1(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	wire.InitWorld(context, 200, 200, 800)

	s := shape.Clone()
	s.TranslateX(blmath.LoopSin(percent*2, -1900, 1900))
	for _, p := range s.Points {
		if math.Abs(p.X) < 400 {
			a := math.Cos(p.X / 400 * pi)
			a = blmath.Map(a, -1, 1, 1, 0.25)
			p.Scale(1, a, a)
		}
	}

	s2 := wire.Box(3000, 400, 400)
	s2.Subdivide(10)
	s2.TranslateX(blmath.LoopSin(percent*2, -1900, 1900))
	for _, p := range s2.Points {
		if math.Abs(p.X) < 400 {
			a := math.Cos(p.X / 400 * pi)
			a = blmath.Map(a, -1, 1, 1, 0.25)
			p.Scale(1, a, a)
		}
	}

	t := wire.Torus(200, 100, tau, 32, 20, true, true)
	t.RotateZ(pi / 2)

	s.Rotate(percent*tau, (percent+0.25)*tau, percent*2*tau)
	s2.Rotate(percent*tau, (percent+0.25)*tau, percent*2*tau)
	t.Rotate(percent*tau, (percent+0.25)*tau, percent*2*tau)

	context.SetLineWidth(1)
	t.Stroke()
	s2.Stroke()
	s.RenderPoints(1.5)

	context.GaussianBlur(12)
	context.SetLineWidth(0.5)
	s.RenderPoints(0.75)
	s2.Stroke()
	t.Stroke()
}
