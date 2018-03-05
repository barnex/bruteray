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


# br
`import "github.com/barnex/bruteray/br"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
Bruteray is a ray tracer that does bi-directional path tracing.

## <a name="pkg-imports">Imported Packages</a>

No packages beyond the Go standard library are imported.

## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func DiaCircle(rng \*rand.Rand) (x, y float64)](#DiaCircle)
* [func DiaHex(rng \*rand.Rand) (x, y float64)](#DiaHex)
* [func EV(exp float64) float64](#EV)
* [type CSGObj](#CSGObj)
* [type Color](#Color)
  * [func (c Color) Add(b Color) Color](#Color.Add)
  * [func (c Color) At(\_ Vec) Color](#Color.At)
  * [func (c Color) EV(ev float64) Color](#Color.EV)
  * [func (c Color) MAdd(s float64, b Color) Color](#Color.MAdd)
  * [func (c \*Color) Max() float64](#Color.Max)
  * [func (c Color) Mul(s float64) Color](#Color.Mul)
  * [func (c Color) Mul3(b Color) Color](#Color.Mul3)
  * [func (c Color) RGBA() (r, g, b, a uint32)](#Color.RGBA)
  * [func (c Color) Shade(\_ \*Ctx, \_ \*Env, \_ int, \_ \*Ray, \_ Fragment) Color](#Color.Shade)
* [type Ctx](#Ctx)
  * [func NewCtx(seed int) \*Ctx](#NewCtx)
  * [func (c \*Ctx) GetFrags() \*[]Fragment](#Ctx.GetFrags)
  * [func (c \*Ctx) GetRay(start, dir Vec) \*Ray](#Ctx.GetRay)
  * [func (c \*Ctx) PutFrags(fb \*[]Fragment)](#Ctx.PutFrags)
  * [func (c \*Ctx) PutRay(r \*Ray)](#Ctx.PutRay)
* [type Env](#Env)
  * [func NewEnv() \*Env](#NewEnv)
  * [func (e \*Env) Add(o ...Obj)](#Env.Add)
  * [func (e \*Env) AddInvisibleLight(l ...Light)](#Env.AddInvisibleLight)
  * [func (e \*Env) AddLight(l ...Light)](#Env.AddLight)
  * [func (e \*Env) Occludes(ctx \*Ctx, r \*Ray, endpoint float64) bool](#Env.Occludes)
  * [func (e \*Env) SetAmbient(m Material)](#Env.SetAmbient)
  * [func (e \*Env) Shade(ctx \*Ctx, r \*Ray, N int, who []Obj) Color](#Env.Shade)
  * [func (e \*Env) ShadeAll(ctx \*Ctx, r \*Ray, N int) Color](#Env.ShadeAll)
  * [func (e \*Env) ShadeNonLum(ctx \*Ctx, r \*Ray, N int) Color](#Env.ShadeNonLum)
* [type Fragment](#Fragment)
  * [func (frag Fragment) Shade(ctx \*Ctx, e \*Env, recursion int, r \*Ray) Color](#Fragment.Shade)
* [type Insider](#Insider)
* [type Light](#Light)
* [type Material](#Material)
* [type Matrix3](#Matrix3)
* [type Matrix4](#Matrix4)
  * [func RotX4(θ float64) \*Matrix4](#RotX4)
  * [func RotY4(θ float64) \*Matrix4](#RotY4)
  * [func RotZ4(θ float64) \*Matrix4](#RotZ4)
  * [func Transl4(d Vec) \*Matrix4](#Transl4)
  * [func UnitMatrix4() \*Matrix4](#UnitMatrix4)
  * [func (a \*Matrix4) Inv() \*Matrix4](#Matrix4.Inv)
  * [func (a \*Matrix4) Mul(b \*Matrix4) \*Matrix4](#Matrix4.Mul)
  * [func (a \*Matrix4) String() string](#Matrix4.String)
  * [func (T \*Matrix4) TransfDir(v Vec) Vec](#Matrix4.TransfDir)
  * [func (T \*Matrix4) TransfPoint(v Vec) Vec](#Matrix4.TransfPoint)
* [type Obj](#Obj)
* [type Pool](#Pool)
  * [func (p \*Pool) Get() interface{}](#Pool.Get)
  * [func (p \*Pool) Put(v interface{})](#Pool.Put)
* [type Ray](#Ray)
  * [func (r \*Ray) At(t float64) Vec](#Ray.At)
  * [func (r \*Ray) Dir() Vec](#Ray.Dir)
  * [func (r \*Ray) SetDir(dir Vec)](#Ray.SetDir)
  * [func (r \*Ray) Transf(t \*Matrix4)](#Ray.Transf)
* [type Vec](#Vec)
  * [func RandVec(rng \*rand.Rand) Vec](#RandVec)
  * [func RandVecCos(rng \*rand.Rand, dir Vec) Vec](#RandVecCos)
  * [func (a Vec) Add(b Vec) Vec](#Vec.Add)
  * [func (a Vec) Cross(b Vec) Vec](#Vec.Cross)
  * [func (v Vec) Div(a float64) Vec](#Vec.Div)
  * [func (a Vec) Dot(b Vec) float64](#Vec.Dot)
  * [func (v Vec) Len() float64](#Vec.Len)
  * [func (v Vec) Len2() float64](#Vec.Len2)
  * [func (a Vec) MAdd(s float64, b Vec) Vec](#Vec.MAdd)
  * [func (v Vec) Mul(a float64) Vec](#Vec.Mul)
  * [func (v Vec) Mul3(a Vec) Vec](#Vec.Mul3)
  * [func (v Vec) Normalized() Vec](#Vec.Normalized)
  * [func (v Vec) Reflect(n Vec) Vec](#Vec.Reflect)
  * [func (a Vec) Sub(b Vec) Vec](#Vec.Sub)
  * [func (n Vec) Towards(d Vec) Vec](#Vec.Towards)
  * [func (v \*Vec) Transl(d Vec)](#Vec.Transl)
* [type Vec4](#Vec4)
  * [func (a Vec4) Dot(b Vec4) float64](#Vec4.Dot)

#### <a name="pkg-examples">Examples</a>
* [Matrix4.Inv](#example_Matrix4_Inv)
* [Matrix4.Mul](#example_Matrix4_Mul)

#### <a name="pkg-files">Package files</a>
[color.go](./color.go) [ctx.go](./ctx.go) [doc.go](./doc.go) [env.go](./env.go) [fragment.go](./fragment.go) [light.go](./light.go) [material.go](./material.go) [matrix.go](./matrix.go) [obj.go](./obj.go) [pool.go](./pool.go) [rand.go](./rand.go) [ray.go](./ray.go) [util.go](./util.go) [vec.go](./vec.go) 

## <a name="pkg-constants">Constants</a>
``` go
const (
    Pi  = math.Pi
    Deg = Pi / 180
)
```
``` go
const (
    X = 0
    Y = 1
    Z = 2
    W = 3
)
```
``` go
const DefaultRec = 6
```
Default recursion depth for NewEnv

## <a name="pkg-variables">Variables</a>
``` go
var (
    BLACK   = Color{0, 0, 0}
    WHITE   = Color{1, 1, 1}
    GREY    = Color{0.5, 0.5, 0.5}
    RED     = Color{1, 0, 0}
    GREEN   = Color{0, 1, 0}
    BLUE    = Color{0, 0, 1}
    CYAN    = Color{0, 1, 1}
    MAGENTA = Color{1, 0, 1}
    YELLOW  = Color{1, 1, 0}
)
```
``` go
var (
    Ex = Unit[X]
    Ey = Unit[Y]
    Ez = Unit[Z]
)
```
``` go
var Unit = Matrix3{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
```

## <a name="DiaCircle">func</a> [DiaCircle](./rand.go#L63)
``` go
func DiaCircle(rng *rand.Rand) (x, y float64)
```
DiaCircle draws a point from the unit disk.

## <a name="DiaHex">func</a> [DiaHex](./rand.go#L74)
``` go
func DiaHex(rng *rand.Rand) (x, y float64)
```
DiaHex draws a point from the unit hexagon.

## <a name="EV">func</a> [EV](./color.go#L79)
``` go
func EV(exp float64) float64
```
Exposure value, 2^exp.

## <a name="CSGObj">type</a> [CSGObj](./obj.go#L15-L25)
``` go
type CSGObj interface {
    Obj

    // HitAll calculates the intersection between the object and a ray.
    // It appends to *f a surface fragment for each intersection with the ray.
    // The fragments do not need to be sorted.
    HitAll(r *Ray, f *[]Fragment)

    // Inside returns true if point p lies inside the object.
    Inside(p Vec) bool
}
```
CSGObj is an object that can be used with Constructive Solid Geometry.

## <a name="Color">type</a> [Color](./color.go#L24-L26)
``` go
type Color struct {
    R, G, B float64
}
```
Color represents either a reflectivity or intensity.

In case of reflectivity, values should be [0..1],
1 meaning 100% reflectivity for that color.
In case of intensity, values are in W/m² and should be >= 0.

The color space is linear.

### <a name="Color.Add">func</a> (Color) [Add](./color.go#L69)
``` go
func (c Color) Add(b Color) Color
```
Adds two colors (i.e. blends them).

### <a name="Color.At">func</a> (Color) [At](./color.go#L34)
``` go
func (c Color) At(_ Vec) Color
```
At implements mat.Texture

### <a name="Color.EV">func</a> (Color) [EV](./color.go#L58)
``` go
func (c Color) EV(ev float64) Color
```
Multiplies the color by 2^ev.
E.g.:

	RED.EV(-1) // 50% reflective red, i.e. dark red.

### <a name="Color.MAdd">func</a> (Color) [MAdd](./color.go#L74)
``` go
func (c Color) MAdd(s float64, b Color) Color
```
Adds s*b to color c.

### <a name="Color.Max">func</a> (\*Color) [Max](./color.go#L83)
``` go
func (c *Color) Max() float64
```

### <a name="Color.Mul">func</a> (Color) [Mul](./color.go#L51)
``` go
func (c Color) Mul(s float64) Color
```
Multiplies the color, making it darker (s<1) or brighter (s>1).
E.g.:

	RED.Mul(0.5) // 50% reflective red, i.e. dark red.

### <a name="Color.Mul3">func</a> (Color) [Mul3](./color.go#L64)
``` go
func (c Color) Mul3(b Color) Color
```
Point-wise multiplication of two colors.
E.g.: light reflecting off a colored surface.

### <a name="Color.RGBA">func</a> (Color) [RGBA](./color.go#L40)
``` go
func (c Color) RGBA() (r, g, b, a uint32)
```
Implements color.Color.
Converts from float64 linear space to 16-bit srgb.

### <a name="Color.Shade">func</a> (Color) [Shade](./color.go#L29)
``` go
func (c Color) Shade(_ *Ctx, _ *Env, _ int, _ *Ray, _ Fragment) Color
```
Shade implements Material

## <a name="Ctx">type</a> [Ctx](./ctx.go#L11-L15)
``` go
type Ctx struct {
    Rng *rand.Rand // Random-number generator for use by one thread
    // contains filtered or unexported fields
}
```
Ctx is a thread-local context
for passing mutable state like
random number generators and allocation pools.

### <a name="NewCtx">func</a> [NewCtx](./ctx.go#L17)
``` go
func NewCtx(seed int) *Ctx
```

### <a name="Ctx.GetFrags">func</a> (\*Ctx) [GetFrags](./ctx.go#L41)
``` go
func (c *Ctx) GetFrags() *[]Fragment
```
GetFrags returns a new []Fragment, allocated from a pool.
PutFrags should be called to recycle.

### <a name="Ctx.GetRay">func</a> (\*Ctx) [GetRay](./ctx.go#L27)
``` go
func (c *Ctx) GetRay(start, dir Vec) *Ray
```
GetRay returns a new Ray, allocated from a pool.
PutRay should be called to recycle the Ray.

### <a name="Ctx.PutFrags">func</a> (\*Ctx) [PutFrags](./ctx.go#L48)
``` go
func (c *Ctx) PutFrags(fb *[]Fragment)
```
PutFrags recycles values returned by GetFrags.

### <a name="Ctx.PutRay">func</a> (\*Ctx) [PutRay](./ctx.go#L35)
``` go
func (c *Ctx) PutRay(r *Ray)
```
PutRay recycles Rays returned by GetRay.

## <a name="Env">type</a> [Env](./env.go#L10-L21)
``` go
type Env struct {
    Lights []Light // light sources

    Ambient     Fragment // Shades the background at infinity, when no object is hit
    Recursion   int      // Maximum allowed recursion depth. // TODO: rm?
    Fog         float64  // Fog distance
    IndirectFog bool     // Include fog interreflection

    Cutoff float64 // Maximum allowed brightness. Used to suppress spurious caustics. TODO rm
    // contains filtered or unexported fields
}
```
Env stores the entire environment
(all objects, light sources, ... in the scene)
as well as a random-number generator needed for iterative rendering.

### <a name="NewEnv">func</a> [NewEnv](./env.go#L25)
``` go
func NewEnv() *Env
```
NewEnv creates an empty environment
to which objects can be added later.

### <a name="Env.Add">func</a> (\*Env) [Add](./env.go#L37)
``` go
func (e *Env) Add(o ...Obj)
```
Adds an object to the scene.

### <a name="Env.AddInvisibleLight">func</a> (\*Env) [AddInvisibleLight](./env.go#L52)
``` go
func (e *Env) AddInvisibleLight(l ...Light)
```
Adds a light source to the scene.
The source itself is not visible, only its light.

### <a name="Env.AddLight">func</a> (\*Env) [AddLight](./env.go#L43)
``` go
func (e *Env) AddLight(l ...Light)
```
Adds a light source to the scene.

### <a name="Env.Occludes">func</a> (\*Env) [Occludes](./env.go#L178)
``` go
func (e *Env) Occludes(ctx *Ctx, r *Ray, endpoint float64) bool
```
Occludes returns true when an object intersects r
between t=0 and t=endpoint.
This means a light source at endpoint casts a shadow at the ray start point.

### <a name="Env.SetAmbient">func</a> (\*Env) [SetAmbient](./env.go#L57)
``` go
func (e *Env) SetAmbient(m Material)
```
Sets the background color.

### <a name="Env.Shade">func</a> (\*Env) [Shade](./env.go#L79)
``` go
func (e *Env) Shade(ctx *Ctx, r *Ray, N int, who []Obj) Color
```
Calculate intensity seen by ray, with maximum recursion depth N.
who = objs, lights, or all.

### <a name="Env.ShadeAll">func</a> (\*Env) [ShadeAll](./env.go#L65)
``` go
func (e *Env) ShadeAll(ctx *Ctx, r *Ray, N int) Color
```
Calculate intensity seen by ray,
caused by all objects including lights.
Used by specular surfaces
who make no distinction between light sources and regular objects.

### <a name="Env.ShadeNonLum">func</a> (\*Env) [ShadeNonLum](./env.go#L73)
``` go
func (e *Env) ShadeNonLum(ctx *Ctx, r *Ray, N int) Color
```
Calculate intensity seen by ray,
caused by objects but excluding lights.
Used for diffuse inter reflection
where contributions of light sources are added separately.

## <a name="Fragment">type</a> [Fragment](./fragment.go#L9-L21)
``` go
type Fragment struct {
    // The distance where the ray hit the object.
    // Used to determine the frontmost Shader.
    T float64

    // Surface normal where the ray hit the object.
    // Does not need to be normalized, does not need to point outwards.
    Norm Vec

    // Material.Shade will be called with the relevant position and normal to finally calculate the color.
    Material
    Object Obj
}
```
A Fragment is an infinitesimally small surface element.

Fragment shading is lazily evaluated:
only when the frontmost shader has been determined
will we call its Shade method. Shaders returned by
objects hidden behind others will eventually not be used.

### <a name="Fragment.Shade">func</a> (Fragment) [Shade](./fragment.go#L30)
``` go
func (frag Fragment) Shade(ctx *Ctx, e *Env, recursion int, r *Ray) Color
```
Calculate the color seen by ray.

## <a name="Insider">type</a> [Insider](./obj.go#L34-L36)
``` go
type Insider interface {
    Inside(pos Vec) bool
}
```

## <a name="Light">type</a> [Light](./light.go#L5-L14)
``` go
type Light interface {

    // Sample returns a position on the light's surface
    // (used to determine shadows for extended sources),
    // and the intensity at given target position.
    Sample(ctx *Ctx, target Vec) (pos Vec, intens Color)

    // Obj is rendered when the camera looks at the light directly.
    Obj
}
```
Light is an arbitrary kind of light source.
Implementations are in package light.

## <a name="Material">type</a> [Material](./material.go#L5-L13)
``` go
type Material interface {

    // Shade must return the color of the given surface fragment,
    // as seen by Ray r.
    // If Shade uses recursion, e.g., to calculate reflections,
    // it must pass N-1 as the new recursion depth, so that
    // recursion can eventually be terminated (by Env.Shade).
    Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color
}
```
A Material determines the color of a surface fragment.
E.g.: mate white, glossy red, ...

## <a name="Matrix3">type</a> [Matrix3](./matrix.go#L107)
``` go
type Matrix3 [3]Vec
```

## <a name="Matrix4">type</a> [Matrix4](./matrix.go#L8)
``` go
type Matrix4 [4]Vec4
```

### <a name="RotX4">func</a> [RotX4](./matrix.go#L57)
``` go
func RotX4(θ float64) *Matrix4
```

### <a name="RotY4">func</a> [RotY4](./matrix.go#L68)
``` go
func RotY4(θ float64) *Matrix4
```

### <a name="RotZ4">func</a> [RotZ4](./matrix.go#L79)
``` go
func RotZ4(θ float64) *Matrix4
```

### <a name="Transl4">func</a> [Transl4](./matrix.go#L90)
``` go
func Transl4(d Vec) *Matrix4
```

### <a name="UnitMatrix4">func</a> [UnitMatrix4](./matrix.go#L48)
``` go
func UnitMatrix4() *Matrix4
```

### <a name="Matrix4.Inv">func</a> (\*Matrix4) [Inv](./matrix.go#L32)
``` go
func (a *Matrix4) Inv() *Matrix4
```

#### Example:

```go
T := Transl4(Vec{2, 3, 4}).Mul(RotX4(Pi / 2).Mul(Transl4(Vec{1, 2, -3})).Mul(RotY4(Pi / 2)))
//fmt.Println(T)
fmt.Println(T.Mul(T.Inv()))
fmt.Println(T.Inv().Mul(T))
```

![fig](/doc/[1 0 0 0].jpg)
![fig](/doc/[0 1 0 0].jpg)
![fig](/doc/[0 0 1 0].jpg)
![fig](/doc/[0 0 0 1].jpg)
![fig](/doc/.jpg)
![fig](/doc/[1 0 0 0].jpg)
![fig](/doc/[0 1 0 0].jpg)
![fig](/doc/[0 0 1 0].jpg)
![fig](/doc/[0 0 0 1].jpg)

### <a name="Matrix4.Mul">func</a> (\*Matrix4) [Mul](./matrix.go#L10)
``` go
func (a *Matrix4) Mul(b *Matrix4) *Matrix4
```

#### Example:

```go
I := UnitMatrix4()
I = I.Mul(I)
fmt.Println(I)

T := Transl4(Vec{2, 3, 4})
fmt.Println(I.Mul(T))
fmt.Println(T.Mul(I))

```

![fig](/doc/[1 0 0 0].jpg)
![fig](/doc/[0 1 0 0].jpg)
![fig](/doc/[0 0 1 0].jpg)
![fig](/doc/[0 0 0 1].jpg)
![fig](/doc/.jpg)
![fig](/doc/[1 0 0 2].jpg)
![fig](/doc/[0 1 0 3].jpg)
![fig](/doc/[0 0 1 4].jpg)
![fig](/doc/[0 0 0 1].jpg)
![fig](/doc/.jpg)
![fig](/doc/[1 0 0 2].jpg)
![fig](/doc/[0 1 0 3].jpg)
![fig](/doc/[0 0 1 4].jpg)
![fig](/doc/[0 0 0 1].jpg)

### <a name="Matrix4.String">func</a> (\*Matrix4) [String](./matrix.go#L99)
``` go
func (a *Matrix4) String() string
```

### <a name="Matrix4.TransfDir">func</a> (\*Matrix4) [TransfDir](./matrix.go#L27)
``` go
func (T *Matrix4) TransfDir(v Vec) Vec
```

### <a name="Matrix4.TransfPoint">func</a> (\*Matrix4) [TransfPoint](./matrix.go#L22)
``` go
func (T *Matrix4) TransfPoint(v Vec) Vec
```

## <a name="Obj">type</a> [Obj](./obj.go#L5-L12)
``` go
type Obj interface {

    // Hit1 calculates the intersection between the object and a ray.
    // It appends to *f a surface fragment for at least the frontmost intersection.
    // More fragments may be added if convenient, these will be ignored later on.
    // The fragments do not need to be sorted.
    Hit1(r *Ray, f *[]Fragment)
}
```
Obj is an object that can be rendered as part of a scene.
E.g., a red sphere, a blue cube, ...

## <a name="Pool">type</a> [Pool](./pool.go#L3-L6)
``` go
type Pool struct {
    New func() interface{}
    // contains filtered or unexported fields
}
```

### <a name="Pool.Get">func</a> (\*Pool) [Get](./pool.go#L8)
``` go
func (p *Pool) Get() interface{}
```

### <a name="Pool.Put">func</a> (\*Pool) [Put](./pool.go#L17)
``` go
func (p *Pool) Put(v interface{})
```

## <a name="Ray">type</a> [Ray](./ray.go#L14-L18)
``` go
type Ray struct {
    Start Vec

    InvDir Vec // pre-calculated inverse direction for marginal speed improvements
    // contains filtered or unexported fields
}
```
A Ray is a half-line,
starting at the Start point (exclusive) and extending in direction Dir.

### <a name="Ray.At">func</a> (\*Ray) [At](./ray.go#L22)
``` go
func (r *Ray) At(t float64) Vec
```
Returns point Start + t*Dir.
t must be > 0 for the point to lie on the Ray.

### <a name="Ray.Dir">func</a> (\*Ray) [Dir](./ray.go#L3)
``` go
func (r *Ray) Dir() Vec
```

### <a name="Ray.SetDir">func</a> (\*Ray) [SetDir](./ray.go#L7)
``` go
func (r *Ray) SetDir(dir Vec)
```

### <a name="Ray.Transf">func</a> (\*Ray) [Transf](./ray.go#L34)
``` go
func (r *Ray) Transf(t *Matrix4)
```

## <a name="Vec">type</a> [Vec](./vec.go#L7)
``` go
type Vec [3]float64
```

### <a name="RandVec">func</a> [RandVec](./rand.go#L33)
``` go
func RandVec(rng *rand.Rand) Vec
```
Random unit vector.

### <a name="RandVecCos">func</a> [RandVecCos](./rand.go#L54)
``` go
func RandVecCos(rng *rand.Rand, dir Vec) Vec
```
Random unit vector, sampled with probability cos(angle with dir).
Used for diffuse inter-reflection importance sampling.

### <a name="Vec.Add">func</a> (Vec) [Add](./vec.go#L24)
``` go
func (a Vec) Add(b Vec) Vec
```

### <a name="Vec.Cross">func</a> (Vec) [Cross](./vec.go#L94)
``` go
func (a Vec) Cross(b Vec) Vec
```

### <a name="Vec.Div">func</a> (Vec) [Div](./vec.go#L55)
``` go
func (v Vec) Div(a float64) Vec
```
Scalar division.

### <a name="Vec.Dot">func</a> (Vec) [Dot](./vec.go#L36)
``` go
func (a Vec) Dot(b Vec) float64
```

### <a name="Vec.Len">func</a> (Vec) [Len](./vec.go#L65)
``` go
func (v Vec) Len() float64
```
Length (norm).

### <a name="Vec.Len2">func</a> (Vec) [Len2](./vec.go#L70)
``` go
func (v Vec) Len2() float64
```
Length squared

### <a name="Vec.MAdd">func</a> (Vec) [MAdd](./vec.go#L28)
``` go
func (a Vec) MAdd(s float64, b Vec) Vec
```

### <a name="Vec.Mul">func</a> (Vec) [Mul](./vec.go#L40)
``` go
func (v Vec) Mul(a float64) Vec
```

### <a name="Vec.Mul3">func</a> (Vec) [Mul3](./vec.go#L44)
``` go
func (v Vec) Mul3(a Vec) Vec
```

### <a name="Vec.Normalized">func</a> (Vec) [Normalized](./vec.go#L75)
``` go
func (v Vec) Normalized() Vec
```
Returns a copy of v, scaled to unit length.

### <a name="Vec.Reflect">func</a> (Vec) [Reflect](./vec.go#L90)
``` go
func (v Vec) Reflect(n Vec) Vec
```
Reflects v against the plane normal to n.

### <a name="Vec.Sub">func</a> (Vec) [Sub](./vec.go#L32)
``` go
func (a Vec) Sub(b Vec) Vec
```

### <a name="Vec.Towards">func</a> (Vec) [Towards](./vec.go#L82)
``` go
func (n Vec) Towards(d Vec) Vec
```
May invert v to assure it points towards direction d.
Used to ensure normal vectors point outwards.

### <a name="Vec.Transl">func</a> (\*Vec) [Transl](./vec.go#L48)
``` go
func (v *Vec) Transl(d Vec)
```

## <a name="Vec4">type</a> [Vec4](./vec.go#L110)
``` go
type Vec4 [4]float64
```

### <a name="Vec4.Dot">func</a> (Vec4) [Dot](./vec.go#L112)
``` go
func (a Vec4) Dot(b Vec4) float64
```

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

#### <a name="pkg-files">Package files</a>
[box.go](./box.go) [csg.go](./csg.go) [cylinder.go](./cylinder.go) [doc.go](./doc.go) [quad.go](./quad.go) [rect.go](./rect.go) [sheet.go](./sheet.go) [slab.go](./slab.go) [sphere.go](./sphere.go) [util.go](./util.go) 

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
  * [func NewImgTex(img raster.Image, mapper UVMapper) \*ImgTex](#NewImgTex)
  * [func (c \*ImgTex) At(pos Vec) Color](#ImgTex.At)
  * [func (c \*ImgTex) Shade(ctx \*Ctx, e \*Env, N int, r \*Ray, frag Fragment) Color](#ImgTex.Shade)
* [type ShadeDir](#ShadeDir)
  * [func (s ShadeDir) Shade(ctx \*Ctx, e \*Env, N int, r \*Ray, frag Fragment) Color](#ShadeDir.Shade)
* [type Texture](#Texture)
* [type UVAffine](#UVAffine)
  * [func (c \*UVAffine) Map(pos Vec) (u, v float64)](#UVAffine.Map)
* [type UVCyl](#UVCyl)
  * [func (c \*UVCyl) Map(pos Vec) (u, v float64)](#UVCyl.Map)
* [type UVMapper](#UVMapper)

#### <a name="pkg-examples">Examples</a>
* [Blend](#example_Blend)
* [Checkboard](#example_Checkboard)
* [DebugShape](#example_DebugShape)
* [Diffuse](#example_Diffuse)
* [Flat](#example_Flat)
* [Reflective](#example_Reflective)
* [Refractive](#example_Refractive)
* [UVAffine](#example_UVAffine)
* [UVCyl](#example_UVCyl)

#### <a name="pkg-files">Package files</a>
[diffuse.go](./diffuse.go) [diffuse_noshadow.go](./diffuse_noshadow.go) [flat.go](./flat.go) [material.go](./material.go) [procedural.go](./procedural.go) [texture.go](./texture.go) [uvmapper.go](./uvmapper.go) 

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
mat := Blend(0.95, white, 0.05, refl)
doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
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

#### Example:

```go
m1 := Diffuse(WHITE)
m2 := Reflective(WHITE.EV(-3))
mat := Checkboard(0.1, m1, m2)
doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
```

![fig](/doc/ExampleCheckboard.jpg)
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
mat := Diffuse(WHITE)
doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
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

## <a name="Load">func</a> [Load](./texture.go#L62)
``` go
func Load(name string) (raster.Image, error)
```

## <a name="MustLoad">func</a> [MustLoad](./texture.go#L54)
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
mat := Reflective(WHITE.EV(-1))
doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
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
mat := Refractive(1, 1.5)
doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
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
mat := Flat(WHITE)
doc.Show(shape.NewSphere(1, mat).Transl(Vec{0, 0.5, 0}))
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
func NewImgTex(img raster.Image, mapper UVMapper) *ImgTex
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

## <a name="UVAffine">type</a> [UVAffine](./uvmapper.go#L22-L24)
``` go
type UVAffine struct {
    P0, Pu, Pv Vec
}
```
UVAffine maps an affine coordinate system.
Most suited to map textures on plane surfaces.

	P0 -> (0, 0)
	Pu -> (1, 0)
	Pv -> (0, 1)

Often, Pu and Pv are chosen orthogonally.

#### Example:

```go
img := MustLoad("../assets/monalisa.jpg")
cube := shape.NewBox(1/img.Aspect(), 1, 0.2, nil)
cube.Transl(Vec{0, 0.5, 0})
uvmap := &UVAffine{
P0: cube.Corner(-1, -1, 1),
Pu: cube.Corner(1, -1, 1),
Pv: cube.Corner(-1, 1, 1)}
cube.Mat = Diffuse(NewImgTex(img, uvmap))
doc.Show(cube)
```

![fig](/doc/ExampleUVAffine.jpg)

### <a name="UVAffine.Map">func</a> (\*UVAffine) [Map](./uvmapper.go#L26)
``` go
func (c *UVAffine) Map(pos Vec) (u, v float64)
```

## <a name="UVCyl">type</a> [UVCyl](./uvmapper.go#L39-L41)
``` go
type UVCyl struct {
    P0, Pu, Pv Vec
}
```
UVCyl maps a cylindrical coordinate system.

	P0: center
	Pu: point on the equator
	Pv: north pole

#### Example:

```go
img := MustLoad("../assets/earth.jpg") // cylindrical projection
r := 0.5
globe := shape.NewSphere(2*r, nil)
globe.Transl(Vec{0, r, 0})
th := -30 * Deg
uvmap := &UVCyl{
P0: Vec{0, 0, 0},
Pu: Vec{math.Sin(th), 0, -math.Cos(th)},
Pv: Vec{0, 2 * r, 0},
}
globe.Mat = Diffuse(NewImgTex(img, uvmap))
doc.Show(globe)
```

![fig](/doc/ExampleUVCyl.jpg)

### <a name="UVCyl.Map">func</a> (\*UVCyl) [Map](./uvmapper.go#L43)
``` go
func (c *UVCyl) Map(pos Vec) (u, v float64)
```

## <a name="UVMapper">type</a> [UVMapper](./uvmapper.go#L12-L14)
``` go
type UVMapper interface {
    Map(pos Vec) (u, v float64)
}
```
A UVMapper maps 3D coordinates (x,y,z) on the surface of a shape
onto 2D coordinates (u,v) suitable for indexing a texture.
(u,v) coordinates typically lie within the range [0, 1].

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
