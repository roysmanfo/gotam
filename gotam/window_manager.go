package gotam

import (
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	targetFPS = 60
	frameTime = time.Second / targetFPS
)

var deltaTime time.Duration

func StartWindow(window *glfw.Window) {
	var lastTime = time.Now()
	var frameCount int

	for !window.ShouldClose() {

		// update deltaTime
		now := time.Now()
		deltaTime = now.Sub(lastTime)
		lastTime = now

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update + Rendering

		window.SwapBuffers()
		glfw.PollEvents()

		// Frame rate control
		frameCount++
		elapsed := time.Since(lastTime)
		if elapsed < frameTime {
			time.Sleep(frameTime - elapsed)
		}
	}
}

func GetDeltaTime() float64 {
	return deltaTime.Seconds()
}
