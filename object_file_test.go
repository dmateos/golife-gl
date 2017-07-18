package main

import (
	"testing"
)

var data = `
v 1.000000 -1.000000 -1.000000\n
v 1.000000 -1.000000 1.000000\n
vn 0.0000 -1.0000 0.0000\n
vn 0.0000 1.0000 0.0000\n
f 2//1 4//1 1//1
f 8//2 6//2 5//2
`

func TestObjFileParse(t *testing.T) {
	objFile := NewObjFile()
	status := objFile.Parse(data)

	if status != nil {
		t.Error("parser returned error")
	}

	if len(objFile.Vertex) != 6 {
		t.Error("object vertex count incorrect")
	}

	if len(objFile.Normals) != 6 {
		t.Error("object vertex normal count incorrect")
	}

	if len(objFile.VertexIndex) != 6 {
		t.Error("object vertex index count incorrect")
	}

	if len(objFile.NormalIndex) != 6 {
		t.Error("object vertex index count incorrect")
	}

	if objFile.Vertex[1] != -1 || objFile.VertexIndex[1] != 3 {
		t.Error("vertex data does not match test data")
	}

	if objFile.Normals[0] != 0 || objFile.NormalIndex[0] != 0 {
		t.Error("normal data does not match test data")
	}
}
