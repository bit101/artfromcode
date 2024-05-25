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

//revive:disable:unused-parameter
const (
	tau = blmath.Tau
	pi  = math.Pi
)

func main() {
	renderTarget := target.Video

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/out.png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 720)
		program.RenderAndPlayVideo("out/frames", "out/out.mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	c0 := wire.Spring(300, 100, 0, 10, 20)
	c0.TranslateY(blmath.LoopSin(percent*2, -300, -600))

	c1 := c0.RotatedZ(pi / 2)
	c2 := c1.RotatedZ(pi / 2)
	c3 := c2.RotatedZ(pi / 2)
	c4 := c3.RotatedY(pi / 2)
	c5 := c4.RotatedY(pi)
	s := wire.Sphere(170, 12, 24, true, true)

	shape := wire.NewShape()
	shape.AddShape(c0)
	shape.AddShape(c1)
	shape.AddShape(c2)
	shape.AddShape(c3)
	shape.AddShape(c4)
	shape.AddShape(c5)
	shape.AddShape(s)

	txt := wire.NewString("satellite", wire.FontArcade).AsCylinder(blmath.LoopSin(percent, 1000, 600), 40)
	txt.UniScale(0.5)

	txt.RotateY(percent * tau * 2)
	txt.RotateZ(-pi / 4)
	txt.Rotate(percent*tau, percent*2*tau, 0)
	shape.Rotate(percent*tau, percent*2*tau, 0)

	context.SetLineWidth(2)
	shape.Stroke()
	txt.Stroke()

	context.GaussianBlur(20)
	context.SetLineWidth(1)
	shape.Stroke()
	txt.Stroke()

	context.Starfield(200, 1)

}
