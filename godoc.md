# Bruteray

Bi-directional path tracing, a physically accurate method that produces realistic images.

Features:

  * Indirect lighting (global illuminiation)
  * Volumetric lighting
  * Refraction
  * Depth of field

Sub-packages:

    br        core raytracing logic and types
    mat       materials and textures
    light     various types of light sources
    shape     shapes and objects
    csg       constructive solid geometry: combine shapes
    transf    affine transformations on shapes
    raster    turns a scene into a pixel image

Additional material:

	cmd/raywatch    web interface for developing scenes
	serve           web server used by raywatch
	scenes          source files of some scenes
	tutorial        explains some ray-tracing basics


# shape
`import "github.com/barnex/bruteray/shape"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
Package shape implements various shapes and objects.

## <a name="pkg-imports">Imported Packages</a>

- [github.com/barnex/bruteray/br](./../br)

## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func And(a, b CSGObj) CSGObj](#And)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cutout(a CSGObj, b Insider) CSGObj](#Cutout)
* [func Hollow(o CSGObj) CSGObj](#Hollow)
* [func Inverse(o CSGObj) CSGObj](#Inverse)
* [func Minus(a, b CSGObj) CSGObj](#Minus)
* [func MultiOr(o ...CSGObj) CSGObj](#MultiOr)
* [func NewCylinder(dir int, center Vec, diam, h float64, m Material) CSGObj](#NewCylinder)
* [func NewInfCylinder(dir int, diam float64, m Material) \*quad](#NewInfCylinder)
* [func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj](#OldBox)
* [func Or(a, b CSGObj) CSGObj](#Or)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [func SurfaceAnd(a Obj, b CSGObj) Obj](#SurfaceAnd)
* [type Box](#Box)
  * [func NewBox(w, h, d float64, m Material) \*Box](#NewBox)
  * [func (s \*Box) Center() Vec](#Box.Center)
  * [func (s \*Box) Corner(x, y, z int) Vec](#Box.Corner)
  * [func (s \*Box) Hit1(r \*Ray, f \*[]Fragment)](#Box.Hit1)
  * [func (s \*Box) HitAll(r \*Ray, f \*[]Fragment)](#Box.HitAll)
  * [func (s \*Box) Inside(v Vec) bool](#Box.Inside)
  * [func (s \*Box) Normal(p Vec) Vec](#Box.Normal)
  * [func (s \*Box) Transl(d Vec) \*Box](#Box.Transl)
* [type Sheet](#Sheet)
  * [func NewSheet(dir Vec, off float64, m Material) \*Sheet](#NewSheet)
  * [func (s \*Sheet) Hit1(r \*Ray, f \*[]Fragment)](#Sheet.Hit1)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewBox](#example_NewBox)
* [NewSheet](#example_NewSheet)
* [NewSphere](#example_NewSphere)

#### <a name="pkg-files">Package files</a>
[box.go](./box.go) [csg.go](./csg.go) [cylinder.go](./cylinder.go) [doc.go](./doc.go) [quad.go](./quad.go) [rect.go](./rect.go) [sheet.go](./sheet.go) [slab.go](./slab.go) [sphere.go](./sphere.go) [util.go](./util.go) 

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="And">func</a> [And](./csg.go#L30)
``` go
func And(a, b CSGObj) CSGObj
```
Intersection (boolean AND) of two objects.

## <a name="Cube">func</a> [Cube](./box.go#L53)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cutout">func</a> [Cutout](./csg.go#L201)
``` go
func Cutout(a CSGObj, b Insider) CSGObj
```

## <a name="Hollow">func</a> [Hollow](./csg.go#L255)
``` go
func Hollow(o CSGObj) CSGObj
```
Hollow turns a into a hollow surface.
E.g.: a filled cylinder into a hollow tube.

## <a name="Inverse">func</a> [Inverse](./csg.go#L267)
``` go
func Inverse(o CSGObj) CSGObj
```

## <a name="Minus">func</a> [Minus](./csg.go#L161)
``` go
func Minus(a, b CSGObj) CSGObj
```
Subtraction (logical AND NOT) of two objects

## <a name="MultiOr">func</a> [MultiOr](./csg.go#L111)
``` go
func MultiOr(o ...CSGObj) CSGObj
```

## <a name="NewCylinder">func</a> [NewCylinder](./cylinder.go#L10)
``` go
func NewCylinder(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).
TODO: Transl

## <a name="NewInfCylinder">func</a> [NewInfCylinder](./quad.go#L18)
``` go
func NewInfCylinder(dir int, diam float64, m Material) *quad
```

