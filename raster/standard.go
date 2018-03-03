package raster

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/barnex/bruteray/br"
)

func Example(e *br.Env) {
	pc, _, _, _ := runtime.Caller(1)
	base := strings.Replace(path.Base(runtime.FuncForPC(pc).Name()), ".", "_", 1)
	name := "doc/" + base + ".jpg"
	img := MakeImage(960, 540)
	cam := Camera(1).Transl(0, 0.5, -2).Transf(br.RotX4(10 * br.Deg))
	MultiPass(cam, e, img, 5)
	MustEncode(img, "../"+name)
	fmt.Printf("![fig](%v)\n", name)
}
