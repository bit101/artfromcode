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
	fileName := "20240529"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 600)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	context.SetSourceHSV(percent*360, 0.125, 1)
	wire.InitWorld(context, 200, 200, 800)
	shape := wire.NewShape()

	random.Seed(0)
	x := random.FloatRange(-0.1, 0.1)
	y := random.FloatRange(-0.1, 0.1)
	z := random.FloatRange(-0.1, 0.1)

	a := blmath.LoopSin(percent, 1.5, 2.9)
	b := blmath.LoopSin(percent+0.125, 1.5, 2.8)

	for range 10000 {
		x, y, z = sprott(x, y, z, a, b)
		shape.AddXYZ(x, y, z)
	}
	shape.TranslateX(-0.75)
	shape.Scale(600, 600, 1200)
	shape.Rotate(percent*tau, percent*2*tau, 0)

	s := wire.Sphere(80, 8, 16, true, true)
	s.Rotate(percent*tau, percent*2*tau, 0)

	shape.RenderPoints(2)
	context.SetLineWidth(2)
	s.Stroke()

	context.GaussianBlur(20)
	shape.RenderPoints(1)
	context.SetLineWidth(0.5)
	s.Stroke()

}

func sprott(x, y, z, a, b float64) (float64, float64, float64) {
	x1 := y + a*x*y + x*z
	y1 := 1 - b*x*x + y*z
	z1 := x - x*x - y*y
	dt := 0.01
	return x + x1*dt, y + y1*dt, z + z1*dt
}
