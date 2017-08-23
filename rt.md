# Basics

In its most basic form, ray tracing works by:
  - casting rays from a camera, 
  - determining where they intersect the scene, 
  - and coloring the pixels accordingly.


## Intersection

A ray is a half-line. Rays are represented in parametric form: `p = s + t*d`, where:
  - `p` is a point on the ray, 
  - `s` is the position where the ray starts,
  - `d` its direction (unit vector by convention),
  - `t` the free parameter. Different `t`s (`t > 0`) yield different points along the ray.

### Sphere
In order to intersect a ray with, e.g., a sphere, we simply find a `t` so that the corresponding point `p` lies on both the ray and the sphere:
```
(p-center)² = r²      (sphere)
p = s + t*d           (ray)
```
yielding:
```
t = (v.d)±sqrt((v.d)²-(v²-r²))   with v = s-center
```

We must be careful to:
  - pick a positive `t` (we don't want to see objects *behind* the camera).
  - pick the smallest `t` if there are multiple solutions (we can only see the closest intersection point, the one on the back is hidden).

If we now color pixels where the ray intersects the sphere (valid `t`) white and the others black, we might get an image like:

![fig](testdata/001-sphere.png)

### Sheet
Similarly for an infinite, horizontal sheet at height `y0` (e.g. a floor plane):

```
y = y0         (sheet)
y = ys + t*yd  (ray)

=> t = (y0-ys)/yd
```

