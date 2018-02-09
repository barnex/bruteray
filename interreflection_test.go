package bruteray_test

import (
	"fmt"
	"testing"

	. "github.com/barnex/bruteray"
	"github.com/barnex/bruteray/sample"
)

// Test convergence of diffuse interreflection:
//
// 1) We are inside a highly reflective white box containing a point light.
// If the pre-factor for interreflection is slightly too high,
// deeper recursion will diverge to an infinitely bright image instead of converge.
//
// 2) We are inside a 100% reflective white box containing a point light.
// This is not a physical situation and the intensity should diverge to infinity.
// If the pre-factor for interreflection is slightly too low, divergence will not happen.
func TestDiffuse1(t *testing.T) {
	for _, refl := range []float64{0.8, 1} {
		refl := refl
		for _, r := range []int{1, 16, 128} {
			e := whitebox(refl)
			e.Recursion = r
			t.Run(fmt.Sprintf("refl=%v,rec=%v", refl, e.Recursion), func(t *testing.T) {
				t.Parallel()
				img := sample.MakeImage(testW/4, testH/4)
				nPass := 2
				sample.MultiPass(e, img, nPass)
				name := fmt.Sprintf("diffuse1-refl%v-rec%v", refl, e.Recursion)
				CompareImg(t, e, img, 0, name, 10)
			})
		}
	}
}

func whitebox(refl float64) *Env {
	e := NewEnv()
	white := Diffuse(WHITE.Mul(refl))
	e.Add(
		Sheet(Ey, -1, white),
		Sheet(Ey, 1, white),
		Sheet(Ex, -1, white),
		Sheet(Ex, 1, white),
		Sheet(Ez, -1, white),
		Sheet(Ez, 1, white),
	)
	e.AddLight(PointLight(Vec{}, WHITE.Mul(EV(-3)).Mul(4*Pi)))
	e.Camera = Camera(0.75).Transl(0, 0, -1)
	e.Camera.AA = true
	return e
}
