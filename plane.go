package main

import (
	"github.com/ojrac/opensimplex-go"
)

type Plane struct {
	entity *Entity
}

func NewPlane() *Plane {
	plane := Plane{}
	return &plane
}

func (p *Plane) LoadPlane(program *Program) {
	vertexData := NewObjFile()
	vertexData.ParseFile("obj/grid.obj")
	vertexPlane := NewVertex(vertexData.Vertex,
		vertexData.Normals,
		vertexData.VertexIndex,
		program,
	)

	p.entity = NewEntity([3]float32{0, 0, 0}, vertexPlane)

}

func (p *Plane) Draw(program *Program) {
	p.entity.Draw(program)
}

func (p *Plane) Perlin(program *Program) {
	noise := opensimplex.New()

	for x := 1; x < 30000; x += 3 {
		thing := p.entity.GetVertexData()
		thing[x] = float32(noise.Eval2(0.0, float64(x))) * 5
	}

	p.entity.vertex.RefreshBuffer(program)
}
