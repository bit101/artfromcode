// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/lures"
	"github.com/bit101/lures/l3d"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240613"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 180)
		program.AddSceneWithFrames(scene2, 180)
		program.AddSceneWithFrames(scene3, 180)
		program.AddSceneWithFrames(scene4, 180)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau   = blmath.Tau
	pi    = math.Pi
	blur  = true
	count = 20000
)

var (
	attr lures.Lure3d
	s    *wire.Shape
	box  *wire.Shape
	box2 *wire.Shape
)

func init() {
	attr = l3d.NewChenLee()
	s = wire.ShapeFromLure(attr, count)
	// box = wire.Box(80, 700, 900)
	box = wire.GridPlane(700, 1000, 20, 20*10/7)
	box.Subdivide(5)
	box.RotateZ(pi / 2)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	x := 800.0 * (1 - percent)
	box1 := box.Clone()
	box2 := box.Clone()
	box1.TranslateX(-x - 40)
	box2.TranslateX(x + 40)

	s1 := s.Clone()
	s1.RotateX(-pi / 2)
	for _, p := range s1.Points {
		if p.X > x {
			p.X = x
		}
		if p.X < -x {
			p.X = -x
		}
	}

	s1.RotateY(percent * pi)
	box1.RotateY(percent * pi)
	box2.RotateY(percent * pi)

	// s1.Rotate(percent*tau, percent*tau*2, 0)
	s1.RenderPoints(3)
	box1.Stroke(2)
	box2.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s1.RenderPoints(1.5)
		box1.Stroke(1)
		box2.Stroke(1)
	}
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	x := 0.0
	box1 := box.Clone()
	box2 := box.Clone()
	box1.TranslateX(-x - 40)
	box2.TranslateX(x + 40)

	s1 := s.Clone()
	s1.RotateX(-pi / 2)
	for _, p := range s1.Points {
		if p.X > x {
			p.X = x
		}
		if p.X < -x {
			p.X = -x
		}
	}

	s1.RotateY(pi + percent*pi)
	box1.RotateY(pi + percent*pi)
	box2.RotateY(pi + percent*pi)

	// s1.Rotate(percent*tau, percent*tau*2, 0)
	s1.RenderPoints(3)
	box1.Stroke(2)
	box2.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s1.RenderPoints(1.5)
		box1.Stroke(1)
		box2.Stroke(1)
	}
}

func scene3(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	x := 800.0 * (percent)
	box1 := box.Clone()
	box2 := box.Clone()
	box1.TranslateX(-x - 40)
	box2.TranslateX(x + 40)

	s1 := s.Clone()
	s1.RotateX(-pi / 2)
	for _, p := range s1.Points {
		p.X = 0
	}

	s1.RotateY(percent * pi)
	box1.RotateY(percent * pi)
	box2.RotateY(percent * pi)

	// s1.Rotate(percent*tau, percent*tau*2, 0)
	s1.RenderPoints(3)
	box1.Stroke(2)
	box2.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s1.RenderPoints(1.5)
		box1.Stroke(1)
		box2.Stroke(1)
	}
}

func scene4(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	x := 800.0
	box1 := box.Clone()
	box2 := box.Clone()
	box1.TranslateX(-x - 40)
	box2.TranslateX(x + 40)

	s1 := s.Clone()
	s2 := s.Clone()
	s1.RotateX(-pi / 2)
	s2.RotateX(-pi / 2)
	for _, p := range s1.Points {
		p.X = 0
	}

	s1.TranslateY(easing.CubicEaseInOut(percent, 1, 1100))
	s2.TranslateY(easing.CubicEaseInOut(percent, -1100, 0))

	s1.RotateY(pi + percent*pi)
	s2.RotateY(pi + percent*pi)
	box1.RotateY(pi + percent*pi)
	box2.RotateY(pi + percent*pi)

	// s1.Rotate(percent*tau, percent*tau*2, 0)
	s1.RenderPoints(3)
	s2.RenderPoints(3)
	box1.Stroke(2)
	box2.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s1.RenderPoints(1.5)
		s2.RenderPoints(1.5)
		box1.Stroke(1)
		box2.Stroke(1)
	}
}
