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
	fileName := "20240617"

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
	shape = wire.Sphere(200, 16, 32, true, true)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	percent += 0.5
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := shape.Clone()
	scale := blmath.LoopSin(percent, 0.0, 0.7)
	iterate(s, scale, 200.0*scale, 8, 5)
	s.Rotate(percent*tau, percent*2*tau, 0)

	s.Stroke(1.5)

	if blur {
		context.GaussianBlur(16)
		s.Stroke(0.75)
	}
}

func iterate(shape *wire.Shape, scale, radius float64, lines, depth int) {
	s := wire.Sphere(radius, lines, lines*2, true, true)
	for a := 0.0; a < tau; a += tau / 3 {
		s1 := s.Clone()
		if depth > 0 {
			iterate(s1, scale, radius*scale, lines-1, depth-1)
		}
		s1.TranslateX(radius/scale + radius)
		s1.RotateX(pi / 2)
		s1.RotateY(a)
		shape.AddShape(s1)
	}
}
