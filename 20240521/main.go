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
	tau  = blmath.Tau
	pi   = math.Pi
	blur = 20.0
)

func main() {
	renderTarget := target.Video
	fileName := "20240521"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 180)
		program.AddSceneWithFrames(scene2, 180)
		program.AddSceneWithFrames(scene3, 180)
		program.AddSceneWithFrames(scene4, 180)
		program.AddSceneWithFrames(scene5, 180)
		program.AddSceneWithFrames(scene6, 180)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

// grow
func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.RandomInnerSphere(400, 5000)
	shape.UniScale(percent * 1)
	shape.Rotate(percent*0.5*tau, percent*tau, 0)

	shape.RenderPoints(2)

	if blur > 0 {
		context.GaussianBlur(20)
		shape.RenderPoints(0.5)
	}
}

// hollow out
func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.RandomInnerSphere(400, 5000)
	shape.Points.Push(wire.NewPoint(0, 0, 0), percent*400)
	shape.Rotate((percent*0.5+0.5)*tau, percent*tau, 0)

	shape.RenderPoints(2)

	if blur > 0 {
		context.GaussianBlur(20)
		shape.RenderPoints(0.5)
	}
}

// order points
func scene3(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.RandomInnerSphere(400, 5000)
	shape.Points.Push(wire.NewPoint(0, 0, 0), 400)
	shape.Rotate(percent*0.5*tau, percent*tau, 0)

	shape2 := wire.Sphere(400, 12, 24, false, false)
	shape2.Rotate(percent*0.5*tau, percent*tau, 0)

	shape.RenderPoints(2 - percent*2)
	shape2.RenderPoints(1 + percent*3) // 1-4

	if blur > 0 {
		context.GaussianBlur(20)
		shape.RenderPoints(0.5 - percent*0.5)
		shape2.RenderPoints(0.5 + percent*0.5) // 0.5-1
	}
}

// latitude
func scene4(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.Sphere(400, 12, 24, false, false)
	shape.Rotate((percent*0.5+0.5)*tau, percent*tau, 0)

	shape2 := wire.Sphere(400, 12, 24, true, false)
	shape2.Rotate((percent*0.5+0.5)*tau, percent*tau, 0)

	shape.RenderPoints(4 - percent*4) // 4-0
	shape.Stroke()

	context.SetLineWidth(percent * 2)
	shape2.Stroke()

	if blur > 0 {
		context.GaussianBlur(20)
		shape.RenderPoints(1 - percent) // 1-0
		context.SetLineWidth(percent * 1)
		shape2.Stroke()
	}
}

// longitude
func scene5(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.Sphere(400, 12, 24, true, false)
	shape.Rotate(percent*0.5*tau, percent*tau, 0)

	shape2 := wire.Sphere(400, 12, 24, false, true)
	shape2.Rotate(percent*0.5*tau, percent*tau, 0)

	context.SetLineWidth(2)
	shape.Stroke()

	context.SetLineWidth(percent * 2)
	shape2.Stroke()

	if blur > 0 {
		context.GaussianBlur(20)
		context.SetLineWidth(1)
		shape.Stroke()
		context.SetLineWidth(percent)
		shape2.Stroke()
	}
}

// shrink
func scene6(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(0)
	wire.InitWorld(context, 200, 200, 800)

	shape := wire.Sphere(400, 12, 24, true, true)
	shape.UniScale(1 - percent)
	shape.Rotate((percent*0.5+0.5)*tau, percent*tau, 0)

	context.SetLineWidth(2)
	shape.Stroke()

	if blur > 0 {
		context.GaussianBlur(20)
		context.SetLineWidth(1)
		shape.Stroke()
	}
}
