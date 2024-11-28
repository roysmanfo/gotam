package gotam

import (
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func InitWindow(width, height int, title string) *glfw.Window {

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	if width == CONF_FULLSCREEN_WIDTH {
		// TODO: set maximum width
	}

	if height == CONF_FULLSCREEN_HEIGHT {
		// TODO: set maximum height
	}

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	StartWindow(window)

	return window
}
