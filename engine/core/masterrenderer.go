package core

import (
	"math"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func (andromeda *Andromeda) masterRenderer() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.7, 0.6, 0.8, 1) // Solid Clear color
	gl.ClearColor(
		float32(math.Cos(andromeda.meta.Time)),
		float32(math.Sin(andromeda.meta.Time)),
		float32(math.Cos(andromeda.meta.Time)), 0.1,
	)

	if andromeda.window.GetKey(glfw.KeyL) == 1 {
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	} else {
		gl.Disable(gl.BLEND)
	}

}
