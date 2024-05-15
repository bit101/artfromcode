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
		render.CreateAndViewImage(400, 400, "out/out.png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 420)
		program.RenderAndPlayVideo("out/frames", "out/out.mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	wire.InitWorld(context, 200, 200, 800)
	wire.SetFog(true, 500, blmath.LoopSin(percent, 1500, 10000))

	txt := "sometimes we need the fog to remind ourselves that all of life is not black and white"
	author := "jonathan lockwood huie"
	t0 := wire.NewString(txt, wire.FontAsteroid).AsCylinder(2000, 20)
	t2 := wire.NewString(author, wire.FontAsteroid).AsLine(20)

	t0.RotateY(percent * blmath.Tau)
	t0.Translate(-1200, blmath.LoopSin(percent, -200, 200), 1500)
	t1 := t0.ScaledY(-1)
	t1.TranslateY(550)

	t2.UniScale(0.6)
	t2.Translate(-200, -600, 400)
	t2.RotateY(0.3)

	gp := wire.GridPlane(5000, 5000, 20, 20)
	gp.TranslateY(400)
	gp.RotateY(0.5)

	context.SetLineWidth(2)
	t0.Stroke()
	t1.Stroke()
	t2.Stroke()
	context.SetLineWidth(1)
	gp.Stroke()

	context.GaussianBlur(10)
	context.SetLineWidth(1)
	t0.Stroke()
	t2.Stroke()
	context.SetLineWidth(0.5)
	gp.Stroke()

}
