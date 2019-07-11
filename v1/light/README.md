# light

Package light implements various types of light sources.
They all implement br.Light.

## <a name="pkg-index">Index</a>
* [func DirLight(pos Vec, intensity Color) Light](#DirLight)
* [func PointLight(pos Vec, intensity Color) Light](#PointLight)
* [func RectLight(pos Vec, rx, ry, rz float64, c Color) Light](#RectLight)
* [func Sphere(pos Vec, radius float64, intensity Color) Light](#Sphere)

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