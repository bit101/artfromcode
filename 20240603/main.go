// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240603"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 240)
		program.AddSceneWithFrames(scene2, 60)
		program.AddSceneWithFrames(scene3, 240)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau = blmath.Tau
	pi  = math.Pi
)

var (
	blur = true
	lat  = 16
	lon  = 23
)

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := wire.Sphere(300, lat, lon, true, true)
	s0 := s.Clone()
	s0.Cull(func(p *wire.Point) bool {
		return p.X > -1
	})
	s0.TranslateY(-300)
	angle := easing.CubicEaseOut(percent, 0, pi/4)
	s0.RotateZ(angle)
	s0.TranslateY(300)
	s1 := s0.ScaledX(-1)
	s0.AddShape(s1)
	s0.RotateY(percent * tau * 2)

	random.Seed(0)
	cloud := wire.RandomInnerSphere(300, 10000)
	cloud.RotateY(-percent * tau)
	random.Seed(1)
	for _, p := range cloud.Points {
		p.Y = easing.SineEaseIn(percent, p.Y, p.Y-random.FloatRange(800, 3000))
	}

	// s0.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	s0.Stroke()
	cloud.RenderPoints(2)

	//===========================
	if blur {
		context.GaussianBlur(20)
		context.SetLineWidth(1)
		s0.Stroke()
		cloud.RenderPoints(1)
	}
	//===========================
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	random.Seed(0)

	s := wire.Sphere(300, lat, lon, true, true)
	s0 := s.Clone()
	s0.Cull(func(p *wire.Point) bool {
		return p.X > -1
	})
	s0.TranslateY(-300)
	angle := easing.SineEaseIn(percent, pi/4, 0)
	s0.RotateZ(angle)
	s0.TranslateY(300)
	s1 := s0.ScaledX(-1)
	s0.AddShape(s1)
	s0.RotateY(percent * pi)

	// s0.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	s0.Stroke()

	//===========================
	if blur {
		context.GaussianBlur(20)
		context.SetLineWidth(1)
		s0.Stroke()
	}
	//===========================
}

func scene3(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	random.Seed(0)

	s := wire.Sphere(300, lat, lon, true, true)
	s.RotateY(percent * tau * 2)
	if percent < 0.05 {
		p := 1 - percent*20
		s.Randomize(10 * p)
	}

	random.Seed(0)
	cloud := wire.RandomInnerSphere(300, int(percent*10000))
	cloud.RotateY(-percent * tau)

	context.SetLineWidth(2)
	s.Stroke()
	cloud.RenderPoints(2)

	//===========================
	if blur {
		context.GaussianBlur(20)
		context.SetLineWidth(1)
		s.Stroke()
		cloud.RenderPoints(1)
	}
	//===========================
}
