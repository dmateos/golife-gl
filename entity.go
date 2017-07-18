package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Entity struct {
	pos mgl32.Vec3
}

func NewEntity() *Entity {
	e := Entity{}
	return &e
}

func (e *Entity) GetTranslationMatrix() mgl32.Mat4 {
	return mgl32.Translate3D(e.pos.X(), e.pos.Y(), e.pos.Z())
}
