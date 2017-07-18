package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	pos                   mgl32.Vec3
	horizAngle, vertAngle float32
}

func NewCamera() *Camera {
	c := Camera{}
	c.pos = mgl32.Vec3{0, 0, 0}
	return &c
}

func (c *Camera) Orientation() mgl32.Mat4 {
	orientation := mgl32.HomogRotate3D(c.vertAngle, [3]float32{1, 0, 0})
	orientation = orientation.Mul4(mgl32.HomogRotate3D(c.horizAngle, [3]float32{0, 1, 0}))
	return orientation
}

func (c *Camera) OffsetOrientation(up, right float32) {
	c.horizAngle += right
	c.vertAngle += up
}

func (c *Camera) OffsetPosition(offset [3]float32) {
	c.pos = c.pos.Add(offset)
}

func (c *Camera) Forward(step float32) mgl32.Vec3 {
	vec := c.Orientation().Inv().Mul4x1([4]float32{0, 0, -step, 1})
	return mgl32.Vec3{vec.X(), vec.Y(), vec.Z()}
}

func (c *Camera) Right(step float32) mgl32.Vec3 {
	vec := c.Orientation().Inv().Mul4x1([4]float32{step, 0, 0, 1})
	return mgl32.Vec3{vec.X(), vec.Y(), vec.Z()}
}

func (c *Camera) Up(step float32) mgl32.Vec3 {
	vec := c.Orientation().Inv().Mul4x1([4]float32{0, step, 0, 1})
	return mgl32.Vec3{vec.X(), vec.Y(), vec.Z()}
}

func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	return c.Orientation().Mul4(mgl32.Translate3D(c.pos.X(), c.pos.Y(), c.pos.Z()))
}

func (c *Camera) GetPerspectiveMatrix() mgl32.Mat4 {
	return mgl32.Perspective(45.0, 1280/1024, 0.01, 100.0)
}
