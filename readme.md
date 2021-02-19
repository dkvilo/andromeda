# Andromeda Renderer

Core Renderer of andromeda Game Engine written in go, using Modern OpenGL API

## Run

```sh
  # install module deps
  $ go mod download

  # run example scenes from sandbox layer
  $ make build && make run
```

# Preview 3d scene example
<img src="https://github.com/dkvilo/andromeda/blob/develop/screen-space-reflections-low-res.gif" />


## Framework usage
```go

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


```
