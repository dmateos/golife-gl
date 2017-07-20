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
	case glfw.KeyC:
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	}
}

func handleMouse(w *glfw.Window) {
	x, y := w.GetCursorPos()
	camera.OffsetOrientation(float32(y)*0.01, float32(x)*0.01)
	w.SetCursorPos(0.0, 0.0)
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
	window := NewWindow()
	program := make_basic_gl_shader_program()
	camera = NewCamera()

	window.GetWindow().SetKeyCallback(onKey)

	vertexDataMan := NewObjFile()
	vertexDataMan.ParseFile("obj/model.obj")
	vertexMan := NewVertex(vertexDataMan.Vertex,
		vertexDataMan.Normals, vertexDataMan.VertexIndex, program,
	)
	entity := NewEntity([3]float32{0, 0, 0}, vertexMan)
	entity_two := NewEntity([3]float32{20.0, 0.0, 0.0}, vertexMan)
	entity_three := NewEntity([3]float32{0.0, 0.0, 20.0}, vertexMan)
	entity_four := NewEntity([3]float32{20.0, 0.0, 20.0}, vertexMan)

	vertexDataMonk := NewObjFile()
	vertexDataMonk.ParseFile("obj/monkey.obj")
	vertexMonk := NewVertex(vertexDataMonk.Vertex,
		vertexDataMonk.Normals, vertexDataMonk.VertexIndex, program,
	)
	monkEntity := NewEntity([3]float32{10, 20.0, 10.0}, vertexMonk)

	plane := NewPlane()
	plane.LoadPlane(program)
	plane.Perlin(program)

	defer glfw.Terminate()
	defer program.Free()

	program.Use()

	for !window.ShouldClose() {
		window.Clear()
		program.SetUniform("camera", camera.GetViewMatrix())
		program.SetUniform("projection", camera.GetPerspectiveMatrix())

		plane.Draw(program)
		monkEntity.Draw(program)
		entity.Draw(program)
		entity_two.Draw(program)
		entity_three.Draw(program)
		entity_four.Draw(program)

		window.Swap()
		glfw.PollEvents()
		handleMouse(window.GetWindow())
	}
}
