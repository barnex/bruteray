package shape

import (
	"github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
)

func ExampleSphere() {
	e := br.NewEnv()
	e.Add(NSphere(1, mat.ShadeShape(br.RED)))
	raster.Example(e)
	//Output:
	//![fig](/doc/shape_ExampleSphere.jpg)
}
