package components

import (
	"fmt"

	"github.com/dkvilo/andromeda/framework/entity"
	"github.com/go-gl/mathgl/mgl32"
)

// Transform Structure
type Transform struct {
	Label              string
	Container          *entity.Entity
	Position, Rotation mgl32.Vec3
}

// NewTransform constructor
func NewTransform(container *entity.Entity) *Transform {
	return &Transform{
		Container: container,
	}
}

// Update method for Transform ...
func (t *Transform) Update(time, elapsed float64) {
	fmt.Println("Transform", t.Label, ":\n  - Position:", t.Position, "\n  - Rotation:", t.Rotation)
}
