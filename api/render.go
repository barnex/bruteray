package api

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"time"

	imagef "github.com/barnex/bruteray/image"
	"github.com/barnex/bruteray/image/ppm"
	"github.com/barnex/bruteray/sampler"
)

var (
	flagHTTP = flag.String("http", "", "HTTP port")
)

func Render(spec Spec) {
	flag.Parse()
	spec.InitDefaults()

	switch {
	default:
		renderLocal(spec)
	case *flagHTTP != "":
		fmt.Println("Serving", *flagHTTP)
		check(Serve(*flagHTTP, spec))
	}
}

func renderLocal(spec Spec) {
	//print("rendering:", *flagO, Width, "x", Height, ",", NumPass, "passes, ", Recursion, "recursion depth...")
	s := sampler.NewAdaptive(spec.ImageFunc(), spec.Width, spec.Height, true)

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

	check(savePPM(noExt(*flagO)+".ppm", s.StoredImage()))

	print("DONE\n")
}

func save(img image.Image, suffix string) error {
	return savef(img, *flagO, suffix)
}

func savef(img image.Image, fname, suffix string) error {
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

	fname = noExt(fname) + suffix + ext
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

func noExt(fname string) string {
	ext := path.Ext(fname)
	return fname[:len(fname)-len(ext)]
}

func savePPM(fname string, img imagef.Image) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	return ppm.EncodeAscii16(w, img)
}
