package api

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"time"

	"github.com/barnex/bruteray/sampler"
)

func Render(spec Spec) {
	flag.Parse()
	initDefaults(&spec)

	//print("rendering:", *flagO, Width, "x", Height, ",", NumPass, "passes, ", Recursion, "recursion depth...")
	s := sampler.NewAdaptive(spec.imageFunc(), spec.Width, spec.Height, true)

	passBeforeSave := 0
	i := 0
	for i < spec.NumPass {
		passBeforeSave++
		s.Sample(passBeforeSave)
		i += passBeforeSave
		printTime("render")
		print(s.Stats())

		//pp := Postprocess.ApplyTo(s.StoredImage(), imagef.PixelSize(s.Bounds().Dx(), s.Bounds().Dy()))
		//printTime("postprocess")
		//		check(save(pp, ""))

		check(save(s.StoredImage(), ""))
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

var start = time.Now()

func printTime(msg string) {
	print(msg, ":", time.Since(start).Round(time.Millisecond))
	start = time.Now()
}
