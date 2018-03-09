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

    Ambient     Material // Shades the background at infinity, when no object is hit
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

### <a name="Env.Occludes">func</a> (\*Env) [Occludes](./env.go#L177)
``` go
func (e *Env) Occludes(ctx *Ctx, r *Ray, endpoint float64) bool
```
Occludes returns true when an object intersects r
between t=0 and t=endpoint.
This means a light source at endpoint casts a shadow at the ray start point.

### <a name="Env.SetAmbient">func</a> (\*Env) [SetAmbient](./env.go#L112)
``` go
func (e *Env) SetAmbient(m Material)
```

### <a name="Env.Shade">func</a> (\*Env) [Shade](./env.go#L74)
``` go
func (e *Env) Shade(ctx *Ctx, r *Ray, N int, who []Obj) Color
```
Calculate intensity seen by ray, with maximum recursion depth N.
who = objs, lights, or all.

### <a name="Env.ShadeAll">func</a> (\*Env) [ShadeAll](./env.go#L60)
``` go
func (e *Env) ShadeAll(ctx *Ctx, r *Ray, N int) Color
```
Calculate intensity seen by ray,
caused by all objects including lights.
Used by specular surfaces
who make no distinction between light sources and regular objects.

### <a name="Env.ShadeNonLum">func</a> (\*Env) [ShadeNonLum](./env.go#L68)
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