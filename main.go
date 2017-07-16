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

func onKey(w *glfw.Window, key glfw.Key, scancode int,
	action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		return
	}

	switch key {
	case glfw.KeyW:
	case glfw.KeyA:
	case glfw.KeyS:
	case glfw.KeyD:
	}
}

func window_setup() *glfw.Window {
	err := glfw.Init()

	if err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(1289, 1024, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.SetKeyCallback(onKey)

	window.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
	gl.ClearColor(0.11, 0.545, 0.765, 0.0)

	return window
}

func main() {
	defer glfw.Terminate()

	window := window_setup()

	vert_shader := NewShader("shaders/vertex_shader.shader", 0)
	frag_shader := NewShader("shaders/frag_shader.shader", 1)

	if !vert_shader.Status() {
		log.Fatal("could not compile vertex shader")
	}

	if !frag_shader.Status() {
		log.Fatal("could not compile fragment shader")
	}

	program := NewProgram()
	program.AddShader(vert_shader)
	program.AddShader(frag_shader)
	program.Compile()

	if !program.Status() {
		log.Fatal("Could not compile GL program")
	}

	program.Use()

	camera := NewCamera()

	vertexBufferData := []float32{
		-1.0, -1.0, 0.0,
		1.0, -1.0, 0.0,
		0.0, 1.0, 0.0,
	}

	vertex := NewVertex(vertexBufferData, program)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		vertex.Bind()

		program.SetUniform("camera", camera.GetMatrix())

		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		//gl.DisableVertexAttribArray(0)

		vertex.UnBind()

		window.SwapBuffers()
		glfw.PollEvents()
	}

	vert_shader.Free()
	frag_shader.Free()
	program.Free()
}
