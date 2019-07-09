package api

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/barnex/bruteray/v2/builder"
	"github.com/barnex/bruteray/v2/color"
	"github.com/barnex/bruteray/v2/geom"
	imagef "github.com/barnex/bruteray/v2/image"
	"github.com/barnex/bruteray/v2/light"
	"github.com/barnex/bruteray/v2/material"
	"github.com/barnex/bruteray/v2/post"
	"github.com/barnex/bruteray/v2/sampler"

	_ "net/http/pprof"
)

var (
	flagH     = flag.Int("h", 0, "override image height")
	flagN     = flag.Int("n", 0, "override number of passes")
	flagR     = flag.Int("r", 0, "override recursion depth")
	flagW     = flag.Int("w", 0, "override image width")
	flagQ     = flag.Int("q", 0, "override jpeg quality")
	flagO     = flag.String("o", "out.jpeg", "output file")
	flagScale = flag.Int("s", 1, "scale size")
	flagPProf = flag.String("pprof", "", "Remote pprof port.")
)

//
var (
	Width       = 1920 / 4
	Height      = 1080 / 4
	Recursion   = 1
	NumPass     = 1
	JPEGQuality = 95
)

var (
	scene       builder.SceneBuilder
	Camera      = &scene.Camera
	AntiAlias   bool
	Postprocess post.Params
)

func init() {
	log.SetFlags(0)
	reset()
}

func reset() {
	scene = builder.SceneBuilder{}
	Camera = &scene.Camera
	Camera.Frame = geom.XYZ
	AntiAlias = true
	Postprocess = post.Params{}
}

type (
	Color   = color.Color
	Vec     = geom.Vec
	Builder = builder.Builder
)

func Add(b ...Builder) { scene.Add(b...) }

var (
	RectangleLight = light.NewRectangleLight
	PointLight     = light.NewPointLight

	Rectangle  = builder.NewRectangle
	Parametric = builder.NewParametric
	Sphere     = builder.NewSphere
	Tree       = builder.NewTree
	PlyFile    = builder.PlyFile
	Sheet      = builder.NewSheet
	BumpMapped = builder.BumpMapped

	Blend          = material.Blend
	Flat           = material.Flat
	Mate           = material.Mate
	Normal         = material.Normal()
	Normal2        = material.Normal2()
	Grid           = material.Grid()
	ReflectFresnel = material.ReflectFresnel
	Reflective     = material.Reflective
	Shiny          = material.Shiny

	Pitch       = geom.Pitch
	Roll        = geom.Roll
	Yaw         = geom.Yaw
	Scale       = geom.Scale
	Translate   = geom.Translate
	TranslateTo = geom.TranslateTo

	White = color.White
	Black = color.Black
	Red   = color.Red
	Green = color.Green
	Blue  = color.Blue
	Gray  = color.Gray
	EV    = color.EV

	Ex = geom.Ex
	Ey = geom.Ey
	Ez = geom.Ez
)

const (
	Pi  = math.Pi
	Deg = geom.Deg
	X   = geom.X
	Y   = geom.Y
	Z   = geom.Z
)

func Render() {
	initFlags()

	if *flagPProf != "" {
		log.Println("pprof", *flagPProf)
		go func() {
			log.Fatal(http.ListenAndServe(*flagPProf, nil))
		}()
	}
	printTime("user scene definition")

	built := scene.Build()
	printTime("build scene graph")

	print("rendering:", *flagO, Width, "x", Height, ",", NumPass, "passes, ", Recursion, "recursion depth...")

	s := sampler.NewStratifier(built.ImageFunc(Recursion), Width, Height, AntiAlias)

	passBeforeSave := 0
	for i := 0; i < NumPass; i++ {

		passBeforeSave++
		s.Sample(passBeforeSave)
		printTime("render")
		print(s.Stats())

		pp := Postprocess.ApplyTo(s.StoredImage(), imagef.PixelSize(s.Bounds().Dx(), s.Bounds().Dy()))
		printTime("postprocess")

		check(save(pp, ""))
		printTime("encode")

		check(save(s.SamplingImage(), "-sampling"))
		printTime("sampling image")
	}

	print("DONE\n")
}

func save(img image.Image, suffix string) error {
	fname := *flagO
	ext := path.Ext(fname)

	var b bytes.Buffer
	var err error
	switch ext {
	case ".png":
		err = png.Encode(&b, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(&b, img, &jpeg.Options{Quality: JPEGQuality})
	default:
		err = fmt.Errorf("save %q: unknown image format extension", fname)
	}
	if err != nil {
		return err
	}

	fname = fname[:len(fname)-len(ext)] + suffix + ext
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(b.Bytes())
	return err
}

func initFlags() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if flag.NArg() != 0 {
		fatal("No command-line arguments allowed, have:", flag.Args())
	}
	if *flagW != 0 {
		Width = *flagW * *flagScale
	}
	if *flagH != 0 {
		Height = *flagH * *flagScale
	}
	if *flagR != 0 {
		Recursion = *flagR
	}
	if *flagN != 0 {
		NumPass = *flagN
	}
}

var start = time.Now()

func printTime(msg string) {
	print(msg, ":", time.Since(start).Round(time.Millisecond))
	start = time.Now()
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
