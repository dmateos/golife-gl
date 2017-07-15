package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Vertex struct {
	bufferData                    []float32
	vertexArrayID, vertexBufferID uint32
}

func CreateVertex(data []float32) *Vertex {
	vertex := Vertex{}
	vertex.bufferData = data
	return &vertex
}

func (v *Vertex) SetupBuffer() {
	gl.GenVertexArrays(1, &v.vertexArrayID)
	gl.BindVertexArray(v.vertexArrayID)

	gl.GenBuffers(1, &v.vertexBufferID)
	gl.BindBuffer(gl.ARRAY_BUFFER, v.vertexBufferID)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(v.bufferData)*4,
		gl.Ptr(v.bufferData),
		gl.STATIC_DRAW,
	)
}
