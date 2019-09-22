package sampler

import "fmt"

// IndexToCam maps a pixel index {ix, iy} inside an image with given width and height
// onto a u,v coordinate strictly inside the interval [0,1].
// If the image's aspect ratio width:height is not square,
// then either u or v will not span the entire [0,1] interval.
//
// Half-pixel offsets are applied so that the borders in u,v correspond
// exactly to pixel borders (not centers). This transformation is sketched below:
//
// 	             +----------------+ (u,v=1,1)
// 	             |                |
//  (x,y=-.5,-.5)+----------------+
// 	             |                |
// 	             |                |
// 	             +----------------+ (x,y=w-.5,h-.5)
// 	             |                |
// 	    (u,v=0,0)+----------------+
//
// Note that the v axis points up, while the y axis points down.
func IndexToCam(w, h int, ix, iy float64) (u, v float64) {
	W := float64(w)
	H := float64(h)

	if ix < -0.5 || iy < -0.5 || ix > W-0.5 || iy > H-0.5 {
		panic(fmt.Sprintf("IndexToCam: pixel index out of range: w=%v, h=%v, x=%v, y=%v",
			w, h, ix, iy))
	}

	u = linterp(-0.5, 0, W-0.5, 1, ix)
	v = linterp(-0.5, 0.5+0.5*(H/W), H-0.5, 0.5-0.5*(H/W), iy)
	return u, v
}

// linear interpolation
// 	x1 -> y1
// 	x2 -> y2
// 	x  -> y
func linterp(x1, y1, x2, y2, x float64) (y float64) {
	return y1 + (y2-y1)*(x-x1)/(x2-x1)
}
