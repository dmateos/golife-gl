package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"runtime"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

var camera *Camera

func onKey(w *glfw.Window, key glfw.Key, scancode int,
	action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	switch key {
	case glfw.KeyW:
		camera.Move(0, 0.00, 0.4)
	case glfw.KeyS:
		camera.Move(0, -0.00, -0.4)
	case glfw.KeyA:
		camera.Move(0.4, 0, 0)
	case glfw.KeyD:
		camera.Move(-0.4, 0, 0)
	case glfw.KeyC:
		camera.Move(0, -0.4, 0.0)
	case glfw.KeyV:
		camera.Move(0, 0.4, 0.0)
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

	window, err := glfw.CreateWindow(1280, 1024, "Testing", nil, nil)
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
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	gl.DepthFunc(gl.LESS)
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

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

	camera = NewCamera()

	vertexData := NewObjFile()
	//vertexData.Read("obj/Crate1.obj")
	vertexData.Read("obj/simple_man.obj")
	//vertexData.Read("obj/monkey.obj")
	vertex := NewVertex(vertexData.Vertex, vertexData.VertexIndex, program)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		vertex.Bind()

		program.SetUniform("camera", camera.GetViewMatrix())
		program.SetUniform("perspective", camera.GetPerspectiveMatrix())

		gl.DrawElements(
			gl.TRIANGLES,
			int32(len(vertexData.VertexIndex)),
			gl.UNSIGNED_INT,
			unsafe.Pointer(uintptr(0)),
		)
		//gl.DisableVertexAttribArray(0)

		vertex.UnBind()

		window.SwapBuffers()
		glfw.PollEvents()
	}

	vert_shader.Free()
	frag_shader.Free()
	program.Free()
}
