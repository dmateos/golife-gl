package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type ObjFile struct {
	Vertex, Normals          []float32
	VertexIndex, NormalIndex []uint32
}

func NewObjFile() *ObjFile {
	objFile := ObjFile{}
	return &objFile
}

func (o *ObjFile) Parse(d string) error {
	data := strings.Split(d, "\n")

	for _, line := range data {
		if strings.HasPrefix(line, "v ") {
			var x, y, z float32
			fmt.Sscanf(line, "v %f %f %f", &x, &y, &z)
			o.Vertex = append(o.Vertex, x)
			o.Vertex = append(o.Vertex, y)
			o.Vertex = append(o.Vertex, z)
		} else if strings.HasPrefix(line, "vn") {
			var x, y, z float32
			fmt.Sscanf(line, "vn %f %f %f", &x, &y, &z)
			o.Normals = append(o.Normals, x)
			o.Normals = append(o.Normals, y)
			o.Normals = append(o.Normals, z)
		} else if strings.HasPrefix(line, "f ") {
			var vix, viy, viz, vnx, vny, vnz uint32
			fmt.Sscanf(line, "f %d//%d %d//%d %d//%d",
				&vix, &vnx, &viy, &vny, &viz, &vnz,
			)
			o.VertexIndex = append(o.VertexIndex, vix-1)
			o.VertexIndex = append(o.VertexIndex, viy-1)
			o.VertexIndex = append(o.VertexIndex, viz-1)
			o.NormalIndex = append(o.NormalIndex, vnx-1)
			o.NormalIndex = append(o.NormalIndex, vny-1)
			o.NormalIndex = append(o.NormalIndex, vnz-1)
		}
	}

	return nil
}

func (o *ObjFile) ParseFile(filename string) error {
	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	o.Parse(string(b))
	return nil
}
