package main

import (
	"image/png"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	cursorFile, err := os.Open("cursor.png")
	if err != nil {
		log.Fatal(err)
	}
	defer cursorFile.Close()

	cursorImage, err := png.Decode(cursorFile)
	if err != nil {
		log.Fatal(err)
	}

	err = glfw.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()

	cursor := glfw.CreateCursor(cursorImage, 10, 10)

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	window.SetCursor(cursor)
	window.Focus()
	window.MakeContextCurrent()

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}

}
