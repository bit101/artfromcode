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
	fileName := "20240528"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 360)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	shape := wire.NewShape()

	random.Seed(0)
	x := random.FloatRange(-0.1, 0.1)
	y := random.FloatRange(-0.1, 0.1)
	z := random.FloatRange(-0.1, 0.1)

	for range 10000 {
		x, y, z = lorenz(x, y, z)
		shape.AddXYZ(x, y, z)
	}
	shape.TranslateZ(-20)
	shape.UniScale(20)
	shape.Rotate(percent*tau, percent*2*tau, 0)

	txt := wire.NewString("lorenz", wire.FontAsteroid).AsLine(20)
	txt.RotateZ(pi)
	txt.Rotate(-0.5, 1.25, 0)
	txt.TranslateX(-300)
	txt.Rotate(percent*tau, percent*2*tau, 0)

	shape.RenderPoints(2)
	context.SetLineWidth(2)
	txt.Stroke()

	context.GaussianBlur(20)
	shape.RenderPoints(1)
	context.SetLineWidth(1)
	txt.Stroke()

}

func lorenz(x, y, z float64) (float64, float64, float64) {
	a := 28.0
	b := 10.0
	c := 8.0 / 3.0
	x1 := b * (y - x)
	y1 := x*(a-z) - y
	z1 := x*y - c*z
	dt := 0.01
	return x + x1*dt, y + y1*dt, z + z1*dt
}
