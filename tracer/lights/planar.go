package lights

import (
	"github.com/barnex/bruteray/geom"
	"github.com/barnex/bruteray/tracer/materials"
	"github.com/barnex/bruteray/tracer/objects"
	"github.com/barnex/bruteray/tracer/sequence"
	. "github.com/barnex/bruteray/tracer/types"
)

func RectangleLight(brightness Color, w, h float64, center Vec) Light {
	return planarLight(brightness, w, h, center, 1, nil, transfRect)
}

func DiskLight(brightness Color, diam float64, center Vec) Light {
	relSurf := Pi / 4 // Surface area relative to enclosing rectangle.
	return planarLight(brightness, diam, diam, center, relSurf, objects.Cylinder(nil, diam, Inf, center), sequence.UniformDisk)
}

// SunLight constructs a far-away circlular light source, like the sun.
// angularDiam is the angular diameter (in radians).
//
// The angular diameter of the sun seem from the earth is about 0.53 degrees
// (https://en.wikipedia.org/wiki/Angular_diameter#Use_in_astronomy)
//
// brightnessAtEarth is the brightness that a 100% reflective matte surface
// receives when perpendicular to the sunlight. A brightness of (1,1,1) is thus
// easy to understand: this is the highest brightness that is guaranteed
// not to overexpose any scene.
//
// yaw is the horizontal angle with respect to the -z axis.
// Thus 0 degree yaw means the sun is in the north.
//
// pitch is the vertical angle above the horizon.
// 90 degree pitch means the sun is in the zenith.
func SunLight(brightnessAtEarth Color, angularDiam float64, yaw, pitch float64) Light {
	const r = 1e6 // distance at which we place the sun. Nearly infinite.

	// brightness of sun surface to get desired brighness on matte surface
	// with sun in zenith. This is an approximation for small angular size:
	// the solid angle of the sun, dividied by the surface of the hemisphere (2 pi)
	// times 1/(2pi) BDRF normalization.
	bright2 := brightnessAtEarth.Mul(4 / (angularDiam * angularDiam))

	disk := DiskLight(bright2, angularDiam*r, Vec{})
	transf := geom.ComposeLR(
		geom.Rotate(O, Ex, -90*Deg),
		geom.Translate(Vec{0, 0, -r}),
		geom.Rotate(O, Ex, pitch),
		geom.Rotate(O, Ey, yaw),
	)
	return Transformed(disk, transf)
}

// planar is a rectangle or disk light.
type planar struct {
	w, h         float64
	center       Vec
	totalPower   Color // W/m2 (pi?)
	object       objects.Interface
	restrict     objects.Interface
	transfSample func(u, v float64) (x, y float64)
}

func planarLight(brightness Color, w, h float64, center Vec, relSurf float64, restrict objects.Interface, transfSample func(u, v float64) (x, y float64)) Light {
	mat := materials.TwoSided(
		materials.Flat(brightness),
		materials.Flat(Color{0, 0, 0}),
	)
	if restrict == nil { // TODO: remove this, since it can only be a disk or rectangle.
		restrict = allSpace{}
	}
	return &planar{
		w:            w,
		h:            h,
		center:       center,
		totalPower:   brightness.Mul(w * h * relSurf),
		restrict:     restrict,
		object:       objects.Restrict(objects.Rectangle(mat, w, h, center), restrict),
		transfSample: transfSample,
	}
}

func (l *planar) Sample(ctx *Ctx, target Vec) (Vec, Color) {
	p := l.samplePos(ctx)
	n := Vec{0, -1, 0}

	delta := target.Sub(p)
	cosTheta := n.Dot(delta.Normalized())
	I := (1 / Pi) / delta.Len2() * cosTheta
	if I < 0 {
		I = 0
	}
	return p, l.totalPower.Mul(I)
}

func (l *planar) samplePos(ctx *Ctx) Vec {
	x, y := l.transfSample(ctx.Generate2())
	return Vec{(0.5 * x) * l.w, Tiny, (0.5 * y) * l.h}.Add(l.center)

	//u, v := ctx.Generate2()
	//p := Vec{(u - 0.5) * l.w, Tiny, (v - 0.5) * l.h}.Add(l.center)

	//n := 0
	//// TODO: use disk mapping for disk light
	//for !l.restrict.Inside(p) {
	//	u, v = ctx.Generate2()
	//	p = Vec{(u - 0.5) * l.w, Tiny, (v - 0.5) * l.h}.Add(l.center)

	//	// if we made a mistake so there is zero overlap,
	//	// panic rather than hang.
	//	n++
	//	if n == 1024*1024 {
	//		panic("light: sampleUV: no point found inside after 1M samples")
	//	}
	//}
	//return p
}

func (l *planar) Object() Object {
	return l.object
}

type allSpace struct{}

func (allSpace) Inside(Vec) bool {
	return true
}

func (allSpace) Bounds() objects.BoundingBox {
	return objects.BoundingBox{Min: Vec{-Inf, -Inf, -Inf}, Max: Vec{Inf, Inf, Inf}}
}

func (allSpace) Intersect(*Ray) HitRecord { panic("unused") }

func transfRect(u, v float64) (x, y float64) {
	return 2*u - 1, 2*v - 1
}
