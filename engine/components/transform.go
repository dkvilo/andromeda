package components

import (
	"fmt"

	"github.com/dkvilo/andromeda/engine/entity"
	"github.com/dkvilo/andromeda/engine/utils/vector"
)

//   Structure
type Transform struct {
	Container *entity.Entity
	Position vector.Vector3d
	Rotation vector.Vector3d
}

// NewTransform constructor
func NewTransform(container *entity.Entity) *Transform {
	return &Transform{
		Container: container,
	}
}

// OnUpdate method for Transform ...
func (t *Transform) OnUpdate() {
	t.Rotation.X += 2.4
	t.Rotation.Y += 2.4
	fmt.Println("Transform On Update on [", t.Container, "]")
}
