# shape

Package shape implements various shapes and objects.

## <a name="pkg-index">Index</a>
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
  * [func NewCube(d float64, m Material) \*Box](#NewCube)
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
  * [func (s \*Sphere) Radius() float64](#Sphere.Radius)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [And](#example_And)
* [Cutout](#example_Cutout)
* [Minus](#example_Minus)
* [NewBox](#example_NewBox)
* [NewCube](#example_NewCube)
* [NewCylinder](#example_NewCylinder)
* [NewInfCylinder](#example_NewInfCylinder)
* [NewSheet](#example_NewSheet)
* [NewSphere](#example_NewSphere)
* [Or](#example_Or)

## <a name="And">func</a> [And](./csg.go#L30)
``` go
func And(a, b CSGObj) CSGObj
```
Intersection (boolean AND) of two objects.

#### Example:

```go
cube := NewBox(1, 1, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
sphere := NewSphere(1.5, mat.Diffuse(BLUE)).Transl(cube.Corner(1, 1, -1))
doc.Show(And(cube, sphere))
```

![fig](/doc/ExampleAnd.jpg)
## <a name="Cube">func</a> [Cube](./box.go#L57)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cutout">func</a> [Cutout](./csg.go#L201)
``` go
func Cutout(a CSGObj, b Insider) CSGObj
```

#### Example:

```go
cube := NewBox(1, 1, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
sphere := NewSphere(1.5, mat.Diffuse(BLUE)).Transl(cube.Corner(1, 1, -1))
doc.Show(Cutout(cube, sphere))
```

![fig](/doc/ExampleCutout.jpg)
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

#### Example:

```go
cube := NewBox(1, 1, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
sphere := NewSphere(1.5, mat.Diffuse(BLUE)).Transl(cube.Corner(1, 1, -1))
doc.Show(Minus(cube, sphere))
```

![fig](/doc/ExampleMinus.jpg)
## <a name="MultiOr">func</a> [MultiOr](./csg.go#L111)
``` go
func MultiOr(o ...CSGObj) CSGObj
```

## <a name="NewCylinder">func</a> [NewCylinder](./cylinder.go#L7)
``` go
func NewCylinder(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).
TODO: Transl

#### Example:

```go
cyl := NewCylinder(Y, Vec{0, 0.5, 0}, 1, 1, mat.Diffuse(RED))
doc.Show(cyl)
```

![fig](/doc/ExampleNewCylinder.jpg)
## <a name="NewInfCylinder">func</a> [NewInfCylinder](./quad.go#L18)
``` go
func NewInfCylinder(dir int, diam float64, m Material) *quad
```

#### Example:

```go
cyl := NewInfCylinder(Y, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
doc.Show(cyl)
```

![fig](/doc/ExampleNewInfCylinder.jpg)
## <a name="OldBox">func</a> [OldBox](./box.go#L49)
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

#### Example:

```go
cube := NewBox(1, 1, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
sphere := NewSphere(1.0, mat.Diffuse(BLUE)).Transl(cube.Corner(1, 1, -1))
doc.Show(Or(cube, sphere))
```

![fig](/doc/ExampleOr.jpg)
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

## <a name="Box">type</a> [Box](./box.go#L23-L26)
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
box := NewBox(1, 0.5, 1, mat.Diffuse(RED)).Transl(Vec{0, 0.25, 0})
doc.Show(box)
```

![fig](/doc/ExampleNewBox.jpg)### <a name="NewCube">func</a> [NewCube](./box.go#L19)
``` go
func NewCube(d float64, m Material) *Box
```

#### Example:

```go
cube := NewCube(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
doc.Show(cube)
```

![fig](/doc/ExampleNewCube.jpg)

### <a name="Box.Center">func</a> (\*Box) [Center](./box.go#L28)
``` go
func (s *Box) Center() Vec
```

### <a name="Box.Corner">func</a> (\*Box) [Corner](./box.go#L43)
``` go
func (s *Box) Corner(x, y, z int) Vec
```
Corner returns one of the box's corners:

	Corner( 1, 1, 1) -> right top  back
	Corner(-1,-1,-1) -> left bottom front
	Corner( 1,-1,-1) -> right bottom front
	...

### <a name="Box.Hit1">func</a> (\*Box) [Hit1](./box.go#L61)
``` go
func (s *Box) Hit1(r *Ray, f *[]Fragment)
```

### <a name="Box.HitAll">func</a> (\*Box) [HitAll](./box.go#L63)
``` go
func (s *Box) HitAll(r *Ray, f *[]Fragment)
```

### <a name="Box.Inside">func</a> (\*Box) [Inside](./box.go#L97)
``` go
func (s *Box) Inside(v Vec) bool
```

### <a name="Box.Normal">func</a> (\*Box) [Normal](./box.go#L103)
``` go
func (s *Box) Normal(p Vec) Vec
```

### <a name="Box.Transl">func</a> (\*Box) [Transl](./box.go#L32)
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
sheet := NewSheet(Ey, 0.1, mat.Diffuse(RED))
doc.Show(sheet)
```

![fig](/doc/ExampleNewSheet.jpg)

### <a name="Sheet.Hit1">func</a> (\*Sheet) [Hit1](./sheet.go#L15)
``` go
func (s *Sheet) Hit1(r *Ray, f *[]Fragment)
```

## <a name="Sphere">type</a> [Sphere](./sphere.go#L10-L14)
``` go
type Sphere struct {
    Center Vec

    Mat Material
    // contains filtered or unexported fields
}
```

### <a name="NewSphere">func</a> [NewSphere](./sphere.go#L6)
``` go
func NewSphere(diam float64, m Material) *Sphere
```

#### Example:

```go
sphere := NewSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
doc.Show(sphere)
```

![fig](/doc/ExampleNewSphere.jpg)

### <a name="Sphere.Hit1">func</a> (\*Sphere) [Hit1](./sphere.go#L30)
``` go
func (s *Sphere) Hit1(r *Ray, f *[]Fragment)
```

### <a name="Sphere.HitAll">func</a> (\*Sphere) [HitAll](./sphere.go#L32)
``` go
func (s *Sphere) HitAll(r *Ray, f *[]Fragment)
```

### <a name="Sphere.Inside">func</a> (\*Sphere) [Inside](./sphere.go#L20)
``` go
func (s *Sphere) Inside(p Vec) bool
```

### <a name="Sphere.Normal">func</a> (\*Sphere) [Normal](./sphere.go#L49)
``` go
func (s *Sphere) Normal(pos Vec) Vec
```

### <a name="Sphere.Radius">func</a> (\*Sphere) [Radius](./sphere.go#L16)
``` go
func (s *Sphere) Radius() float64
```

### <a name="Sphere.Transl">func</a> (\*Sphere) [Transl](./sphere.go#L25)
``` go
func (s *Sphere) Transl(d Vec) *Sphere
```