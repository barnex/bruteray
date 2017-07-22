/*
Bruteray is a brute-force ray tracer. More precisely it implements bi-directional path tracing, a ray tracing method that:
	- produces very realistic images
	- is relatively simple to implement
	- but uses a lot of compute power

TODO

	- Refactor: Object(SomeShape(...), material) -> SomeShape(..., material)
	- Slab not correct if ray starts inside (happens often with ObjAnd!)
	- Box must handle div by zero
*/
package bruteray
