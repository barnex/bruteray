// Package api provides a high-level API for defining BruteRay scenes.
// This is the only package users usually need to interact with.
package api

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/imagef/colorf"
	"github.com/barnex/bruteray/tracer"
	"github.com/barnex/bruteray/tracer/lights"
	"github.com/barnex/bruteray/tracer/materials"
	"github.com/barnex/bruteray/tracer/media"
	"github.com/barnex/bruteray/tracer/objects"
)

var (
	flagH       = flag.Int("h", 0, "override image height")
	flagN       = flag.Int("n", 0, "override number of passes")
	flagR       = flag.Int("r", 0, "override recursion depth")
	flagW       = flag.Int("w", 0, "override image width")
	flagQ       = flag.Int("q", 0, "override jpeg quality")
	flagO       = flag.String("o", "out.jpeg", "output file")
	flagScale   = flag.Int("s", 1, "scale size")
	flagPProf   = flag.String("pprof", "", "Remote pprof port.")
	JPEGQuality = 90
)

type (
	Color    = colorf.Color
	Light    = tracer.Light
	Material = tracer.Material
	Medium   = tracer.Medium
	Vec      = geom.Vec
)

var (
	And = objects.And
	Not = objects.Not

	Black   = colorf.Black
	Blue    = colorf.Blue
	Cyan    = colorf.Cyan
	Gray    = colorf.Gray
	Green   = colorf.Green
	Magenta = colorf.Magenta
	Red     = colorf.Red
	White   = colorf.White
	Yellow  = colorf.Yellow

	RectangleLight = lights.RectangleLight
	PointLight     = lights.PointLight
	DiskLight      = lights.DiskLight
	SunLight       = lights.SunLight

	Matte          = materials.Matte
	Reflective     = materials.Reflective
	Refractive     = materials.Refractive
	Transparent    = materials.Transparent
	Flat           = materials.Flat
	Shiny          = materials.Shiny
	Blend          = materials.Blend
	BlendMap       = materials.BlendMap
	ReflectFresnel = materials.ReflectFresnel

	ExpFog = media.ExpFog
	Fog    = media.Fog

	Ex = geom.Ex
	Ey = geom.Ey
	Ez = geom.Ez
	O  = geom.O
)

func V(x, y, z float64) Vec {
	return Vec{x, y, z}
}

func C(r, g, b float64) Color {
	return Color{R: r, G: g, B: b}
}

const (
	Pi  = math.Pi
	Deg = geom.Deg
	X   = geom.X
	Y   = geom.Y
	Z   = geom.Z
)

func initFlags() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if flag.NArg() != 0 {
		fatal("No command-line arguments allowed, have:", flag.Args())
	}
	//	if *flagR != 0 {
	//		Recursion = *flagR
	//	}
	//	if *flagN != 0 {
	//		NumPass = *flagN
	//	}
}

func print(x ...interface{}) {
	fmt.Println(x...)
}

func check(e error) {
	if e != nil {
		fatal(e)
	}
}

func fatal(e ...interface{}) {
	fmt.Fprintln(os.Stderr, e...)
	os.Exit(1)
}
