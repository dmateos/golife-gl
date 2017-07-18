package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Vertex struct {
	bufferData                                         []float32
	indexData                                          []uint32
	vertexArrayID, vertexBufferID, vertexIndexBufferID uint32
}

func NewVertex(data []float32, indexData []uint32, program *Program) *Vertex {
	vertex := Vertex{}
	vertex.bufferData = data
	vertex.indexData = indexData
	vertex.setupBuffer(program)
	return &vertex
}

func (v *Vertex) Bind() {
	gl.BindVertexArray(v.vertexArrayID)
}

func (v *Vertex) UnBind() {
	gl.BindVertexArray(0)
}

func (v *Vertex) setupBuffer(program *Program) {
	gl.GenVertexArrays(1, &v.vertexArrayID)
	v.Bind()

	gl.GenBuffers(1, &v.vertexBufferID)
	gl.BindBuffer(gl.ARRAY_BUFFER, v.vertexBufferID)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(v.bufferData)*4,
		gl.Ptr(v.bufferData),
		gl.STATIC_DRAW,
	)

	gl.GenBuffers(1, &v.vertexIndexBufferID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, v.vertexIndexBufferID)
	gl.BufferData(
		gl.ELEMENT_ARRAY_BUFFER,
		len(v.indexData)*4,
		gl.Ptr(v.indexData),
		gl.STATIC_DRAW,
	)

	gl.EnableVertexAttribArray(program.GetAttribute("vp"))
	gl.VertexAttribPointer(
		program.GetAttribute("vp"),
		3,
		gl.FLOAT,
		false,
		0,
		nil,
	)
	v.UnBind()
}
