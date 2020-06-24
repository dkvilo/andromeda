package sandbox

import (
	. "github.com/dkvilo/andromeda/engine/core"
	_ "image/jpeg"
	_ "image/png"
	"runtime"
)

func init() {
	runtime.LockOSThread();
}

var (
	Example_1 Example = Example{
		Andromeda {
			Width:  1080,
			Height: 720,
			Title:  "Andromeda Sandbox",
		},
	}
)
