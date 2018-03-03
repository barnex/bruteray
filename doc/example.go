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

func Example(e *Env) {
	pc, _, _, _ := runtime.Caller(1)
	base := path.Ext(path.Base(runtime.FuncForPC(pc).Name()))[1:]
	file := "../doc/" + base + ".jpg"
	defer fmt.Printf(base)

	if _, err := os.Stat(file); err == nil {
		return // already there
	}

	m := mat.Checkboard(1, mat.Diffuse(WHITE.EV(-.3)), mat.Diffuse(WHITE.EV(0)))
	e.Add(shape.Sheet(Ey, 0, m))
	e.AddLight(light.Sphere(Vec{2, 2, -2}, 0.5, WHITE.EV(6)))
	e.SetAmbient(WHITE.EV(-5))

	img := raster.MakeImage(1920/3, 1080/3)
	cam := raster.Camera(1).Transl(0, 0.8, -2).Transf(RotX4(10 * Deg))
	cam.AA = true
	raster.MultiPass(cam, e, img, 300)
	e.Recursion = 3

	raster.MustEncode(img, file)
}
