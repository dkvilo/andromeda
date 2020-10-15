package components

import (
	"fmt"

	"github.com/dkvilo/andromeda/engine/entity"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// Transform Structure
type Transform struct {
	Label string
	Container *entity.Entity
	Position, Rotation mgl32.Vec3
}

// NewTransform constructor
func NewTransform(container *entity.Entity) *Transform {
	return &Transform{
		Container: container,
	}
}

// OnRender Override
func (t *Transform) OnRender() {}

// OnStart Override
func (t *Transform) OnStart() {}

// OnUpdate method for Transform ...
func (t *Transform) OnUpdate(window *glfw.Window, time, prevTime, elapsed float64) {
	fmt.Println("Transform", t.Label, ":\n  - Position:", t.Position, "\n  - Rotation:", t.Rotation)
}
