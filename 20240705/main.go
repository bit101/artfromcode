// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
)

func main() {
	renderTarget := target.Video
	fileName := "20240705"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 180)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau  = blmath.Tau
	pi   = math.Pi
	blur = false
)

func scene1(context *cairo.Context, width, height, percent float64) {
	random.Seed(0)
	context.WhiteOnBlack()
	context.Save()
	context.TranslateCenter()
	context.Rotate(-pi / 2)
	context.Translate(-width/2, -height/2)

	a := percent * tau
	xc := 0.5 + math.Cos(a)*0.25
	yc := 0.5 + math.Sin(a)*0.25
	iter(context, 0, 0, width, height, xc, yc, 5)

	context.Restore()
}

func iter(context *cairo.Context, x, y, w, h, xc, yc float64, depth int) {
	rand := 0.25
	xc += random.FloatRange(-rand, rand)
	yc += random.FloatRange(-rand, rand)
	if depth > 0 {
		iter(context, x, y, w*xc, h*yc, xc, yc, depth-1)
		iter(context, x+w*xc, y, w-w*xc, h*yc, xc, yc, depth-1)
		iter(context, x, y+h*yc, w*xc, h-h*yc, xc, yc, depth-1)
		iter(context, x+w*xc, y+h*yc, w-w*xc, h-h*yc, xc, yc, depth-1)
	} else {
		context.SetSourceWhite()
		context.FillRectangle(x, y, w*xc, h*yc)
		context.FillRectangle(x+w*xc, y+h*yc, w-w*xc, h-h*yc)
	}

}
