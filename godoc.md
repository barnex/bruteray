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
* [func Box(center Vec, rx, ry, rz float64, m Material) CSGObj](#Box)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj](#Cyl)
* [func NBox(w, h, d float64, m Material) \*box](#NBox)
* [func NCyl(dir int, diam float64, m br.Material) \*cyl](#NCyl)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Sheet(dir Vec, off float64, m Material) Obj](#Sheet)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewSphere](#example_NewSphere)

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="Box">func</a> [Box](./box.go#L44)
``` go
func Box(center Vec, rx, ry, rz float64, m Material) CSGObj
```

## <a name="Cube">func</a> [Cube](./box.go#L52)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cyl">func</a> [Cyl](./quad.go#L10)
``` go
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).

## <a name="NBox">func</a> [NBox](./box.go#L10)
``` go
func NBox(w, h, d float64, m Material) *box
```
NBox constructs a box with given width, depth and height.

## <a name="NCyl">func</a> [NCyl](./cylinder.go#L5)
``` go
func NCyl(dir int, diam float64, m br.Material) *cyl
```

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
e := NewEnv()
sphere := NewSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
e.Add(sphere)
doc.Example(e)
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
* [func Box(center Vec, rx, ry, rz float64, m Material) CSGObj](#Box)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj](#Cyl)
* [func NBox(w, h, d float64, m Material) \*box](#NBox)
* [func NCyl(dir int, diam float64, m br.Material) \*cyl](#NCyl)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Sheet(dir Vec, off float64, m Material) Obj](#Sheet)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewSphere](#example_NewSphere)

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="Box">func</a> [Box](./box.go#L44)
``` go
func Box(center Vec, rx, ry, rz float64, m Material) CSGObj
```

## <a name="Cube">func</a> [Cube](./box.go#L52)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cyl">func</a> [Cyl](./quad.go#L10)
``` go
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).

## <a name="NBox">func</a> [NBox](./box.go#L10)
``` go
func NBox(w, h, d float64, m Material) *box
```
NBox constructs a box with given width, depth and height.

## <a name="NCyl">func</a> [NCyl](./cylinder.go#L5)
``` go
func NCyl(dir int, diam float64, m br.Material) *cyl
```

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
e := NewEnv()
sphere := NewSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
e.Add(sphere)
doc.Example(e)
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
* [func Box(center Vec, rx, ry, rz float64, m Material) CSGObj](#Box)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj](#Cyl)
* [func NBox(w, h, d float64, m Material) \*box](#NBox)
* [func NCyl(dir int, diam float64, m br.Material) \*cyl](#NCyl)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Sheet(dir Vec, off float64, m Material) Obj](#Sheet)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewSphere](#example_NewSphere)

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="Box">func</a> [Box](./box.go#L44)
``` go
func Box(center Vec, rx, ry, rz float64, m Material) CSGObj
```

## <a name="Cube">func</a> [Cube](./box.go#L52)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cyl">func</a> [Cyl](./quad.go#L10)
``` go
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).

## <a name="NBox">func</a> [NBox](./box.go#L10)
``` go
func NBox(w, h, d float64, m Material) *box
```
NBox constructs a box with given width, depth and height.

## <a name="NCyl">func</a> [NCyl](./cylinder.go#L5)
``` go
func NCyl(dir int, diam float64, m br.Material) *cyl
```

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
e := NewEnv()
sphere := NewSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
e.Add(sphere)
doc.Example(e)
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
* [func Box(center Vec, rx, ry, rz float64, m Material) CSGObj](#Box)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj](#Cyl)
* [func NBox(w, h, d float64, m Material) \*box](#NBox)
* [func NCyl(dir int, diam float64, m br.Material) \*cyl](#NCyl)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Sheet(dir Vec, off float64, m Material) Obj](#Sheet)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewSphere](#example_NewSphere)

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="Box">func</a> [Box](./box.go#L44)
``` go
func Box(center Vec, rx, ry, rz float64, m Material) CSGObj
```

## <a name="Cube">func</a> [Cube](./box.go#L52)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cyl">func</a> [Cyl](./quad.go#L10)
``` go
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).

## <a name="NBox">func</a> [NBox](./box.go#L10)
``` go
func NBox(w, h, d float64, m Material) *box
```
NBox constructs a box with given width, depth and height.

## <a name="NCyl">func</a> [NCyl](./cylinder.go#L5)
``` go
func NCyl(dir int, diam float64, m br.Material) *cyl
```

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
e := NewEnv()
sphere := NewSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
e.Add(sphere)
doc.Example(e)
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
* [func Box(center Vec, rx, ry, rz float64, m Material) CSGObj](#Box)
* [func Cube(center Vec, r float64, m Material) CSGObj](#Cube)
* [func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj](#Cyl)
* [func NBox(w, h, d float64, m Material) \*box](#NBox)
* [func NCyl(dir int, diam float64, m br.Material) \*cyl](#NCyl)
* [func Quad(center Vec, a Vec, b float64, m Material) CSGObj](#Quad)
* [func Rect(pos, dir Vec, rx, ry, rz float64, m Material) Obj](#Rect)
* [func Sheet(dir Vec, off float64, m Material) Obj](#Sheet)
* [func Slab(dir Vec, off1, off2 float64, m Material) CSGObj](#Slab)
* [type Sphere](#Sphere)
  * [func NewSphere(diam float64, m Material) \*Sphere](#NewSphere)
  * [func (s \*Sphere) Hit1(r \*Ray, f \*[]Fragment)](#Sphere.Hit1)
  * [func (s \*Sphere) HitAll(r \*Ray, f \*[]Fragment)](#Sphere.HitAll)
  * [func (s \*Sphere) Inside(p Vec) bool](#Sphere.Inside)
  * [func (s \*Sphere) Normal(pos Vec) Vec](#Sphere.Normal)
  * [func (s \*Sphere) Transl(d Vec) \*Sphere](#Sphere.Transl)

#### <a name="pkg-examples">Examples</a>
* [NewSphere](#example_NewSphere)

## <a name="pkg-variables">Variables</a>
``` go
var CsgAnd_ func(a, b CSGObj) CSGObj
```
TODO: remove

## <a name="Box">func</a> [Box](./box.go#L44)
``` go
func Box(center Vec, rx, ry, rz float64, m Material) CSGObj
```

## <a name="Cube">func</a> [Cube](./box.go#L52)
``` go
func Cube(center Vec, r float64, m Material) CSGObj
```

## <a name="Cyl">func</a> [Cyl](./quad.go#L10)
``` go
func Cyl(dir int, center Vec, diam, h float64, m Material) CSGObj
```
Cyl constructs a (capped) cylinder along a direction (X, Y, or Z).

## <a name="NBox">func</a> [NBox](./box.go#L10)
``` go
func NBox(w, h, d float64, m Material) *box
```
NBox constructs a box with given width, depth and height.

## <a name="NCyl">func</a> [NCyl](./cylinder.go#L5)
``` go
func NCyl(dir int, diam float64, m br.Material) *cyl
```

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
e := NewEnv()
sphere := NewSphere(1, mat.Diffuse(RED)).Transl(Vec{0, 0.5, 0})
e.Add(sphere)
doc.Example(e)
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