## <a name="OldBox">func</a> [OldBox](./box.go#L45)
``` go
func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj
```
TODO rm

## <a name="Or">func</a> [Or](./csg.go#L70)
``` go
func Or(a, b CSGObj) CSGObj
```
Union (logical OR) of two objects.
TODO: remove in favor of MultiOr

## <a name="Quad">func</a> [Quad](./quad.go#L6)
``` go
func Quad(center Vec, a Vec, b float64, m Material) CSGObj
```

## <a name="Rect">func</a> [Rect](./rect.go#L9)
``` go
func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj
```
A rectangle (i.e. finite sheet) at given position,
with normal vector dir and half-axes rx, ry, rz.

TODO: pass Vec normal, U, V

## <a name="Slab">func</a> [Slab](./slab.go#L5)
``` go
func Slab(dir Vec, off1, off2 float64, m Material) CSGObj
```

## <a name="SurfaceAnd">func</a> [SurfaceAnd](./csg.go#L227)
``` go
func SurfaceAnd(a Obj, b CSGObj) Obj
```
Intersection, treating A as a hollow object.
Equivalent to, but more efficient than And(Hollow(a), b)

## <a name="Box">type</a> [Box](./box.go#L19-L22)
``` go
type Box struct {
    Min, Max Vec
    Mat      Material
}
```

### <a name="NewBox">func</a> [NewBox](./box.go#L10)
``` go
func NewBox(w, h, d float64, m Material) *Box
```
NewBox constructs a box with given width, depth and height.

#### Example:

```go
doc.Show(
NewBox(1, 1, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0}),
)
```

![fig](/doc/ExampleNewBox.jpg)

### <a name="Box.Center">func</a> (\*Box) [Center](./box.go#L24)
``` go
func (s *Box) Center() Vec
```

### <a name="Box.Corner">func</a> (\*Box) [Corner](./box.go#L39)
``` go
func (s *Box) Corner(x, y, z int) Vec
```
Corner returns one of the box's corners:

	Corner( 1, 1, 1) -> right top  back
	Corner(-1,-1,-1) -> left bottom front
	Corner( 1,-1,-1) -> right bottom front
	...

### <a name="Box.Hit1">func</a> (\*Box) [Hit1](./box.go#L57)
``` go
func (s *Box) Hit1(r *Ray, f *[]Fragment)
```

### <a name="Box.HitAll">func</a> (\*Box) [HitAll](./box.go#L59)
``` go
func (s *Box) HitAll(r *Ray, f *[]Fragment)
```

### <a name="Box.Inside">func</a> (\*Box) [Inside](./box.go#L93)
``` go
func (s *Box) Inside(v Vec) bool
```

### <a name="Box.Normal">func</a> (\*Box) [Normal](./box.go#L99)
``` go
func (s *Box) Normal(p Vec) Vec
```

### <a name="Box.Transl">func</a> (\*Box) [Transl](./box.go#L28)
``` go
func (s *Box) Transl(d Vec) *Box
```

## <a name="Sheet">type</a> [Sheet](./sheet.go#L9-L13)
``` go
type Sheet struct {
    // contains filtered or unexported fields
}
```

### <a name="NewSheet">func</a> [NewSheet](./sheet.go#L5)
``` go
func NewSheet(dir Vec, off float64, m Material) *Sheet
```

#### Example:

```go
doc.Show(
NewSheet(Ey, 0.1, mat.Diffuse(RED)),
)
```

![fig](/doc/ExampleNewSheet.jpg)

### <a name="Sheet.Hit1">func</a> (\*Sheet) [Hit1](./sheet.go#L15)
``` go
func (s *Sheet) Hit1(r *Ray, f *[]Fragment)
```

## <a name="Sphere">type</a> [Sphere](./sphere.go#L10-L14)
``` go
type Sphere struct {
    // contains filtered or unexported fields
}
```

### <a name="NewSphere">func</a> [NewSphere](./sphere.go#L6)
``` go
func NewSphere(diam float64, m Material) *Sphere
```

#### Example:

```go
doc.Show(
NewSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0}),
)
```

![fig](/doc/ExampleNewSphere.jpg)

### <a name="Sphere.Hit1">func</a> (\*Sphere) [Hit1](./sphere.go#L26)
``` go
func (s *Sphere) Hit1(r *Ray, f *[]Fragment)
```

### <a name="Sphere.HitAll">func</a> (\*Sphere) [HitAll](./sphere.go#L28)
``` go
func (s *Sphere) HitAll(r *Ray, f *[]Fragment)
```

