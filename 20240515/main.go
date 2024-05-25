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
		program.AddSceneWithFrames(scene1, 360)
		program.RenderAndPlayVideo("out/frames", "out/out.mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 120, 800)

	shape := wire.Sphere(300, 20, 40, true, true)

	x := blmath.LoopSin(percent, 0, 600)
	shape.Rotate(percent*blmath.Tau, percent*2*blmath.Tau, 0)
	shape.TranslateY(x)
	for _, p := range shape.Points {
		if p.Y > 300 {
			p.Y = 600 - p.Y
		}
	}

	floor := wire.NewShape()
	for range 10000 {
		floor.AddPoint(wire.RandomPointInCircle(2000))
	}
	floor.TranslateY(300)
	floor.RotateY(-percent * blmath.Tau)

	t := wire.NewString("impenetrable", wire.FontAsteroid).AsCylinder(1000, 40)
	t.UniScale(0.5)
	t.TranslateY(250)
	t.RotateY(percent * blmath.Tau)

	context.SetLineWidth(2)
	shape.Stroke()
	floor.RenderPoints(2)
	t.Stroke()

	context.GaussianBlur(20)
	context.SetLineWidth(0.5)
	shape.Stroke()
	floor.RenderPoints(1)
	t.Stroke()
}
