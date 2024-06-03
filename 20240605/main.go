// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/easing"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240605"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 90)
		program.AddSceneWithFrames(scene2, 180)
		program.AddSceneWithFrames(scene3, 90)
		program.AddSceneWithFrames(scene4, 360)
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
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	random.Seed(0)

	shape := wire.Sphere(300, 20, 40, true, true)
	shape.Cull(func(p *wire.Point) bool {
		return p.Y < 1
	})

	wobble := percent * 0.5

	shape.TranslateY(easing.CubicEaseOut(percent, 0, -150))
	shape.AddShape(shape.ScaledY(-1))
	shape.RotateY(percent * tau / 2)
	shape.RotateX(blmath.LoopSin(percent, -wobble, wobble))

	cloud := wire.RandomInnerTorus(easing.CubicEaseOut(percent, 300, 450), percent*150, tau, 20000)
	cloud.RotateY(-percent * tau / 2)
	cloud.RotateX(blmath.LoopSin(percent, -wobble, wobble))

	shape.Stroke(2)
	cloud.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		shape.Stroke(1)
		cloud.RenderPoints(0.5)
	}
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	random.Seed(0)

	shape := wire.Sphere(300, 20, 40, true, true)
	shape.Cull(func(p *wire.Point) bool {
		return p.Y < 1
	})
	wobble := (1 - percent) * 0.5

	// shape.Rotate(percent*tau, percent*2*tau, 0)
	shape.TranslateY(-150)
	shape.AddShape(shape.ScaledY(-1))
	shape.RotateY(percent * tau)
	shape.RotateX(blmath.LoopSin(percent*2, -wobble, wobble))

	cloud := wire.RandomInnerTorus(450, 150, tau, 20000)
	cloud.RotateY(-percent*tau + pi)
	cloud.Randomize(100 * percent)
	for _, p := range cloud.Points {
		p.X += easing.CubicEaseIn(percent, 0, random.FloatRange(1200, 3200))
	}
	cloud.RotateX(blmath.LoopSin(percent*2, -wobble, wobble))

	shape.Stroke(2)
	cloud.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		shape.Stroke(1)
		cloud.RenderPoints(0.5)
	}
}
func scene3(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	random.Seed(0)

	shape := wire.Sphere(300, 20, 40, true, true)
	shape.Cull(func(p *wire.Point) bool {
		return p.Y < 1
	})

	shape.TranslateY(easing.BounceEaseOut(percent, -150, 0))
	shape.AddShape(shape.ScaledY(-1))
	shape.RotateY(percent * tau / 2)

	shape.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		shape.Stroke(1)
	}
}

func scene4(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	random.Seed(0)

	shape := wire.Sphere(300, 20, 40, true, true)
	shape.Randomize(blmath.LoopSin(percent, 0, 20))
	shape.Rotate(percent*tau*2, easing.CubicEaseInOut(percent, 0, 4*tau), 0)
	col := blcolor.HSV(percent*360*3, blmath.LoopSin(percent, 0, 0.25), 1)
	wire.SetRGB(col.R, col.G, col.B)

	shape.Stroke(2)

	if blur {
		context.GaussianBlur(20)
		shape.Stroke(1)
	}
}