### <a name="Sphere.Inside">func</a> (\*Sphere) [Inside](./sphere.go#L16)
``` go
func (s *Sphere) Inside(p Vec) bool
```

### <a name="Sphere.Normal">func</a> (\*Sphere) [Normal](./sphere.go#L45)
``` go
func (s *Sphere) Normal(pos Vec) Vec
```

### <a name="Sphere.Transl">func</a> (\*Sphere) [Transl](./sphere.go#L21)
``` go
func (s *Sphere) Transl(d Vec) *Sphere
```

# mat
`import "github.com/barnex/bruteray/mat"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
Package mat implements various types of materials.

## <a name="pkg-imports">Imported Packages</a>

- [github.com/barnex/bruteray/br](./../br)
- [github.com/barnex/bruteray/raster](./../raster)

## <a name="pkg-index">Index</a>
* [func Blend(a float64, matA Material, b float64, matB Material) Material](#Blend)
* [func Bricks(stride, width float64, a, b Material) Material](#Bricks)
* [func Checkboard(stride float64, a, b Material) Material](#Checkboard)
* [func DebugShape(c Color) Material](#DebugShape)
* [func Diffuse(c Texture) Material](#Diffuse)
* [func Diffuse0(c Texture) Material](#Diffuse0)
* [func Diffuse00(c Color) Material](#Diffuse00)
* [func Distort(seed int, n int, K Vec, ampli float64, orig Material) Material](#Distort)
* [func Load(name string) (raster.Image, error)](#Load)
* [func MustLoad(name string) raster.Image](#MustLoad)
* [func Reflective(c Color) Material](#Reflective)
* [func Refractive(n1, n2 float64) Material](#Refractive)
* [func ShadeNormal(dir Vec) Material](#ShadeNormal)
* [func Shiny(c Color, reflectivity float64) Material](#Shiny)
* [func Waves(seed int, n int, K Vec, col func(float64) Material) Material](#Waves)
* [type FlatColor](#FlatColor)
  * [func Flat(c br.Color) \*FlatColor](#Flat)
  * [func (s \*FlatColor) At(\_ br.Vec) br.Color](#FlatColor.At)
  * [func (s \*FlatColor) Shade(\_ \*br.Ctx, \_ \*br.Env, \_ int, \_ \*br.Ray, \_ br.Fragment) br.Color](#FlatColor.Shade)
* [type ImgTex](#ImgTex)
  * [func NewImgTex(img raster.Image, p0, pu, pv Vec) \*ImgTex](#NewImgTex)
  * [func (c \*ImgTex) At(pos Vec) Color](#ImgTex.At)
  * [func (c \*ImgTex) Shade(ctx \*Ctx, e \*Env, N int, r \*Ray, frag Fragment) Color](#ImgTex.Shade)
* [type ShadeDir](#ShadeDir)
  * [func (s ShadeDir) Shade(ctx \*Ctx, e \*Env, N int, r \*Ray, frag Fragment) Color](#ShadeDir.Shade)
* [type Texture](#Texture)

#### <a name="pkg-examples">Examples</a>
* [Blend](#example_Blend)
* [DebugShape](#example_DebugShape)
* [Diffuse](#example_Diffuse)
* [Flat](#example_Flat)
* [Reflective](#example_Reflective)
* [Refractive](#example_Refractive)

#### <a name="pkg-files">Package files</a>
[diffuse.go](./diffuse.go) [diffuse_noshadow.go](./diffuse_noshadow.go) [flat.go](./flat.go) [material.go](./material.go) [procedural.go](./procedural.go) [texture.go](./texture.go) 

## <a name="Blend">func</a> [Blend](./material.go#L109)
``` go
func Blend(a float64, matA Material, b float64, matB Material) Material
```
Blend mixes two materials with certain weights. E.g.:

	Blend(0.9, Diffuse(WHITE), 0.1, Reflective(WHITE))  // 90% mate + 10% reflective, like a shiny billiard ball.

#### Example:

```go
white := Diffuse(WHITE)
refl := Reflective(WHITE)
doc.Show(
shape.NewSphere(1, Blend(0.95, white, 0.05, refl)).Transl(Vec{0, 0.5, 0}),
)
```

![fig](/doc/ExampleBlend.jpg)
## <a name="Bricks">func</a> [Bricks](./procedural.go#L30)
``` go
func Bricks(stride, width float64, a, b Material) Material
```

## <a name="Checkboard">func</a> [Checkboard](./procedural.go#L9)
``` go
func Checkboard(stride float64, a, b Material) Material
```

## <a name="DebugShape">func</a> [DebugShape](./material.go#L154)
``` go
func DebugShape(c Color) Material
```
DebugShape is a debug material that renders the object's shape only,
even if no lighting is present. Useful while defining a scene before
worrying about lighting.

#### Example:

```go
e := NewEnv()
e.Add(shape.NewSheet(Ey, 0, DebugShape(WHITE)))
e.Add(shape.NewSphere(1, DebugShape(WHITE)).Transl(Vec{0, 0.5, 0}))
// Note: no light source added
doc.Example(e)
```

![fig](/doc/ExampleDebugShape.jpg)
## <a name="Diffuse">func</a> [Diffuse](./diffuse.go#L10)
``` go
func Diffuse(c Texture) Material
```
A Diffuse material appears perfectly mate,
like paper or plaster.
See <a href="https://en.wikipedia.org/wiki/Lambertian_reflectance">https://en.wikipedia.org/wiki/Lambertian_reflectance</a>.

#### Example:

```go
doc.Show(
shape.NewSphere(1, Diffuse(WHITE)).Transl(Vec{0, 0.5, 0}),
)
```

![fig](/doc/ExampleDiffuse.jpg)
## <a name="Diffuse0">func</a> [Diffuse0](./diffuse.go#L46)
``` go
func Diffuse0(c Texture) Material
```
Diffuse material with direct illumination only (no interreflection).
Intended for debugging or rapid previews. Diffuse is much more realistic.

## <a name="Diffuse00">func</a> [Diffuse00](./diffuse_noshadow.go#L9)
``` go
func Diffuse00(c Color) Material
```
Diffuse material with direct illumination only and no shadows.
Intended for the tutorial.

## <a name="Distort">func</a> [Distort](./procedural.go#L67)
``` go
func Distort(seed int, n int, K Vec, ampli float64, orig Material) Material
```

## <a name="Load">func</a> [Load](./texture.go#L67)
``` go
func Load(name string) (raster.Image, error)
```

## <a name="MustLoad">func</a> [MustLoad](./texture.go#L59)
``` go
func MustLoad(name string) raster.Image
```

## <a name="Reflective">func</a> [Reflective](./material.go#L26)
``` go
func Reflective(c Color) Material
```
A Reflective surface. E.g.:

	Reflective(WHITE)        // perfectly reflective, looks like shiny metal
	Reflective(WHITE.EV(-1)) // 50% reflective, looks like darker metal
	Reflective(RED)          // Reflects only red, looks like metal in transparent red candy-wrap.

#### Example:

```go
doc.Show(
shape.NewSphere(1, Reflective(WHITE.EV(-1))).Transl(Vec{0, 0.5, 0}),
)
```

![fig](/doc/ExampleReflective.jpg)
## <a name="Refractive">func</a> [Refractive](./material.go#L45)
``` go
func Refractive(n1, n2 float64) Material
```
Refractive material with index of refraction n1 outside and n2 inside.
E.g.:

	Refractive(1, 1.5) // glass in air
	Refractive(1.5, 1) // air in glass

#### Example:

```go
doc.Show(
shape.NewSphere(1, Refractive(1, 1.5)).Transl(Vec{0, 0.5, 0}),
)
```

![fig](/doc/ExampleRefractive.jpg)
## <a name="ShadeNormal">func</a> [ShadeNormal](./material.go#L134)
``` go
func ShadeNormal(dir Vec) Material
```
ShadeNormal is a debug shader that colors according to the normal vector projected on dir.

## <a name="Shiny">func</a> [Shiny](./material.go#L115)
``` go
func Shiny(c Color, reflectivity float64) Material
```
Shiny is shorthand for Blend-ing diffuse + reflection, e.g.:
Shiny(WHITE, 0.1) // a white billiard ball, 10% specular reflection

## <a name="Waves">func</a> [Waves](./procedural.go#L94)
``` go
func Waves(seed int, n int, K Vec, col func(float64) Material) Material
```

## <a name="FlatColor">type</a> [FlatColor](./flat.go#L12-L14)
``` go
type FlatColor struct {
    // contains filtered or unexported fields
}
```

### <a name="Flat">func</a> [Flat](./flat.go#L8)
``` go
func Flat(c br.Color) *FlatColor
```
A Flat material always returns the same color.
Useful for debugging, or for rare cases like
a computer screen or other extended, dimly luminous surfaces.

#### Example:

```go
doc.Show(
shape.NewSphere(1, Flat(WHITE)).Transl(Vec{0, 0.5, 0}),
)
```

![fig](/doc/ExampleFlat.jpg)

### <a name="FlatColor.At">func</a> (\*FlatColor) [At](./flat.go#L20)
``` go
func (s *FlatColor) At(_ br.Vec) br.Color
```

### <a name="FlatColor.Shade">func</a> (\*FlatColor) [Shade](./flat.go#L16)
``` go
func (s *FlatColor) Shade(_ *br.Ctx, _ *br.Env, _ int, _ *br.Ray, _ br.Fragment) br.Color
```

## <a name="ImgTex">type</a> [ImgTex](./texture.go#L23-L26)
``` go
type ImgTex struct {
    // contains filtered or unexported fields
}
```

### <a name="NewImgTex">func</a> [NewImgTex](./texture.go#L19)
``` go
func NewImgTex(img raster.Image, p0, pu, pv Vec) *ImgTex
```

### <a name="ImgTex.At">func</a> (\*ImgTex) [At](./texture.go#L33)
``` go
func (c *ImgTex) At(pos Vec) Color
```

### <a name="ImgTex.Shade">func</a> (\*ImgTex) [Shade](./texture.go#L29)
``` go
func (c *ImgTex) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color
```
TODO: remove?

## <a name="ShadeDir">type</a> [ShadeDir](./material.go#L15)
``` go
type ShadeDir func(dir Vec) Color
```
ShadeDir returns a color based on the direction of a ray.
Used for shading the ambient background, E.g., the sky.

### <a name="ShadeDir.Shade">func</a> (ShadeDir) [Shade](./material.go#L17)
``` go
func (s ShadeDir) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color
```

## <a name="Texture">type</a> [Texture](./texture.go#L15-L17)
``` go
type Texture interface {
    At(Vec) Color
}
```

# light
`import "github.com/barnex/bruteray/light"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package light implements various types of light sources.
They all implement br.Light.

