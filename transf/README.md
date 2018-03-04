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