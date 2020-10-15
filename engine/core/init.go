package core

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/andrebq/gas"
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

	if andromeda.window, andromeda.err = glfw.CreateWindow(int(andromeda.Width), int(andromeda.Height), andromeda.Title, nil, nil); andromeda.err !=  nil {
		log.Fatalf("GLFW CreateWindow: %s", andromeda.err)
	}

	andromeda.window.SetDropCallback(func(window *glfw.Window, names []string) {
		fmt.Println(names)
		window.SetTitle(names[0])
	})
	
	andromeda.window.MakeContextCurrent()
	
	if andromeda.err = gl.Init(); andromeda.err != nil {
		log.Fatalf("GL Init: %s", andromeda.err)
	}
		
	fmt.Println("GL Version:", gl.GoStr(gl.GetString(gl.VERSION)))

	andromeda.running = true

	cursortSrc, err := gas.Abs("github.com/dkvilo/andromeda/resources/assets/ui/cursor.png")
	if err != nil {
		panic(err)
	}

	cursorImgFile, err := os.Open(cursortSrc)
	if err != nil {
		panic(fmt.Errorf("Image %v not found on disk", err))
	}
	
	img, _, err := image.Decode(cursorImgFile)
	if err != nil {
		panic(fmt.Errorf("texture %v not found on disk", err))
	}

	newCursor := glfw.CreateCursor(img, 0, 0)
	andromeda.window.SetCursor(newCursor)

	return andromeda
}
