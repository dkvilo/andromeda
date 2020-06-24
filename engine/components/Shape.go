package components

import (
	"fmt"

	"github.com/dkvilo/andromeda/engine/entity"
)

// Shape Structure
type Shape struct {
	Container *entity.Entity
}

// NewShape constructor
func NewShape(container *entity.Entity) *Shape {
	return &Shape{
		Container: container,
	}
}

// OnUpdate method for Sprite ...
func (s *Shape) OnUpdate() {
	fmt.Println("Shape: Container Component", s.Container.Components)
	fmt.Println("Shape: OnUpdate() fired")
}
