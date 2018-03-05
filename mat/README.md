# mat

Package mat implements various types of materials.

## <a name="pkg-index">Index</a>
* [func Blend(a float64, matA Material, b float64, matB Material) Material](#Blend)
* [func Bricks(stride, width float64, a, b Material) Material](#Bricks)
* [func Checkboard(stride float64, a, b Material) Material](#Checkboard)
* [func DebugShape(c Color) Material](#DebugShape)
* [func Diffuse(c Texture3D) Material](#Diffuse)
* [func Diffuse0(c Texture3D) Material](#Diffuse0)
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
* [type Image](#Image)
* [type ImgTex](#ImgTex)
  * [func NewImgTex(img raster.Image, mapper UVMapper) \*ImgTex](#NewImgTex)
  * [func (c \*ImgTex) At(pos Vec) Color](#ImgTex.At)
  * [func (c \*ImgTex) Shade(ctx \*Ctx, e \*Env, N int, r \*Ray, frag Fragment) Color](#ImgTex.Shade)
* [type ShadeDir](#ShadeDir)
  * [func Skybox(tex Image) ShadeDir](#Skybox)
  * [func (s ShadeDir) Shade(ctx \*Ctx, e \*Env, N int, r \*Ray, frag Fragment) Color](#ShadeDir.Shade)
* [type Texture3D](#Texture3D)
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

## <a name="Blend">func</a> [Blend](./material.go#L100)
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
## <a name="DebugShape">func</a> [DebugShape](./material.go#L145)
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
func Diffuse(c Texture3D) Material
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
func Diffuse0(c Texture3D) Material
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

## <a name="Load">func</a> [Load](./texture.go#L66)
``` go
func Load(name string) (raster.Image, error)
```

## <a name="MustLoad">func</a> [MustLoad](./texture.go#L58)
``` go
func MustLoad(name string) raster.Image
```

## <a name="Reflective">func</a> [Reflective](./material.go#L17)
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
## <a name="Refractive">func</a> [Refractive](./material.go#L36)
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
## <a name="ShadeNormal">func</a> [ShadeNormal](./material.go#L125)
``` go
func ShadeNormal(dir Vec) Material
```
ShadeNormal is a debug shader that colors according to the normal vector projected on dir.

## <a name="Shiny">func</a> [Shiny](./material.go#L106)
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

## <a name="Image">type</a> [Image](./texture.go#L19-L21)
``` go
type Image interface {
    AtUV(u, v float64) Color
}
```

## <a name="ImgTex">type</a> [ImgTex](./texture.go#L27-L30)
``` go
type ImgTex struct {
    // contains filtered or unexported fields
}
```

### <a name="NewImgTex">func</a> [NewImgTex](./texture.go#L23)
``` go
func NewImgTex(img raster.Image, mapper UVMapper) *ImgTex
```

### <a name="ImgTex.At">func</a> (\*ImgTex) [At](./texture.go#L37)
``` go
func (c *ImgTex) At(pos Vec) Color
```

### <a name="ImgTex.Shade">func</a> (\*ImgTex) [Shade](./texture.go#L33)
``` go
func (c *ImgTex) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color
```
TODO: remove?

## <a name="ShadeDir">type</a> [ShadeDir](./skydome.go#L7)
``` go
type ShadeDir func(dir Vec) Color
```
ShadeDir returns a color based on the direction of a ray.
Used for shading the ambient background, E.g., the sky.

### <a name="Skybox">func</a> [Skybox](./skydome.go#L13)
``` go
func Skybox(tex Image) ShadeDir
```

### <a name="ShadeDir.Shade">func</a> (ShadeDir) [Shade](./skydome.go#L9)
``` go
func (s ShadeDir) Shade(ctx *Ctx, e *Env, N int, r *Ray, frag Fragment) Color
```

## <a name="Texture3D">type</a> [Texture3D](./texture.go#L15-L17)
``` go
type Texture3D interface {
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