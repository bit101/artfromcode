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
	dots := wire.GridPlane(800, 800, 50, 50)
	angle := percent * blmath.Tau
	yrot := blmath.LoopSin(percent, 0.6, 0.4)
	p := wire.NewPoint(math.Sin(angle)*200, blmath.LoopSin(percent, -200, 200), math.Cos(angle)*200)
	dots.Points.Push(p, 200)
	dots.Rotate(-0.5, yrot, 0)

	s := wire.Sphere(150, 10, 20, true, true)
	s.Rotate(percent*blmath.Tau, percent*2*blmath.Tau, 0)
	s.Translate(p.X, p.Y, p.Z)
	s.Rotate(-0.5, yrot, 0)

	t := wire.NewString("no barriers", wire.FontAsteroid).AsLine(50)
	t.UniScale(0.5)
	t.Translate(0, -60, 400)
	t.Rotate(-0.5, yrot, 0)

	context.SetLineWidth(1)
	dots.RenderPoints(1.5)
	dots.Stroke()
	s.Stroke()
	t.Stroke()

	context.GaussianBlur(10)
	context.SetLineWidth(0.5)
	dots.RenderPoints(1)
	dots.Stroke()
	s.Stroke()

	context.SetLineWidth(1)
	t.Stroke()
}
