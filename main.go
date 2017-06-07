package main 

import (
  "runtime"
  "github.com/go-gl/glfw/v3.2/glfw"
  "github.com/go-gl/gl/v3.3-core/gl"
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

  window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
  if err != nil {
    panic(err)
  }

  window.MakeContextCurrent()

  err = gl.Init()
  if err != nil {
    panic(err)
  }

  for !window.ShouldClose() {
    window.SwapBuffers()
    glfw.PollEvents()
  }
}
