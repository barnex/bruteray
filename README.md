# Brute force ray tracing

## 6: Diffuse interreflection

![fig](shots/012.jpg)

## 5: Monte Carlo integration

Shooting many secondary rays in random directions and averaging out.

![fig](shots/010.jpg)

## 4: Reflection

When a ray hits a reflective surface, a reflected ray is cast. Recursion does the rest.

![fig](shots/009.jpg)

## 3: Shadow casting

Shadows are cast whenever the line of sight between a point and the light source intersects the scene.

![fig](shots/007.jpg)

![fig](shots/008.jpg)

## 2: Complex geometries

Shapes are simply functions of (x, y, z) that return `true` whenever a point lies inside. Simple shapes can be combined by boolean operations as well as transformed. This allows for complex geometries to be easily specified.

![fig](shots/005.jpg)

![fig](shots/006.jpg)

## 1: Depth and normal vectors

We calculate each ray's exact intersection with the scene, as well as the scene's normal vector at the intersection point. Then, we apply diffuse lighting.

![fig](shots/004.jpg)

## 0: basic intersection

Whenever a ray intersects an object, we color the corresponding pixel white.

![fig](shots/002.png)
