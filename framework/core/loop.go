package core

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func (andromeda *Andromeda) engineLoop() {

	// Issue Load Context
	andromeda.OnLoadContext(andromeda)

	// Engine Loop
	for !andromeda.window.ShouldClose() {

		// Update Engine Meta information
		andromeda.updateMetaOnRender()

		// Extend Master Renderer Behaviour
		andromeda.masterRenderer()

		// Load Update Context
		andromeda.OnUpdateContext(andromeda)

		if andromeda.window.GetKey(glfw.KeyQ) == 1 {
			break
		}

		andromeda.window.SwapBuffers()
		glfw.PollEvents()
	}

}
