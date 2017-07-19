package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"unsafe"
)

type Vertex struct {
	bufferData, normalData                                             []float32
	indexData                                                          []uint32
	vertexArrayID, vertexBufferID, normalBufferID, vertexIndexBufferID uint32
}

func NewVertex(data, nData []float32, indexData []uint32, program *Program) *Vertex {
	vertex := Vertex{}
	vertex.bufferData = data
	vertex.normalData = nData
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

	// Vertex Buffer + Attrib pointer to goto shader.
	gl.GenBuffers(1, &v.vertexBufferID)
	gl.BindBuffer(gl.ARRAY_BUFFER, v.vertexBufferID)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(v.bufferData)*4,
		gl.Ptr(v.bufferData),
		gl.STATIC_DRAW,
	)

	gl.EnableVertexAttribArray(program.GetAttribute("vp"))
	gl.VertexAttribPointer(program.GetAttribute("vp"),
		3,
		gl.FLOAT,
		false,
		0,
		nil,
	)

	// Normal Buffer + Attrib pointer to goto shader.
	gl.GenBuffers(1, &v.normalBufferID)
	gl.BindBuffer(gl.ARRAY_BUFFER, v.normalBufferID)
	gl.BufferData(gl.ARRAY_BUFFER,
		len(v.normalData)*4,
		gl.Ptr(v.normalData),
		gl.STATIC_DRAW,
	)

	gl.EnableVertexAttribArray(program.GetAttribute("nm"))
	gl.VertexAttribPointer(program.GetAttribute("nm"),
		3,
		gl.FLOAT,
		false,
		0,
		nil,
	)

	// Vertex index buffer
	gl.GenBuffers(1, &v.vertexIndexBufferID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, v.vertexIndexBufferID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		len(v.indexData)*4,
		gl.Ptr(v.indexData),
		gl.STATIC_DRAW,
	)

	//Cleanup
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	v.UnBind()
}

func (v *Vertex) Draw() {
	gl.DrawElements(gl.TRIANGLES,
		int32(len(v.indexData)),
		gl.UNSIGNED_INT,
		unsafe.Pointer(uintptr(0)),
	)
}
