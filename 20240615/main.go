// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/noise"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240615"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 420)
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
	sphere01 *wire.Shape
	torus01  *wire.Shape
	cyl01    *wire.Shape
	box01    *wire.Shape
	cone01   *wire.Shape
)

func init() {
	sphere01 = wire.Sphere(100, 24, 48, true, true)
	torus01 = wire.Torus(100, 50, tau, 32, 20, true, true)
	cyl01 = wire.Cylinder(200, 50, 20, 20, true, true)
	box01 = wire.Box(75, 200, 50)
	box01.Subdivide(5)
	cone01 = wire.Cone(300, 75, 0, 20, 20, true, true)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	wire.InitWorld(context, 200, 200, 800)
	wire.SetFog(true, 800, 1200)
	wire.SetWaterLevel(true, 50, 100)
	wire.SetClipping(50, 2000)
	random.Seed(0)

	s := sphere01.Clone()
	s.Rotate(percent*tau, percent*2*tau, 0)
	s.Translate(-200, blmath.LoopSin(percent*4, 120, 80), blmath.Lerp(percent, 800, -800))

	t := torus01.Clone()
	t.Rotate(percent*tau, percent*2*tau, 0)
	perc := math.Mod(percent+0.2, 1)
	t.Translate(250, blmath.LoopSin((percent+random.Float())*4, 120, 80), blmath.Lerp(perc, 800, -800))

	c := cyl01.Clone()
	c.Rotate(percent*2*tau, -percent*tau, 0)
	perc = math.Mod(percent+0.4, 1)
	c.Translate(50, blmath.LoopSin((percent+random.Float())*4, 120, 80), blmath.Lerp(perc, 800, -800))

	b := box01.Clone()
	b.Rotate(percent*tau, -percent*2*tau, 0)
	perc = math.Mod(percent+0.6, 1)
	b.Translate(-100, blmath.LoopSin((percent+random.Float())*4, 120, 80), blmath.Lerp(perc, 800, -800))

	cone := cone01.Clone()
	cone.Rotate(-percent*2*tau, percent*tau, 0)
	perc = math.Mod(percent+0.8, 1)
	cone.Translate(200, blmath.LoopSin((percent+random.Float())*4, 120, 80), blmath.Lerp(perc, 800, -800))

	dots := wire.NewShape()
	for range 1000 {
		dots.AddXYZ(
			random.FloatRange(-1000, 1000),
			50,
			random.FloatRange(-800, 0),
		)
	}
	dots.AddShape(dots.TranslatedZ(800))
	dots.TranslateZ(blmath.Lerp(percent, 0, -800))
	scale := 0.007
	for _, p := range dots.Points {
		n := noise.Simplex2(p.X*scale, p.Z*scale)
		p.Y += n * 15
		p.X += n * 5
	}

	w := 1.5
	s.Stroke(w)
	t.Stroke(w)
	c.Stroke(w)
	b.Stroke(w)
	cone.Stroke(w)
	dots.RenderPoints(w / 3)

	if blur {
		w /= 2
		context.GaussianBlur(10)
		s.Stroke(w)
		t.Stroke(w)
		c.Stroke(w)
		b.Stroke(w)
		cone.Stroke(w)
		dots.RenderPoints(w / 3)
	}
}
