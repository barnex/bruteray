package main

type Image [][]Color

func MakeImage(W, H int) Image {
	img := make(Image, H)
	for i := range img {
		img[i] = make([]Color, W)
	}
	return img
}

func (i Image) Size() (int, int) {
	return len(i[0]), len(i)
}
