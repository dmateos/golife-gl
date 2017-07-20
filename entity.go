package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Entity struct {
	pos    mgl32.Vec3
	vertex *Vertex
}

func NewEntity(pos [3]float32, vertex *Vertex) *Entity {
	e := Entity{}
	e.vertex = vertex
	e.pos = pos
	return &e
}

func (e *Entity) GetTranslationMatrix() mgl32.Mat4 {
	return mgl32.Translate3D(e.pos.X(), e.pos.Y(), e.pos.Z())
}

func (e Entity) GetVertexData() []float32 {
	return e.vertex.bufferData
}

func (e *Entity) Draw(program *Program) {
	e.vertex.Bind()
	program.SetUniform("transform", e.GetTranslationMatrix())
	e.vertex.Draw()
	e.vertex.UnBind()
}
