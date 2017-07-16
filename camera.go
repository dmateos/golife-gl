package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	position *mgl32.Vec3
}

func NewCamera() *Camera {
	c := Camera{}
	c.position = &mgl32.Vec3{0, 0, 0}
	return &c
}

func (c *Camera) GetMatrix() [16]float32 {
	return [16]float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}
