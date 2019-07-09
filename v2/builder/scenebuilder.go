package builder

import (
	. "github.com/barnex/bruteray/v2/tracer"
)

type SceneBuilder struct {
	Camera   Camera // TODO: rm
	Builders []Builder
}

// TODO: rename Scene
func NewSceneBuilder() *SceneBuilder {
	s := &SceneBuilder{}
	s.Camera.Init()
	return s
}

func (b *SceneBuilder) Add(c ...Builder) {
	b.Builders = append(b.Builders, c...)
}

func (b *SceneBuilder) Build() *Scene {
	s := new(Scene)
	s.Camera = b.Camera
	s.Camera.Init()

	for _, b := range b.Builders {
		b.Init()
		switch b := b.(type) {
		case Light:
			s.AddLight(b)
		case *Tree:
			for _, ch := range b.Lights {
				s.AddLight(ch)
			}
			s.AddObject(b)
		default:
			s.AddObject(b)
		}
	}
	return s
}
