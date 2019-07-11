# transf

Package transf provides affine transformations on objects, like rotations.

## <a name="pkg-index">Index</a>
* [func Transf(o CSGObj, T \*Matrix4) CSGObj](#Transf)
* [func TransfNonCSG(o Obj, T \*Matrix4) Obj](#TransfNonCSG)

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