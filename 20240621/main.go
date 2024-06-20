// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240621"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 360)
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
)

func init() {
	// shape = wire.Sphere(400, 10, 20, true, true)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	wire.SetFontSize(55)

	s := wire.GridBox(800, 400, 400, 20, 10, 10, false)
	s.Subdivide(5)
	s.RotateZ(pi / 2)
	s.Points.TwistY(blmath.LoopSin(percent, -0.0034, 0.0034))
	s.Rotate(-percent*tau, -percent*2*tau, 0)

	t := wire.NewString("do the twist").AsLine()
	t.Subdivide(5)
	t.TranslateY(-300)
	t.AddShape(t.RotatedX(pi))
	t.AddShape(t.RotatedX(pi / 2).RotatedZ(pi))
	t.RotateZ(pi / 2)
	t.Points.TwistY(blmath.LoopSin(percent, -0.0034, 0.0034))
	t.Rotate(-percent*tau, -percent*2*tau, 0)

	s.Stroke(2)
	t.Stroke(5)

	if blur {
		context.GaussianBlur(20)
		s.Stroke(1)
		t.Stroke(2.5)
	}
}
