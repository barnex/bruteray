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
