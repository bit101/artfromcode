// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/noise"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
)

func main() {
	renderTarget := target.Video
	fileName := "20240706"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 360)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau  = blmath.Tau
	pi   = math.Pi
	blur = true
)

func scene1(context *cairo.Context, width, height, percent float64) {
	random.Seed(0)
	context.BlackOnWhite()
	context.Save()
	context.TranslateCenter()
	context.Rotate(-pi / 2)
	context.Translate(-width/2, -height/2)

	a := percent * tau
	xc := 0.5 + math.Cos(a)*0.25
	yc := 0.5 + math.Sin(a)*0.25
	iter(context, -100, -100, width+200, height+200, xc, yc, 6)

	context.FillPreserve()
	if blur {
		context.GaussianBlur(2)
		// iter(context, -100, -100, width+200, height+200, xc, yc, 6)
		// context.SetSourceBlack()
		context.Fill()
	}

	context.Restore()
}

func iter(context *cairo.Context, x, y, w, h, xc, yc float64, depth int) {
	// rand := 0.25
	scale := 0.002
	n := noise.Simplex2(x*scale, y*scale)
	a := n * tau
	xc += math.Cos(a) * 0.05
	yc += math.Sin(a) * 0.05
	// xc += random.FloatRange(-rand, rand)
	// yc += random.FloatRange(-rand, rand)

	if depth > 0 {
		iter(context, x, y, w*xc, h*yc, xc, yc, depth-1)
		iter(context, x+w*xc, y, w-w*xc, h*yc, xc, yc, depth-1)
		iter(context, x, y+h*yc, w*xc, h-h*yc, xc, yc, depth-1)
		iter(context, x+w*xc, y+h*yc, w-w*xc, h-h*yc, xc, yc, depth-1)
	} else {
		context.SetSourceBlack()
		context.Rectangle(x, y, w*xc, h*yc)
		context.Rectangle(x+w*xc, y+h*yc, w-w*xc, h-h*yc)
	}

}
