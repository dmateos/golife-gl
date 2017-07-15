package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Program struct {
	programID uint32
	shaders   []*Shader
}

func CreateProgram() *Program {
	p := Program{}
	return &p
}

func (p *Program) Use() {
	gl.UseProgram(p.programID)
}

func (p *Program) Compile() {
	p.programID = gl.CreateProgram()
	for _, shader := range p.shaders {
		gl.AttachShader(p.programID, shader.shaderID)
	}

	gl.LinkProgram(p.programID)
}

func (p *Program) AddShader(shader *Shader) {
	p.shaders = append(p.shaders, shader)
}

func (p *Program) Free() {
	gl.DeleteProgram(p.programID)
}

func (p *Program) Status() {
	var status int32
	gl.GetProgramiv(p.programID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		return false
	}
	return true
}
