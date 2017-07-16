package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	pos *mgl32.Vec3
}

func NewCamera() *Camera {
	c := Camera{}
	c.pos = &mgl32.Vec3{0, -0.10, 0}
	return &c
}

func (c *Camera) GetMatrix() [16]float32 {
	return mgl32.Translate3D(c.pos.X(), c.pos.Y(), c.pos.Z())
}
