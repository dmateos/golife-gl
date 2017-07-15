package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()

	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
	gl.ClearColor(0.11, 0.545, 0.765, 0.0)

	vert_shader := CreateShader("shaders/vertex_shader.shader", 0)
	frag_shader := CreateShader("shaders/frag_shader.shader", 1)

	if !vert_shader.Status() {
		log.Fatal("could not compile vertex shader")
	}

	if !frag_shader.Status() {
		log.Fatal("could not compile fragment shader")
	}

	program := CreateProgram()
	program.AddShader(vert_shader)
	program.AddShader(frag_shader)
	program.Compile()

	if !program.Status() {
		log.Fatal("Could not compile GL program")
	}

	program.Use()

	vertexBufferData := []float32{
		-1.0, -1.0, 0.0,
		1.0, -1.0, 0.0,
		0.0, 1.0, 0.0,
	}

	vertex := CreateVertex(vertexBufferData)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		vertex.Bind()

		gl.EnableVertexAttribArray(0)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.DisableVertexAttribArray(0)

		vertex.UnBind()

		window.SwapBuffers()
		glfw.PollEvents()
	}

	vert_shader.Free()
	frag_shader.Free()
	program.Free()
}
