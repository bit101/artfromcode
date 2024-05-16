// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/bllog"
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
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
		program.AddSceneWithFrames(scene1, 420)
		program.RenderAndPlayVideo("out/frames", "out/out.mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	to1 := wire.Torus(300, 150, math.Pi, 24, 24, true, true)
	to1.Rotate(percent*blmath.Tau, percent*2*blmath.Tau, 0)

	to2 := wire.NewShape()
	for range 10000 {
		to2.AddRandomPointInTorus(290, 150, math.Pi)
	}
	to2.RotateY(percent * blmath.Tau)
	to2.Rotate(percent*blmath.Tau+math.Pi, percent*2*blmath.Tau, 0)

	txt := wire.NewString("emerge", wire.FontAsteroid).AsLine(20)
	txt.Translate(100, blmath.LoopSin(percent-0.25, 1000, 500), 200)
	txt.RotateY(blmath.LoopSin(percent, -0.5, 0.5))

	txt2 := wire.NewString("return", wire.FontAsteroid).AsLine(20)
	txt2.Translate(-100, blmath.LoopSin(percent+0.25, -1000, -500), 200)
	txt2.RotateY(blmath.LoopSin(percent, 0.5, -0.5))

	context.SetLineWidth(2)
	to1.Stroke()
	to2.RenderPoints(2)
	context.SetLineWidth(3)
	txt.Stroke()
	txt2.Stroke()

	context.GaussianBlur(20)
	context.SetLineWidth(0.5)
	to1.Stroke()
	to2.RenderPoints(1.5)
	// context.SetLineWidth(0.75)
	txt.Stroke()
	txt2.Stroke()
}
