// Package main renders an image or video
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
const (
	tau = blmath.Tau
	pi  = math.Pi
)

func main() {
	renderTarget := target.Video
	fileName := "20240526"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 400)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.NewShape()
	for i := range 10000 {
		a := float64(i) / 10000 * tau
		y := random.FloatRange(-100, 100) * 4
		r := 200.0
		p := wire.NewPoint(0, y, 0)
		p.RotateZ(a + percent*tau)
		p.TranslateX(r)
		p.RotateY(a)
		shape.AddPoint(p)
	}

	txt := wire.NewString("mobius   mobius", wire.FontArcade).AsCylinder(330, 10)

	shape.Rotate(percent*tau, percent*2*tau, 0)
	txt.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(4)
	txt.Stroke()
	shape.RenderPoints(2)

	context.GaussianBlur(20)
	context.SetLineWidth(2)
	txt.Stroke()
	shape.RenderPoints(1)
}
