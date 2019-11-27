package api

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/barnex/bruteray/tracer"
)

func Animate(numFrame int, f func(frame int) Spec) {
	flag.Parse()
	for i := 0; i < numFrame; i++ {
		fmt.Println("frame", i+1)
		renderFrame(f(i), i)
	}
	fmt.Println("DONE")
}

func renderFrame(spec Spec, i int) {
	spec.InitDefaults()

	aa := (spec.NumPass > 1)
	s := tracer.NewSampler(spec.ImageFunc(), spec.Width, spec.Height, aa)

	for i := 0; i < spec.NumPass; i++ {
		s.Sample(1) // TODO: Sample(N) is broken for high N
	}

	//pp := Postprocess.ApplyTo(s.StoredImage(), imagef.PixelSize(s.Bounds().Dx(), s.Bounds().Dy()))
	//printTime("postprocess")
	//		check(save(pp, ""))
	dir := noExt(*flagO) + "-frames"
	if err := os.Mkdir(dir, 0777); err != nil {
		fmt.Println(err)
	}
	fname := path.Join(dir, fmt.Sprintf("%05d.jpg", i))
	check(savef(s.Image(), path.Join(dir, "last.jpg"), ""))
	check(savef(s.Image(), fname, ""))
}
