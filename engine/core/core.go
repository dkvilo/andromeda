package core

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Run bundles everithing
func (andromeda *Andromeda) Run() *Andromeda {
	andromeda.init()
	andromeda.engineLoop()

	defer func() {
		andromeda.window.Destroy()
		glfw.Terminate()
	} ()

	return andromeda
}

