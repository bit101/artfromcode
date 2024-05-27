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
	fileName := "20240601"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 360)
		program.AddSceneWithFrames(scene2, 360)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

var shape *wire.Shape

func init() {
	shape = wire.Sphere(400, 20, 40, true, true)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := wire.NewShape()
	s0 := shape.Culled(func(p *wire.Point) bool {
		return p.X >= -1
	})
	s0.TranslateX(percent * 1000)
	s.AddShape(s0)
	s1 := shape.Culled(func(p *wire.Point) bool {
		return p.X <= 1
	})
	s1.TranslateX(percent * -1000)
	s.AddShape(s1)

	s3 := shape.Clone()
	s3.UniScale(percent)

	s.Rotate(percent*tau, percent*2*tau, 0)
	s3.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	s.Stroke()
	s3.Stroke()

	context.GaussianBlur(20)
	context.SetLineWidth(1)
	s.Stroke()
	s3.Stroke()
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	s := wire.NewShape()
	s0 := shape.Culled(func(p *wire.Point) bool {
		return p.Y >= -1
	})
	s0.TranslateY(percent * 1000)
	s.AddShape(s0)
	s1 := shape.Culled(func(p *wire.Point) bool {
		return p.Y <= 1
	})
	s1.TranslateY(percent * -1000)
	s.AddShape(s1)

	s3 := shape.Clone()
	s3.UniScale(percent)

	s.Rotate(percent*tau, percent*2*tau, 0)
	s3.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	s.Stroke()
	s3.Stroke()

	context.GaussianBlur(20)
	context.SetLineWidth(1)
	s.Stroke()
	s3.Stroke()
}
