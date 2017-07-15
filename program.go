package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"log"
	"strings"
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

func (p *Program) Status() bool {
	var status int32
	gl.GetProgramiv(p.programID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(p.programID, gl.INFO_LOG_LENGTH, &logLength)
		logs := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(p.programID, logLength, nil, gl.Str(logs))
		log.Print(logs)
		return false
	}
	return true
}
