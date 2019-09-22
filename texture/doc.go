/*
Package texture provides implementations of the texture interface.
This interface is that of a 3D ("solid") texture, i.e.,
a Color defined at every point in 3D space:

	type Texture interface {
		At(p Vec) Color
	}

The position p where the Texture is evaluated is usually local the
corresponding Object (HitRecord.Local).

Solid textures can either be directly generated (often procedurally),
or obtained by mapping a 2D texture onto 3D space ("UV mapping").

Textures are used wherever a space-dependent value (Color or otherwise) is needed.

  * Materials (package material) use textures for reflectivity values.
  * Some Objects (package object) use textures to modify shapes.
	E.g.: bump mapping, offset mapping. In that case the texture's Color
	is interpeted as number rather than a color.
  * Some Material and texture implementations use textures as masks
	for blending other materials or textures. In that case colors are used as weights.


Texture2D and UV Mapping


Images and interpolation

An image can be turned into a 2D texture by interpolation, i.e.,
indexing the image via floating point coordinates (u, v).


Color and float values

Some textures are inherently scalar-valued, rather than Color-valued.
(E.g.: grayscale images or bump maps.)
We do not expose a separate interface for those. Instead, we simply
return a Color and later use only one component.  This avoids having
to implement texture manipulations for both scalar- and Color-valued
textures.

TODO: provide Adapter.


TODO

Decide the package for this interface. material, builder, both?, tracer?
*/
package texture
