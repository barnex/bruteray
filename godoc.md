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

Package shape implements various shapes and objects.

## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj](#Cyl)
* [func NCyl(dir int, diam float64, m br.Material) \*cyl](#NCyl)
* [func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj](#OldBox)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Sheet(dir Vec, off float64, m Material) Obj](#Sheet)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [type Box](#Box)
  * [func NewBox(w, h, d float64, m Material) \*Box](#NewBox)
  * [func (s \*Box) Center() Vec](#Box.Center)
  * [func (s \*Box) Corner(x, y, z int) Vec](#Box.Corner)
  * [func (s \*Box) Hit1(r \*Ray, f \*[]Fragment)](#Box.Hit1)
  * [func (s \*Box) HitAll(r \*Ray, f \*[]Fragment)](#Box.HitAll)
  * [func (s \*Box) Inside(v Vec) bool](#Box.Inside)
  * [func (s \*Box) Normal(p Vec) Vec](#Box.Normal)
  * [func (s \*Box) Transl(d Vec) \*Box](#Box.Transl)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewBox](#example_NewBox)
* [NewSphere](#example_NewSphere)

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="Cube">func</a> [Cube](./box.go#L53)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cyl">func</a> [Cyl](./quad.go#L10)
``` go
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).

## <a name="NCyl">func</a> [NCyl](./cylinder.go#L5)
``` go
func NCyl(dir int, diam float64, m br.Material) *cyl
```

## <a name="OldBox">func</a> [OldBox](./box.go#L45)
``` go
func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj
```
TODO rm

## <a name="Quad">func</a> [Quad](./quad.go#L19)
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

## <a name="Sheet">func</a> [Sheet](./sheet.go#L5)
``` go
func Sheet(dir Vec, off float64, m Material) Obj
```

## <a name="Slab">func</a> [Slab](./slab.go#L5)
``` go
func Slab(dir Vec, off1, off2 float64, m Material) CSGObj
```

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
NewBox(0.5, 1, 0.5, mat.Diffuse(RED)).Transl(Vec{-1.5, 0.5, 0}),
NewBox(1, 0.5, 1, mat.Diffuse(BLUE)).Transl(Vec{1.5, 0.5, 0}),
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

# shape

Package shape implements various shapes and objects.

## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj](#Cyl)
* [func NCyl(dir int, diam float64, m br.Material) \*cyl](#NCyl)
* [func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj](#OldBox)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Sheet(dir Vec, off float64, m Material) Obj](#Sheet)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [type Box](#Box)
  * [func NewBox(w, h, d float64, m Material) \*Box](#NewBox)
  * [func (s \*Box) Center() Vec](#Box.Center)
  * [func (s \*Box) Corner(x, y, z int) Vec](#Box.Corner)
  * [func (s \*Box) Hit1(r \*Ray, f \*[]Fragment)](#Box.Hit1)
  * [func (s \*Box) HitAll(r \*Ray, f \*[]Fragment)](#Box.HitAll)
  * [func (s \*Box) Inside(v Vec) bool](#Box.Inside)
  * [func (s \*Box) Normal(p Vec) Vec](#Box.Normal)
  * [func (s \*Box) Transl(d Vec) \*Box](#Box.Transl)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewBox](#example_NewBox)
* [NewSphere](#example_NewSphere)

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="Cube">func</a> [Cube](./box.go#L53)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cyl">func</a> [Cyl](./quad.go#L10)
``` go
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).

## <a name="NCyl">func</a> [NCyl](./cylinder.go#L5)
``` go
func NCyl(dir int, diam float64, m br.Material) *cyl
```

## <a name="OldBox">func</a> [OldBox](./box.go#L45)
``` go
func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj
```
TODO rm

## <a name="Quad">func</a> [Quad](./quad.go#L19)
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

## <a name="Sheet">func</a> [Sheet](./sheet.go#L5)
``` go
func Sheet(dir Vec, off float64, m Material) Obj
```

## <a name="Slab">func</a> [Slab](./slab.go#L5)
``` go
func Slab(dir Vec, off1, off2 float64, m Material) CSGObj
```

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
NewBox(0.5, 1, 0.5, mat.Diffuse(RED)).Transl(Vec{-1.5, 0.5, 0}),
NewBox(1, 0.5, 1, mat.Diffuse(BLUE)).Transl(Vec{1.5, 0.5, 0}),
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

# shape

Package shape implements various shapes and objects.

## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj](#Cyl)
* [func NCyl(dir int, diam float64, m br.Material) \*cyl](#NCyl)
* [func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj](#OldBox)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Sheet(dir Vec, off float64, m Material) Obj](#Sheet)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [type Box](#Box)
  * [func NewBox(w, h, d float64, m Material) \*Box](#NewBox)
  * [func (s \*Box) Center() Vec](#Box.Center)
  * [func (s \*Box) Corner(x, y, z int) Vec](#Box.Corner)
  * [func (s \*Box) Hit1(r \*Ray, f \*[]Fragment)](#Box.Hit1)
  * [func (s \*Box) HitAll(r \*Ray, f \*[]Fragment)](#Box.HitAll)
  * [func (s \*Box) Inside(v Vec) bool](#Box.Inside)
  * [func (s \*Box) Normal(p Vec) Vec](#Box.Normal)
  * [func (s \*Box) Transl(d Vec) \*Box](#Box.Transl)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewBox](#example_NewBox)
* [NewSphere](#example_NewSphere)

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="Cube">func</a> [Cube](./box.go#L53)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cyl">func</a> [Cyl](./quad.go#L10)
``` go
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).

## <a name="NCyl">func</a> [NCyl](./cylinder.go#L5)
``` go
func NCyl(dir int, diam float64, m br.Material) *cyl
```

## <a name="OldBox">func</a> [OldBox](./box.go#L45)
``` go
func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj
```
TODO rm

## <a name="Quad">func</a> [Quad](./quad.go#L19)
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

## <a name="Sheet">func</a> [Sheet](./sheet.go#L5)
``` go
func Sheet(dir Vec, off float64, m Material) Obj
```

## <a name="Slab">func</a> [Slab](./slab.go#L5)
``` go
func Slab(dir Vec, off1, off2 float64, m Material) CSGObj
```

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
NewBox(0.5, 1, 0.5, mat.Diffuse(RED)).Transl(Vec{-1.5, 0.5, 0}),
NewBox(1, 0.5, 1, mat.Diffuse(BLUE)).Transl(Vec{1.5, 0.5, 0}),
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

# shape

Package shape implements various shapes and objects.

## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj](#Cyl)
* [func NCyl(dir int, diam float64, m br.Material) \*cyl](#NCyl)
* [func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj](#OldBox)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Sheet(dir Vec, off float64, m Material) Obj](#Sheet)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [type Box](#Box)
  * [func NewBox(w, h, d float64, m Material) \*Box](#NewBox)
  * [func (s \*Box) Center() Vec](#Box.Center)
  * [func (s \*Box) Corner(x, y, z int) Vec](#Box.Corner)
  * [func (s \*Box) Hit1(r \*Ray, f \*[]Fragment)](#Box.Hit1)
  * [func (s \*Box) HitAll(r \*Ray, f \*[]Fragment)](#Box.HitAll)
  * [func (s \*Box) Inside(v Vec) bool](#Box.Inside)
  * [func (s \*Box) Normal(p Vec) Vec](#Box.Normal)
  * [func (s \*Box) Transl(d Vec) \*Box](#Box.Transl)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewBox](#example_NewBox)
* [NewSphere](#example_NewSphere)

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="Cube">func</a> [Cube](./box.go#L53)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cyl">func</a> [Cyl](./quad.go#L10)
``` go
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).

## <a name="NCyl">func</a> [NCyl](./cylinder.go#L5)
``` go
func NCyl(dir int, diam float64, m br.Material) *cyl
```

## <a name="OldBox">func</a> [OldBox](./box.go#L45)
``` go
func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj
```
TODO rm

## <a name="Quad">func</a> [Quad](./quad.go#L19)
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

## <a name="Sheet">func</a> [Sheet](./sheet.go#L5)
``` go
func Sheet(dir Vec, off float64, m Material) Obj
```

## <a name="Slab">func</a> [Slab](./slab.go#L5)
``` go
func Slab(dir Vec, off1, off2 float64, m Material) CSGObj
```

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
NewBox(0.5, 1, 0.5, mat.Diffuse(RED)).Transl(Vec{-1.5, 0.5, 0}),
NewBox(1, 0.5, 1, mat.Diffuse(BLUE)).Transl(Vec{1.5, 0.5, 0}),
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

# shape

Package shape implements various shapes and objects.

## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj](#Cyl)
* [func NCyl(dir int, diam float64, m br.Material) \*cyl](#NCyl)
* [func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj](#OldBox)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Sheet(dir Vec, off float64, m Material) Obj](#Sheet)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [type Box](#Box)
  * [func NewBox(w, h, d float64, m Material) \*Box](#NewBox)
  * [func (s \*Box) Center() Vec](#Box.Center)
  * [func (s \*Box) Corner(x, y, z int) Vec](#Box.Corner)
  * [func (s \*Box) Hit1(r \*Ray, f \*[]Fragment)](#Box.Hit1)
  * [func (s \*Box) HitAll(r \*Ray, f \*[]Fragment)](#Box.HitAll)
  * [func (s \*Box) Inside(v Vec) bool](#Box.Inside)
  * [func (s \*Box) Normal(p Vec) Vec](#Box.Normal)
  * [func (s \*Box) Transl(d Vec) \*Box](#Box.Transl)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewBox](#example_NewBox)
* [NewSphere](#example_NewSphere)

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="Cube">func</a> [Cube](./box.go#L53)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cyl">func</a> [Cyl](./quad.go#L10)
``` go
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).

## <a name="NCyl">func</a> [NCyl](./cylinder.go#L5)
``` go
func NCyl(dir int, diam float64, m br.Material) *cyl
```

## <a name="OldBox">func</a> [OldBox](./box.go#L45)
``` go
func OldBox(center Vec, rx, ry, rz float64, m Material) CSGObj
```
TODO rm

## <a name="Quad">func</a> [Quad](./quad.go#L19)
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

## <a name="Sheet">func</a> [Sheet](./sheet.go#L5)
``` go
func Sheet(dir Vec, off float64, m Material) Obj
```

## <a name="Slab">func</a> [Slab](./slab.go#L5)
``` go
func Slab(dir Vec, off1, off2 float64, m Material) CSGObj
```

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
NewBox(0.5, 1, 0.5, mat.Diffuse(RED)).Transl(Vec{-1.5, 0.5, 0}),
NewBox(1, 0.5, 1, mat.Diffuse(BLUE)).Transl(Vec{1.5, 0.5, 0}),
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

- - -
Generated by a modified [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)
