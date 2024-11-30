package gotam

import (
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	targetFPS = 140
	frameTime = time.Second / targetFPS
)

var deltaTime time.Duration

// this may slow down the refresh rate of the window
// when set to true, causing a loss of performance
// but could be useful during debugging sessions
var useFrameRateControl = false

func StartWindow(window *glfw.Window) {
	var lastTime = time.Now()

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
		if elapsed := time.Since(lastTime); elapsed < frameTime && useFrameRateControl {
			time.Sleep(frameTime - elapsed)
		}
	}
}

func GetDeltaTime() float64 {
	return deltaTime.Seconds()
}
