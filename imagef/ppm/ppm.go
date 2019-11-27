// Package PPM provides support for images in the NetPBM format, as defined by https://en.wikipedia.org/wiki/Netpbm_format.
package ppm

import (
	"fmt"
	"io"

	"github.com/barnex/bruteray/imagef"
	"github.com/barnex/bruteray/imagef/colorf"
)

const ppmMaxCol = (1 << 16) - 1

//https://en.wikipedia.org/wiki/Netpbm_format
func EncodeAscii16(w io.Writer, img imagef.Image) error {
	if err := writeHeader(w, "P3", img); err != nil {
		return err
	}
	for i := range img {
		for j := range img[i] {
			c := img[i][j]
			if _, err := fmt.Fprint(w, trunc(c.R), " ", trunc(c.G), " ", trunc(c.B), " "); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintln(w); err != nil {
			return err
		}
	}
	return nil
}

func Encode48BE(w io.Writer, img imagef.Image) error {
	if err := writeHeader(w, "P6", img); err != nil {
		return err
	}

	b := [2]byte{}
	buf := b[:]
	for i := range img {
		for j := range img[i] {
			c := img[i][j]
			encodeUint16BE(buf, uint16(trunc(c.R)))
			w.Write(buf)
			encodeUint16BE(buf, uint16(trunc(c.G)))
			w.Write(buf)
			encodeUint16BE(buf, uint16(trunc(c.B)))
			w.Write(buf)
		}
	}
	return nil
}

func encodeUint16BE(buf []byte, x uint16) {
	buf[0] = byte(x >> 8)
	buf[1] = byte(x)
}

func writeHeader(w io.Writer, format string, img imagef.Image) error {
	if _, err := fmt.Fprintln(w, format); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, img.Bounds().Dx(), img.Bounds().Dy()); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, ppmMaxCol); err != nil {
		return err
	}
	return nil
}

func trunc(c float64) int {
	c = colorf.LinearToSRGB(c)
	i := int(c * ppmMaxCol)
	if i > ppmMaxCol {
		i = ppmMaxCol
	}
	if i < 0 {
		i = 0
	}
	return i
}
