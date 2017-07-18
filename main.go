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

var camera *Camera

func onKey(w *glfw.Window, k glfw.Key, s int, a glfw.Action, m glfw.ModifierKey) {
	switch k {
	case glfw.KeyW:
		camera.OffsetPosition(camera.Forward(-0.5))
	case glfw.KeyS:
		camera.OffsetPosition(camera.Forward(0.5))
	case glfw.KeyA:
		camera.OffsetPosition(camera.Right(0.5))
	case glfw.KeyD:
		camera.OffsetPosition(camera.Right(-0.5))
	}
}

func handleMouse(w *glfw.Window) {
	x, y := w.GetCursorPos()
	camera.OffsetOrientation(float32(y)*0.01, float32(x)*0.01)
	w.SetCursorPos(0.0, 0.0)
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

	window, err := glfw.CreateWindow(1280, 1024, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	window.SetKeyCallback(onKey)

	window.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.ClearColor(0.11, 0.545, 0.765, 0.0)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	gl.DepthFunc(gl.LESS)
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	return window
}

func make_basic_gl_shader_program() *Program {
	vertShader := NewShader("shaders/vertex_shader.shader", 0)
	fragShader := NewShader("shaders/frag_shader.shader", 1)

	defer vertShader.Free()
	defer fragShader.Free()

	if !vertShader.Status() {
		log.Fatal("could not compile vertex shader")
	}

	if !fragShader.Status() {
		log.Fatal("could not compile fragment shader")
	}

	program := NewProgram()
	program.AddShader(vertShader)
	program.AddShader(fragShader)
	program.Compile()

	if !program.Status() {
		log.Fatal("Could not compile GL program")
	}

	return program
}

func main() {
	window := window_setup()
	program := make_basic_gl_shader_program()
	camera = NewCamera()
	vertexData := NewObjFile()

	vertexData.ParseFile("obj/simple_man.obj")
	vertex := NewVertex(vertexData.Vertex, vertexData.VertexIndex, program)

	defer glfw.Terminate()
	defer program.Free()

	program.Use()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		program.SetUniform("camera", camera.GetViewMatrix())
		program.SetUniform("perspective", camera.GetPerspectiveMatrix())

		vertex.Bind()
		vertex.Draw()
		vertex.UnBind()

		window.SwapBuffers()
		glfw.PollEvents()
		handleMouse(window)
	}
}
