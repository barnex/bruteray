package texture_test

//
//import (
//	"github.com/barnex/bruteray/builder"
//	"github.com/barnex/bruteray/imagef/colorf"
//	. "github.com/barnex/bruteray/tracer/geom"
//	. "github.com/barnex/bruteray/material"
//	"github.com/barnex/bruteray/test"
//	. "github.com/barnex/bruteray/texture"
//)
//
//func ExampleCheckers() {
//	scene := builder.NewSceneBuilder()
//
//	tex2D := Checkers(1, 1, color.White, color.Gray(0.5))
//
//	rect := builder.NewRectangle(Flat(MapLocal(tex2D)), Vec{-1, -1, 0}, Vec{-1, 1, 0}, Vec{1, -1, 0})
//
//	scene.Add(rect)
//
//	api.Render()
//
//	// Output:
//	// exampleCheckers.png
//}
//
