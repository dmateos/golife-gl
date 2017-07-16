package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	position *mgl32.Vec3
}

func (c *Camera) NewCamera() {
	c.position = &mgl32.Vec3{0, 0, 0}
}
