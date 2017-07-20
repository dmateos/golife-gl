package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
)

type Window struct {
	window *glfw.Window
}

func NewWindow() *Window {
	window := Window{}
	window.windowSetup()
	return &window
}

func (w *Window) Swap() {
	w.window.SwapBuffers()
}

func (w *Window) ShouldClose() bool {
	return w.window.ShouldClose()
}

func (w *Window) GetWindow() *glfw.Window {
	return w.window
}

func (w *Window) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (w *Window) windowSetup() {
	err := glfw.Init()

	if err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	w.window, err = glfw.CreateWindow(1280, 1024, "Testing", nil, nil)

	if err != nil {
		panic(err)
	}

	w.window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	w.window.SetKeyCallback(onKey)
	w.window.MakeContextCurrent()

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
}
