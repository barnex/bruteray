package shape

import (
	"github.com/barnex/bruteray/br"
	"github.com/barnex/bruteray/mat"
	"github.com/barnex/bruteray/raster"
)

func ExampleSphere() {
	e := br.NewEnv()
	e.Add(NSphere(1, mat.ShadeShape(br.RED)))
	raster.Standard(e)
	//Output:
	//![fig](shots/062.jpg)
}
