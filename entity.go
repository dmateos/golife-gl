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
