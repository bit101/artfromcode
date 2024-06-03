// Package main renders an image, gif or video
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

func main() {
	renderTarget := target.Video
	fileName := "20240604"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 240)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau = blmath.Tau
	pi  = math.Pi
)

var blur = true

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	random.Seed(0)

	c0 := wire.RandomInnerCylinder(800, 50, 2500)
	c0.TranslateY(-400)
	c0.AddShape(c0.TranslatedY(800))
	c0.TranslateY(-percent*800 + 200)
	for _, p := range c0.Points {
		a := blmath.Map(p.Y, 0, -800, 0, -tau*2)
		s := blmath.Map(p.Y, 200, -800, 0, 3)
		p.Translate(math.Cos(a)*50, 0, math.Sin(a)*50)
		p.Scale(s, 1, s)
	}
	c0.RotateY(-percent * tau)
	c0.Cull(func(p *wire.Point) bool {
		return p.Y < 200
	})
	c0.TranslateX(200)
	c0.AddShape(c0.TranslatedX(-400))
	c0.RotateY(percent * tau)

	c1 := wire.Cone(600, 20, 40, 30, 10, true, true)
	c1.Translate(200, 500, 0)
	c1.AddShape(c1.TranslatedX(-400))
	c1.RotateY(percent * tau)

	c0.RenderPoints(2)
	c1.Stroke(2)

	//===========================
	if blur {
		context.GaussianBlur(20)
		c0.RenderPoints(0.5)
		c1.Stroke(0.5)
	}
	//===========================
}
