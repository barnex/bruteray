/*
Package cameras provides various implementations of the camera interface (tracer.Camera).
This interface defines one method: mapping a 2D position from the image sensor onto a Ray:

	type Camera interface{
		RayFrom(ctx *Ctx, u, v float64) *Ray
	}

The context (ctx) must be used to generate random numbes, if needed
(e.g., for lens samples).

Axes and handedness

The convention in this package is that the X asis points to the right of the screen,
the Y axis points up, and the Z axis points towards the viewer.

This makes the coordinate system right-handed.

All implementations in this package assume the u axis (on the image sensor)
corresponds to pointing to the right of the computer screen,
and the V axis pointing up on the screen. Image samplers (package sampler)
are responsible for following this convention.

Note that some graphics libraries use other conventions.
*/
package cameras
