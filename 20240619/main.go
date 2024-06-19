// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/noise"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/lures/l3d"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240619"

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
		program.AddSceneWithFrames(scene7, 180)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau  = blmath.Tau
	pi   = math.Pi
	res  = 30.0
	blur = true
)

var (
	shape *wire.Shape
	obj   *wire.Shape
	skull *wire.Shape
)

func init() {
	shape = wire.NewShape()
	for i := -1000.0; i < 1000.0; i += res {
		shape.AddXYZ(i, 0, 0)
	}
	for j := 0; j < len(shape.Points)-1; j++ {
		shape.AddSegmentByIndex(j, j+1)
	}
	s := wire.NewShape()
	for z := res / 2; z < 2500; z += res {
		s2 := shape.TranslatedZ(z)
		for _, p := range s2.Points {
			scale := 0.004
			n := noise.Simplex2(p.X*scale, p.Z*scale) * 30
			p.Y = n

		}
		s.AddShape(s2)
	}
	s.AddShape(s.ScaledZ(-1))
	s.AddShape(s.TranslatedZ(5000))
	s.TranslateY(400)
	shape = s
	skull = wire.ShapeFromXYZ("skull.xyz")
	skull.ThinPoints(1, 29)
	skull.UniScale(3.5)
	skull.RotateX(pi)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	obj = wire.Sphere(200, 20, 40, true, true)
	scene(context, width, height, percent, false)
}

func scene2(context *cairo.Context, width, height, percent float64) {
	obj = wire.Cone(400, 0, 200, 20, 20, true, true)
	scene(context, width, height, percent, false)
}

func scene3(context *cairo.Context, width, height, percent float64) {
	obj = wire.Torus(150, 50, tau, 24, 20, true, true)
	scene(context, width, height, percent, false)
}

func scene4(context *cairo.Context, width, height, percent float64) {
	obj = wire.Cylinder(400, 200, 20, 20, true, true)
	scene(context, width, height, percent, false)
}

func scene5(context *cairo.Context, width, height, percent float64) {
	obj = wire.GridBox(400, 400, 400, 10, 10, 10, false)
	scene(context, width, height, percent, false)
}

func scene6(context *cairo.Context, width, height, percent float64) {
	obj = wire.ShapeFromLure(l3d.NewAizawa(), 10000)
	scene(context, width, height, percent, true)
}

func scene7(context *cairo.Context, width, height, percent float64) {
	obj = skull
	scene(context, width, height, percent, true)
}

func scene(context *cairo.Context, width, height, percent float64, renderPoints bool) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	wire.SetFog(true, 800, 5000)

	s := shape.TranslatedZ(blmath.Lerp(percent, 2000, -3000))

	y := blmath.Lerp(percent, 12200, -1200)
	s.Points.Push(wire.NewPoint(0, y, 500), 500)
	s2 := obj.Clone()
	s2.Rotate(percent*tau, percent*2*tau, 0)
	s2.Translate(0, y, 500)
	s.Stroke(2)
	if renderPoints {
		s2.RenderPoints(2)
	} else {
		s2.Stroke(2)
	}

	if blur {
		context.GaussianBlur(20)
		s.Stroke(1)
		if renderPoints {
			s2.RenderPoints(1)
		} else {
			s2.Stroke(1)
		}
	}
}
