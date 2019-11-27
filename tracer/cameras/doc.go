/*
Package cameras provides various implementations of the camera interface tracer.Camera.


Projection type and view direction

A camera implementation is characterized by
   - the kind of projection (plane, spherical, isometric, ...)
   - the camera position and view direction

Constructors return a camera located at position (0,0,0) and looking
in the -Z direction. These can then be changed via the methods
   Translate
   YawPitchRoll

Internally, projection and position are separated wrapping a transform
(change in view direction and position) around an implementation
that has a fixed view direction and only cares about the type of projection.
All constructors return a pre-wrapped camera, which can be conveniently tranformed.


Axes and handedness

The convention in this package is that the X axis points to the right of the screen,
the Y axis points up, and the Z axis points towards the viewer.

This makes the coordinate system right-handed.

All implementations in this package assume the u axis (on the image sensor)
corresponds to pointing to the right of the computer screen,
and the V axis pointing up on the screen. Image samplers (package sampler)
are responsible for following this convention.

Note that some graphics libraries use other conventions.
*/
package cameras
