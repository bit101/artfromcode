// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240622"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 180)
		program.AddSceneWithFrames(scene2, 60)
		program.AddSceneWithFrames(scene3, 120)
		program.AddSceneWithFrames(scene4, 90)
		program.AddSceneWithFrames(scene5, 240)
		program.AddSceneWithFrames(scene6, 90)
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
	shape = wire.GridBox(1600, 200, 200, 64, 8, 8, false)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := shape.Clone()
	s.RotateX(percent * pi)
	s.RotateY(easing.QuadraticEaseInOut(percent, 0, pi))

	s.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s.Stroke(1)
	}
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := shape.Clone()
	percent = easing.QuadraticEaseInOut(percent, 0, 1)
	s.TwistX(percent * tau)

	s.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s.Stroke(1)
	}
}

func scene3(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := shape.Clone()
	s.TwistX(tau)
	s.RotateX(percent * pi)
	s.RotateY(easing.QuadraticEaseInOut(percent, 0, pi))

	s.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s.Stroke(1)
	}
}

func scene4(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := shape.Clone()
	s.TwistX(tau)
	percent = easing.QuadraticEaseInOut(percent, 0, 1)
	s.WrapCylinderWithArc(percent * tau)
	s.TranslateY(percent * -200)

	s.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s.Stroke(1)
	}
}

func scene5(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := shape.Clone()
	s.TwistX(tau)
	s.WrapCylinderWithArc(tau)
	s.TranslateY(-200)

	percent = easing.QuadraticEaseInOut(percent, 0, 1)
	s.Rotate(percent*tau, percent*2*tau, 0)

	s.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s.Stroke(1)
	}
}

func scene6(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := shape.Clone()
	percent = easing.QuadraticEaseInOut(percent, 0, 1)
	s.TwistX(tau * (1 - percent))
	s.WrapCylinderWithArc(tau * (1 - percent))
	s.TranslateY(-200 * (1 - percent))

	// s.Rotate(percent*tau, percent*2*tau, 0)

	s.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		s.Stroke(1)
	}
}
