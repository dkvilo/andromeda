package components

import (
	"fmt"

	"github.com/dkvilo/andromeda/framework/entity"
	"github.com/go-gl/mathgl/mgl32"
)

// Camera struct definition
type Camera struct {
	Container *entity.Entity
	position  mgl32.Mat4
}

// NewCamera constructor
func NewCamera(container *entity.Entity) *Camera {
	return &Camera{
		Container: container,
	}
}

// LookAt Default Position
func (cam *Camera) LookAt(eye, center, up mgl32.Vec3) *Camera {
	// centered position
	// mgl32.Vec3{2, 2, 2}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0}
	cam.position = mgl32.LookAtV(eye, center, up)
	return cam
}

// GetPosition Access camera position in world
func (cam *Camera) GetPosition() mgl32.Mat4 {
	return cam.position
}

// Update method for Sprite ...
func (cam *Camera) Update(time, elapsed float64) {
	fmt.Println("Camera: \n - Position:")
	fmt.Println(cam.position)
}
