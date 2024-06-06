// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/lures/l3d"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	fileName := "20240607"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 180)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

//revive:disable:unused-parameter
const (
	tau   = blmath.Tau
	pi    = math.Pi
	count = 5000
)

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 0, 800)

	//////////////////////////////
	// dequan
	//////////////////////////////
	s0 := wire.NewShape()
	lure0 := l3d.NewDequanLi()
	x, y, z := lure0.InitVals3d()
	s0.AddXYZ(x, y, z)

	for range count {
		x, y, z = lure0.Iterate(x, y, z)
		s0.AddXYZ(x, y, z)
	}
	s0.Translate(lure0.Center3d())
	s0.UniScale(lure0.Scale / 2)

	//////////////////////////////
	// chen
	//////////////////////////////
	s1 := wire.NewShape()
	lure1 := l3d.NewChenLee()
	x, y, z = lure1.InitVals3d()
	s1.AddXYZ(x, y, z)

	for range count {
		x, y, z = lure1.Iterate(x, y, z)
		s1.AddXYZ(x, y, z)
	}
	s1.Translate(lure1.Center3d())
	s1.UniScale(lure1.Scale / 2)

	//////////////////////////////
	// aizawa
	//////////////////////////////
	s2 := wire.NewShape()
	lure2 := l3d.NewAizawa()
	x, y, z = lure1.InitVals3d()
	s2.AddXYZ(x, y, z)

	for range count {
		x, y, z = lure2.Iterate(x, y, z)
		s2.AddXYZ(x, y, z)
	}
	s2.Translate(lure2.Center3d())
	s2.UniScale(lure2.Scale / 2)

	//////////////////////////////
	// 4wing
	//////////////////////////////
	s3 := wire.NewShape()
	lure3 := l3d.NewFourWings()
	x, y, z = lure2.InitVals3d()
	s3.AddXYZ(x, y, z)

	for range count {
		x, y, z = lure3.Iterate(x, y, z)
		s3.AddXYZ(x, y, z)
	}
	s3.Translate(lure3.Center3d())
	s3.UniScale(lure3.Scale / 3)

	//////////////////////////////
	// merge
	//////////////////////////////

	sa := wire.NewShape()
	sa.Points = s0.Points.Clone()
	sa.Points.Lerp(percent, s1.Points)
	sa.RotateY(percent * tau * 3 / 4)
	sa.Translate(-250, 300, 250)
	sa.RotateY(percent * pi / 2)
	sa.RenderPoints(2)

	sb := wire.NewShape()
	sb.Points = s1.Points.Clone()
	sb.Points.Lerp(percent, s2.Points)
	sb.RotateY(percent * tau * 3 / 4)
	sb.Translate(250, 300, 250)
	sb.RotateY(percent * pi / 2)
	sb.RenderPoints(2)

	sc := wire.NewShape()
	sc.Points = s2.Points.Clone()
	sc.Points.Lerp(percent, s3.Points)
	sc.RotateY(percent * tau * 3 / 4)
	sc.Translate(250, 300, -250)
	sc.RotateY(percent * pi / 2)
	sc.RenderPoints(2)

	sd := wire.NewShape()
	sd.Points = s3.Points.Clone()
	sd.Points.Lerp(percent, s0.Points)
	sd.RotateY(percent * tau * 3 / 4)
	sd.Translate(-250, 300, -250)
	sd.RotateY(percent * pi / 2)
	sd.RenderPoints(2)

	sa.TranslatedY(-300).ScaledY(-1).TranslatedY(500).RenderPoints(0.5)
	sb.TranslatedY(-300).ScaledY(-1).TranslatedY(500).RenderPoints(0.5)
	sc.TranslatedY(-300).ScaledY(-1).TranslatedY(500).RenderPoints(0.5)
	sd.TranslatedY(-300).ScaledY(-1).TranslatedY(500).RenderPoints(0.5)

	g := wire.GridPlane(4000, 4000, 20, 20)
	g.Translate(0, 400, 1200)
	g.Subdivide(10)
	g.Stroke(0.5)

	context.GaussianBlur(20)
	sa.RenderPoints(1)
	sb.RenderPoints(1)
	sc.RenderPoints(1)
	sd.RenderPoints(1)
	g.Stroke(0.25)
}
