// Package main renders an image, gif or video
package main

import (
	"math"

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
	fileName := "20240606"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 180)
		program.AddSceneWithFrames(scene2, 180)
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

	shape := wire.Cylinder(400, 100, 20, 20, true, true)
	shape.TranslateY(blmath.LoopSin(percent+0.25, 100, 300))
	shape.RotateX(blmath.LoopSin(percent, -0.5, 0.5))
	shape.RotateY(blmath.LoopSin(percent+0.2, -0.3, 0.3))
	shape.Translate(400, 0, 0)
	shape.Subdivide(4)
	shape.Cull(func(p *wire.Point) bool {
		return p.Y < 200
	})

	ball := wire.Sphere(200, 12, 24, true, true)
	ball.RotateZ(0.5)
	ball.Translate(
		-400,
		easing.ElasticEaseOut(percent, -800, 180),
		600,
	)
	ball.Subdivide(4)
	ball.Cull(func(p *wire.Point) bool {
		return p.Y < 200
	})

	box := wire.RandomSurfaceBox(500, 150, 150, 2000)
	box.TranslateY(blmath.LoopSin(percent, 150, 250))
	box.Rotate(
		blmath.LoopSin(percent-0.2, -0.3, 0.3),
		0.5,
		blmath.LoopSin(percent+0.1, -0.2, 0.2),
	)
	box.Translate(-200, 0, -350)
	box.Cull(func(p *wire.Point) bool {
		return p.Y < 200
	})

	t := wire.Torus(400, 150, tau, 24, 24, true, true)
	t.TranslateY(blmath.LoopSin(percent+0.5, 100, 300))
	t.Rotate(
		blmath.LoopSin(percent, -0.5, -0.25),
		0,
		blmath.LoopSin(percent+0.25, -0.3, 0.3),
	)
	t.Translate(150, 0, 1400)
	t.Subdivide(4)
	t.Cull(func(p *wire.Point) bool {
		return p.Y < 200
	})

	shape.Stroke(2)
	ball.Stroke(2)
	t.Stroke(2)
	box.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		shape.Stroke(1)
		ball.Stroke(1)
		t.Stroke(1)
		box.RenderPoints(1)
	}
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	random.Seed(0)

	shape := wire.Cylinder(400, 100, 20, 20, true, true)
	shape.TranslateY(blmath.LoopSin(percent+0.25, 100, 300))
	shape.RotateX(blmath.LoopSin(percent, -0.5, 0.5))
	shape.RotateY(blmath.LoopSin(percent+0.2, -0.3, 0.3))
	shape.Translate(400, 0, 0)
	shape.Subdivide(4)
	shape.Cull(func(p *wire.Point) bool {
		return p.Y < 200
	})

	ball := wire.Sphere(200, 12, 24, true, true)
	ball.Rotate(percent*0.5, percent, 0.5)
	ball.Translate(
		-400,
		180+percent*230,
		600,
	)
	ball.Subdivide(4)
	ball.Cull(func(p *wire.Point) bool {
		return p.Y < 200
	})

	box := wire.RandomSurfaceBox(500, 150, 150, 2000)
	box.TranslateY(blmath.LoopSin(percent, 150, 250))
	box.Rotate(
		blmath.LoopSin(percent-0.2, -0.3, 0.3),
		0.5,
		blmath.LoopSin(percent+0.1, -0.2, 0.2),
	)
	box.Translate(-200, 0, -350)
	box.Cull(func(p *wire.Point) bool {
		return p.Y < 200
	})

	t := wire.Torus(400, 150, tau, 24, 24, true, true)
	t.TranslateY(blmath.LoopSin(percent+0.5, 100, 300))
	t.Rotate(
		blmath.LoopSin(percent, -0.5, -0.25),
		0,
		blmath.LoopSin(percent+0.25, -0.3, 0.3),
	)
	t.Translate(150, 0, 1400)
	t.Subdivide(4)
	t.Cull(func(p *wire.Point) bool {
		return p.Y < 200
	})

	shape.Stroke(2)
	ball.Stroke(2)
	t.Stroke(2)
	box.RenderPoints(2)

	if blur {
		context.GaussianBlur(20)
		shape.Stroke(1)
		ball.Stroke(1)
		t.Stroke(1)
		box.RenderPoints(1)
	}
}
