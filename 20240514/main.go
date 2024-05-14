// Package main renders an image, gif or video
package main

import (
	"math"

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
	s0 := wire.Sphere(500, 40, 80, true, true)
	s0.Rotate(percent*blmath.Tau, percent*2*blmath.Tau, 0)

	s0.Translate(-400, 0, 400)
	p := wire.NewPoint(blmath.LoopSin(percent, 600, -400), 0, blmath.LoopSin(percent, -200, 400))
	s0.Points.Push(p, 300)

	s1 := wire.Sphere(200, 10, 20, true, true)
	s1.Rotate(percent*blmath.Tau, percent*2*blmath.Tau, 0)
	s1.Translate(p.X, p.Y, p.Z)

	t := wire.NewString("assimilation", wire.FontAsteroid).AsLine(20)
	// t.Translate(blmath.LoopSin(percent, 400, -600), -300, blmath.LoopSin(percent, -200, 400))
	t.TranslateX(blmath.LoopSin(percent, 600, -200))
	t.RotateY(math.Pi / 4)
	t.Translate(0, -500, 500)

	context.SetLineWidth(1)
	s0.Stroke()
	context.SetLineWidth(2)
	s1.Stroke()
	t.Stroke()

	context.GaussianBlur(20)
	context.SetLineWidth(0.5)
	s0.Stroke()
	s1.Stroke()
	context.SetLineWidth(0.75)
	t.Stroke()

}
