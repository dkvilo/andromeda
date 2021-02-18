package sandbox

import (
	"runtime"

	core "github.com/dkvilo/andromeda/framework/core"
)

func init() {
	runtime.LockOSThread()
}

var (
	// Blank Example App
	Blank Example = Example{
		core.Andromeda{
			Width:  1080,
			Height: 720,
			Title:  "Andromeda Empty Sandbox",

			// Issue OnLoad
			SetupOnLoadContext: func(andromeda *core.Andromeda) {
				// Load Assets
			},

			// Issue Entity Update
			SetupOnUpdateContext: func(andromeda *core.Andromeda) {
				// do updates 
			},

			// Issue Draw Call
			SetupOnRenderContext: func(andromeda *core.Andromeda) {
				// draw
			},
		},
	}
)