## <a name="pkg-imports">Imported Packages</a>

- [github.com/barnex/bruteray/br](./../br)
- [github.com/barnex/bruteray/mat](./../mat)
- [github.com/barnex/bruteray/shape](./../shape)

## <a name="pkg-index">Index</a>
* [func DirLight(pos Vec, intensity Color) Light](#DirLight)
* [func PointLight(pos Vec, intensity Color) Light](#PointLight)
* [func RectLight(pos Vec, rx, ry, rz float64, c Color) Light](#RectLight)
* [func Sphere(pos Vec, radius float64, intensity Color) Light](#Sphere)

#### <a name="pkg-files">Package files</a>
[light.go](./light.go) 

## <a name="DirLight">func</a> [DirLight](./light.go#L20)
``` go
func DirLight(pos Vec, intensity Color) Light
```
Directed light source without fall-off.
Position should be far away from the scene (indicates a direction)

## <a name="PointLight">func</a> [PointLight](./light.go#L35)
``` go
func PointLight(pos Vec, intensity Color) Light
```
Point light source (with fall-off).

## <a name="RectLight">func</a> [RectLight](./light.go#L93)
``` go
func RectLight(pos Vec, rx, ry, rz float64, c Color) Light
```

## <a name="Sphere">func</a> [Sphere](./light.go#L52)
``` go
func Sphere(pos Vec, radius float64, intensity Color) Light
```
Spherical light source.
Throws softer shadows than an point source and is visible in specular reflections.
TODO: nearby samples must limit their intensity to the analytical value for that limit.

# transf
`import "github.com/barnex/bruteray/transf"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package transf provides affine transformations on objects, like rotations.

## <a name="pkg-imports">Imported Packages</a>

- [github.com/barnex/bruteray/br](./../br)

## <a name="pkg-index">Index</a>
* [func Transf(o CSGObj, T \*Matrix4) CSGObj](#Transf)
* [func TransfNonCSG(o Obj, T \*Matrix4) Obj](#TransfNonCSG)

#### <a name="pkg-files">Package files</a>
[transf.go](./transf.go) 

## <a name="Transf">func</a> [Transf](./transf.go#L8)
``` go
func Transf(o CSGObj, T *Matrix4) CSGObj
```
Transf returns a transformed version of the object.
TODO: also for non-csg?

## <a name="TransfNonCSG">func</a> [TransfNonCSG](./transf.go#L42)
``` go
func TransfNonCSG(o Obj, T *Matrix4) Obj
```
TODO: rename

- - -
Generated by a modified [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)
