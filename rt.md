# Basics

In its most basic form, ray tracing works by:
  - casting rays from a camera, 
  - determining where they intersect the scene, 
  - and coloring the pixels accordingly.


## Intersection

Rays are represented in parametric form: `p = s + t*d`, where:
  - `p` is a point on the ray, 
  - `s` is the position where the ray starts,
  - `d` its direction (unit vector by convention),
  - `t` the free parameter (`t > 0`).

In order to intersect a ray with, e.g., a sphere, we simply find a `t` so that `p` lies on both the ray and the sphere:
```
(p-center)² = r²      (sphere)
p = start + t * dir   (ray)
```
yielding:
```
t = (v.d)-sqrt((v.d)²-(v²-r²))   with v = s-center
```

We must be careful to:
  - pick a positive `t` (we don't want to see objects that are behind the camera).
  - pick the smallest `t` if there are multiple solutions (we can only see the closest intersection point, the one on the back is hidden).

If we now color pixels where the ray intersects the sphere (valid `t`) white and the others black, we might get an image like:

![fig](shots/002.png)

