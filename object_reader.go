package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type ObjFile struct {
	Vertex      []float32
	VertexIndex []int32
}

func NewObjFile() *ObjFile {
	objFile := ObjFile{}
	return &objFile
}

func (o *ObjFile) Read(filename string) {
	b, err := ioutil.ReadFile(filename)

	if err != nil {

	}

	data := strings.Split(string(b), "\n")

	for _, line := range data {
		if strings.HasPrefix(line, "v ") {
			var x, y, z float32
			fmt.Sscanf(line, "v %f %f %f", &x, &y, &z)
			o.Vertex = append(o.Vertex, x)
			o.Vertex = append(o.Vertex, y)
			o.Vertex = append(o.Vertex, z)
		} else if strings.HasPrefix(line, "f ") {
			var vix, viy, viz, vnx, vny, vnz int32
			fmt.Sscanf(line,
				"f %d//%d %d//%d %d///%d",
				&vix, &vnx, &viy, &vny, &viz, &vnz,
			)
			o.VertexIndex = append(o.VertexIndex, vix)
			o.VertexIndex = append(o.VertexIndex, viy)
			o.VertexIndex = append(o.VertexIndex, viz)
		}
	}
	fmt.Printf("%d\n", len(o.Vertex))
}