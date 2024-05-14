// Package main renders an image, gif or video
package main

import (
	"github.com/bit101/bitlib/bllog"
	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

//revive:disable:unused-parameter

func main() {
	bllog.InitProjectLog("project.log")
	defer bllog.CloseLog()

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
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.Box(200, 200, 200)
	shape.Rotate(percent*blmath.Tau, percent*2*blmath.Tau, percent*blmath.Tau)
	shape.Translate(200, -160, -100)
	shape2 := shape.ScaledY(-1)

	f := wire.GridPlane(2000, 2000, 20, 20)

	t := wire.NewString("reflection", wire.FontAsteroid).AsLine(20)
	t.Translate(120, -100, 400)
	// t.RotateY(-math.Pi / 4)
	t2 := t.ScaledY(-1)

	shape.RotateY(-0.5)
	shape2.RotateY(-0.5)
	f.RotateY(-0.5)
	t.RotateY(-0.5)
	t2.RotateY(-0.5)

	shape.RotateX(blmath.LoopSin(percent, -0.1, -1))
	shape2.RotateX(blmath.LoopSin(percent, -0.1, -1))
	f.RotateX(blmath.LoopSin(percent, -0.1, -1))
	t.RotateX(blmath.LoopSin(percent, -0.1, -1))
	t2.RotateX(blmath.LoopSin(percent, -0.1, -1))

	context.SetLineWidth(5)
	shape.Stroke()
	shape2.Stroke()
	t.Stroke()
	t2.Stroke()
	context.SetLineWidth(1)
	f.Stroke()

	context.GaussianBlur(14)
	context.SetLineWidth(0.5)
	shape.Stroke()
	context.SetLineWidth(0.75)
	t.Stroke()
	context.SetLineWidth(0.25)
	f.Stroke()

}
