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
	fileName := "20240618"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 420)
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
	shape    *wire.Shape
	shape2   *wire.Shape
	rotation = 0.0
	sep      = 1.0
)

func init() {
	shape = wire.GridBox(400, 400, 400, 20, 20, 20, false)
	shape2 = wire.GridBox(100, 100, 100, 10, 10, 10, false)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	random.Seed(0)
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, blmath.LoopSin(percent, 1500, 170))

	rotation = percent * pi
	sep = blmath.LoopSin(percent*4, 1, 1.5)
	s := shape.Clone()
	scale := 0.5
	iterate(s, scale, 400.0*scale, 4, 20/2, 3)
	s.Rotate(percent*tau, percent*2*tau, 0)

	s2 := shape2.Clone()
	s2.Rotate(-percent*tau, -percent*2*tau, 0)

	s.Stroke(1.2)
	s2.Stroke(1)

	if blur {
		context.GaussianBlur(4)
		s.Stroke(0.6)
		s2.Stroke(0.5)
	}
}

func iterate(shape *wire.Shape, scale, size, count float64, lines, depth int) {
	s := wire.GridBox(size, size, size, lines, lines, lines, false)
	for i := 0.0; i < count; i++ {
		a := pi / 2 * (i - 1)
		s1 := s.Clone()
		if depth > 0 {
			iterate(s1, scale, size*scale, 3, lines/2, depth-1)
		}
		s1.TranslateX(size * sep * 1.5)
		s1.RotateX(rotation)
		s1.RotateY(a)
		shape.AddShape(s1)
	}
	for i := 0.0; i < 2; i++ {
		a := pi*i + pi/2
		s1 := s.Clone()
		if depth > 0 {
			iterate(s1, scale, size*scale, 3, lines/2, depth-1)
		}
		s1.TranslateX(size * sep * 1.5)
		s1.RotateX(rotation)
		s1.RotateZ(a)
		shape.AddShape(s1)
	}
}
