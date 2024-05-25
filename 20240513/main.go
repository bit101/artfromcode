// Package main renders an image, gif or video
package main

import (
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

//revive:disable:unused-parameter

func main() {
	renderTarget := target.Video

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/out.png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 240)
		program.RenderAndPlayVideo("out/frames", "out/out.mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	random.Seed(0)
	shape := wire.NewShape()
	for range 5000 {
		shape.AddRandomPointInBox(500, 200, 200)
	}
	shape.Rotate(-0.3, 0.3, 0)
	shape.TranslateX(blmath.LoopSin(percent, -600, 600))
	shape.Points.Push(wire.NewPoint(0, 0, 0), 300)
	shape.RotateY(0.4)

	s := wire.Sphere(200, 10, 20, true, true)
	s.Rotate(percent*blmath.Tau, percent*2*blmath.Tau, 0)

	t := wire.NewString("untouchable", wire.FontAsteroid).AsLine(20)
	t.UniScale(0.8)
	t.Translate(0, blmath.LoopSin(percent*2+0.25, -400, -500), 300)
	t.RotateY(0.4)

	shape.RenderPoints(3)
	context.SetLineWidth(2)
	s.Stroke()
	context.SetLineWidth(4)
	t.Stroke()

	context.GaussianBlur(20)
	shape.RenderPoints(1)
	context.SetLineWidth(0.5)
	s.Stroke()
	context.SetLineWidth(1)
	t.Stroke()
}
