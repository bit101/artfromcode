// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/lures/l3d"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240614"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 300)
		program.AddSceneWithFrames(scene2, 300)
		program.AddSceneWithFrames(scene3, 300)
		program.AddSceneWithFrames(scene4, 300)
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
	earth *wire.Shape
	attr  *wire.Shape
)

func init() {
	earth, _ = wire.LoadShape("earth10000.wire")
	lure := l3d.NewThomas()
	attr = wire.ShapeFromLure(lure, 10000)
	// earth = makeEarth(10000)
	// earth.Save("earth10000.wire")
}

func makeEarth(count int) *wire.Shape {
	mapSurface, _ := cairo.NewSurfaceFromPNG("earth2.png")
	data, _ := mapSurface.GetData()
	earth := wire.NewShape()
	for true {
		w := float64(mapSurface.GetWidth())
		h := float64(mapSurface.GetHeight())
		x := random.FloatRange(0, w)
		y := random.FloatRange(0, h)
		r, _, _, _ := mapSurface.GetPixel(data, int(x), int(y))
		rad := 100.0

		if r < 128 {
			u := blmath.Map(y, 0, h, -1, 1)
			t := blmath.Map(x, 0, w, 0, tau)
			sx := math.Sqrt(1-u*u) * math.Cos(t)
			sz := math.Sqrt(1-u*u) * math.Sin(t)
			sy := u
			earth.AddXYZ(sx*rad, sy*rad, sz*rad)
		}
		if len(earth.Points) == count {
			break
		}
	}
	return earth
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	wire.SetFog(true, 700, 900)
	random.Seed(0)
	s := earth.UniScaled(4)
	s.RotateY(-percent * tau)
	s.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		s.RenderPoints(1)
	}
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	wire.SetFog(true, 700, 900)
	random.Seed(0)
	s := earth.UniScaled(4)
	s.Points.Lerp(percent, attr.Points)
	s.RotateY(-percent * tau)
	s.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		s.RenderPoints(1)
	}
}

func scene3(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	wire.SetFog(true, 700, 900)
	random.Seed(0)
	s := attr.Clone()
	s.RotateY(-percent * tau)
	s.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		s.RenderPoints(1)
	}
}

func scene4(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	wire.SetFog(true, 700, 900)
	random.Seed(0)
	s := earth.UniScaled(4)
	s.Points.Lerp(1-percent, attr.Points)
	s.RotateY(-percent * tau)
	s.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		s.RenderPoints(1)
	}
}
