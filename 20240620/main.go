// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240620"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 360)
		program.AddSceneWithFrames(scene2, 360)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau  = blmath.Tau
	pi   = math.Pi
	blur = true
)

var (
	shape *wire.Shape
)

func init() {
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	wire.SetFontSize(50)

	s := blmath.Lerp(percent, 2000, 4000)
	a := wire.GridBox(s, s, s, 10, 10, 10, false)
	a.Rotate(percent*tau, percent*tau*2, 0)

	b := wire.Sphere(blmath.Lerp(percent, 100, 1000), 20, 40, true, true)
	b.Rotate(-percent*tau, -percent*tau*2, 0)

	s = blmath.Lerp(percent, 0, 200)
	c := wire.GridBox(s, s, s, 10, 10, 10, false)
	c.Rotate(percent*tau, percent*tau*2, 0)

	t := wire.NewString("zoom").AsLine()
	t.UniScale(blmath.Lerp(percent, 1, 10))
	t.TranslateY(blmath.Lerp(percent, -180, -1800))

	t2 := wire.NewString("enhance").AsLine()
	t2.UniScale(blmath.Lerp(percent, 0, 1))
	t2.TranslateY(blmath.Lerp(percent, 0, 200))

	a.Stroke(1 - percent)
	b.Stroke(2 - percent)
	c.Stroke(2)

	t.Stroke(2 - percent)
	t2.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		a.Stroke((1 - percent) / 2)
		b.Stroke((2 - percent) / 2)
		c.Stroke(1)

		t.Stroke((2 - percent) / 2)
		t2.Stroke(1)
	}
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	wire.SetFontSize(50)

	a := wire.Sphere(blmath.Lerp(percent, 1000, 2000), 20, 40, true, true)
	a.Rotate(-percent*tau, -percent*tau*2, 0)

	s := blmath.Lerp(percent, 200, 2000)
	b := wire.GridBox(s, s, s, 10, 10, 10, false)
	b.Rotate(percent*tau, percent*tau*2, 0)

	c := wire.Sphere(blmath.Lerp(percent, 0, 100), 20, 40, true, true)
	c.Rotate(-percent*tau, -percent*tau*2, 0)

	t := wire.NewString("enhance").AsLine()
	t.UniScale(blmath.Lerp(percent, 1, 10))
	t.TranslateY(blmath.Lerp(percent, 200, 2000))

	t2 := wire.NewString("zoom").AsLine()
	t2.UniScale(blmath.Lerp(percent, 0, 1))
	t2.TranslateY(blmath.Lerp(percent, 0, -180))

	a.Stroke(1 - percent)
	b.Stroke(2 - percent)
	c.Stroke(2)

	t.Stroke(2 - percent)
	t2.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		a.Stroke((1 - percent) / 2)
		b.Stroke((2 - percent) / 2)
		c.Stroke(1)

		t.Stroke((2 - percent) / 2)
		t2.Stroke(1)
	}
}
