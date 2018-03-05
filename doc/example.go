package doc

import (
	"fmt"
	"os"
	"path"
	"runtime"

	. "github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/light"
	"github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
	"github.com/barnex/bruteray/shape"
)

func Show(obj ...Obj) {
	e := NewEnv()
	e.Add(obj...)
	example(e)
}

func Example(e *Env) {
	example(e)
}

func example(e *Env) {
	pc, _, _, _ := runtime.Caller(2)
	base := path.Ext(path.Base(runtime.FuncForPC(pc).Name()))[1:]
	file := "../doc/" + base + ".jpg"
	defer fmt.Printf(base)

	if _, err := os.Stat(file); err == nil {
		return // already there
	}

	m := mat.Checkboard(1, mat.Diffuse(WHITE.EV(-.3)), mat.Diffuse(WHITE.EV(0)))
	e.Add(shape.NewSheet(Ey, 0, m))
	e.AddLight(light.Sphere(Vec{2, 2, -2}, 0.5, WHITE.EV(6)))
	e.SetAmbient(WHITE.EV(-5))
	e.Cutoff = EV(8)

	img := raster.MakeImage(1920/3, 1080/3)
	cam := raster.Camera(1).Transl(0, 1.2, -2.2).Transf(RotX4(18 * Deg))
	cam.AA = true
	raster.MultiPass(cam, e, img, 3)
	e.Recursion = 3

	raster.MustEncode(img, file)
}
