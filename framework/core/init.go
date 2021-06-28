package core

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func (andromeda *Andromeda) init() *Andromeda {

	if andromeda.err = glfw.Init(); andromeda.err != nil {
		log.Fatalf("GLFW Init: %s", andromeda.err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	if andromeda.Frameless {
		glfw.WindowHint(glfw.Decorated, glfw.False)
	}

	if andromeda.window, andromeda.err = glfw.CreateWindow(int(andromeda.Width), int(andromeda.Height), andromeda.Title, nil, nil); andromeda.err != nil {
		log.Fatalf("GLFW CreateWindow: %s", andromeda.err)
	}

	andromeda.window.MakeContextCurrent()

	if andromeda.err = gl.Init(); andromeda.err != nil {
		log.Fatalf("GL Init: %s", andromeda.err)
	}

	fmt.Println("GL Version:", gl.GoStr(gl.GetString(gl.VERSION)))

	andromeda.running = true

	return andromeda
}
