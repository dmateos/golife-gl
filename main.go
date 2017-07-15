package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func setupProgram() {
	vert_shader := CreateShader("vertex_shader.shader")
	frag_shader := CreateShader("frag_shader.shader")
	program := CreateProgram()

	if !vert_shader.Status() || !frag_shader.Status() {
		log.Fatal("could not compile shader")
	}

	program.AddShader(vert_shader)
	program.AddShader(frag_shader)
	program.Compile()

	vert_shader.Free()
	frag_shader.Free()
	program.Free()
}

func main() {
	err := glfw.Init()

	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	setupProgram()

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
