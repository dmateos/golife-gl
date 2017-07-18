package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"io/ioutil"
	"log"
	"strings"
)

type Shader struct {
	shaderID uint32
}

func NewShader(filePath string, shaderType uint8) *Shader {
	s := Shader{}

	switch shaderType {
	case 0:
		s.shaderID = gl.CreateShader(gl.VERTEX_SHADER)
	case 1:
		s.shaderID = gl.CreateShader(gl.FRAGMENT_SHADER)
	}

	shaderData, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal("could not read shader file")
	}

	s.compileShader(string(shaderData))
	return &s
}

func (s *Shader) Free() {
	gl.DeleteShader(s.shaderID)
}

func (s *Shader) Status() bool {
	var status int32
	gl.GetShaderiv(s.shaderID, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(s.shaderID, gl.INFO_LOG_LENGTH, &logLength)
		logs := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(s.shaderID, logLength, nil, gl.Str(logs))
		log.Print(logs)
		return false
	}
	return true
}

func (s *Shader) compileShader(shaderData string) {
	shaderCode, freeShaderCode := gl.Strs(shaderData + "\x00")
	defer freeShaderCode()

	gl.ShaderSource(s.shaderID, 1, shaderCode, nil)
	gl.CompileShader(s.shaderID)
}
