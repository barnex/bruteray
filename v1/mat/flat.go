package mat

import "github.com/barnex/bruteray/v1/br"

// A Flat material always returns the same color.
// Useful for debugging, or for rare cases like
// a computer screen or other extended, dimly luminous surfaces.
func Flat(c br.Color) *FlatColor {
	return &FlatColor{c}
}

type FlatColor struct {
	c br.Color
}

func (s *FlatColor) Shade(_ *br.Ctx, _ *br.Env, _ int, _ *br.Ray, _ br.Fragment) br.Color {
	return s.c
}

func (s *FlatColor) At(_ br.Vec) br.Color {
	return s.c
}
