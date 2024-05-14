// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/bllog"
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
)

//revive:disable:unused-parameter

func main() {
	bllog.InitProjectLog("project.log")
	defer bllog.CloseLog()

	renderTarget := target.Video

	if renderTarget == target.Image {
		render.CreateAndViewImage(400, 400, "out/out.png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 360)
		program.RenderAndPlayVideo("out/frames", "out/out.mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	random.Seed(1)
	a := blmath.Tau * percent
	circles := geom.NewCircleList()
	for i := 0; i < 30; i++ {
		angle := a + random.FloatRange(0, math.Pi)
		if random.Boolean() {
			angle *= -1
		}
		x := random.FloatRange(0, width)
		y := random.FloatRange(0, height)
		r0 := random.FloatRange(20, 200)
		r1 := random.FloatRange(10, 50)
		circles.AddXY(x+math.Cos(angle)*r0, y+math.Sin(angle)*r0, r1)
	}

	metaRender(context, circles, 1.0)
	context.SetSourceBlack()
	metaRender(context, circles, 1.5)
	context.SetSourceWhite()
	metaRender(context, circles, 3.0)
	context.SetSourceBlack()
	metaRender(context, circles, 4.5)
	context.GaussianBlur(2)
}

func metaRender(context *cairo.Context, circles geom.CircleList, threshold float64) {
	res := 1.0
	width := context.Width
	height := context.Height

	for x := 0.0; x < width; x += res {
		for y := 0.0; y < height; y += res {
			val := 0.0
			for _, c := range circles {
				val += meta(c, x, y)
			}
			if val > threshold {
				context.FillRectangle(x, y, res, res)
			}
		}
	}

}

func meta(c *geom.Circle, x, y float64) float64 {
	dx := c.X - x
	dy := c.Y - y
	return c.Radius * c.Radius / (dx*dx + dy*dy)
}
