// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240616"

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
	s *wire.Shape
)

func init() {
	s = wire.GridBox(600, 600, 600, 16, 16, 16, false)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	random.Seed(0)

	s1 := s.Clone()
	d := 400.0
	r := blmath.LoopSin(percent, 100, 280)
	s1.Points.Push(wire.NewPoint(0, 0, -d), r)
	s1.Points.Push(wire.NewPoint(0, 0, d), r)
	s1.Points.Push(wire.NewPoint(0, -d, 0), r)
	s1.Points.Push(wire.NewPoint(0, d, 0), r)
	s1.Points.Push(wire.NewPoint(-d, 0, 0), r)
	s1.Points.Push(wire.NewPoint(d, 0, 0), r)
	s1.Rotate(percent*tau, percent*tau*2, 0)
	s1.TranslateZ(blmath.LoopSin(percent*2, 0, -600))

	t := wire.NewString("hello").AsLine()
	wire.SetFontSize(10)
	t.TranslateZ(blmath.LoopSin(percent*2, 0, -600))

	s1.Stroke(2)
	t.Stroke(1.5)

	if blur {
		context.GaussianBlur(20)
		// s1.RenderPoints(1)
		s1.Stroke(1)
		t.Stroke(0.75)
	}
}
