package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	pos *mgl32.Vec3
}

func NewCamera() *Camera {
	c := Camera{}
	c.pos = &mgl32.Vec3{0, 0, 0}
	return &c
}

func (c *Camera) Move(x, y, z float32) {
	c.pos[0] += x
	c.pos[1] += y
	c.pos[2] += z
}

func (c *Camera) GetViewMatrix() [16]float32 {
	return mgl32.Translate3D(c.pos.X(), c.pos.Y(), c.pos.Z())
}

func (c *Camera) GetPerspectiveMatrix() [16]float32 {
	return mgl32.Perspective(90.0, 1280/1024, 0.01, 100.0)
}
