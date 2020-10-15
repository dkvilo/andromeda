package components

import (
	"log"

	"github.com/andrebq/gas"
	"github.com/dkvilo/andromeda/engine/entity"
	"github.com/dkvilo/andromeda/engine/texture"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Material Structure
type Material struct {
	Path string
	Container *entity.Entity
}

// NewMaterial constructor
func NewMaterial(container *entity.Entity) *Material {
	return &Material{
		Container: container,
	}
}

// OnStart method for Sprite ...
func (m *Material) OnStart() {
	
	textureSrc, err := gas.Abs(m.Path)
	if err != nil {
		log.Fatalln("texture path is not valid")
	}

	// Load the texture
	textureData, err := texture.Create(textureSrc, 0)
	if err != nil {
		log.Fatalln(err)
	}

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, textureData)
}

// OnRender method for Sprite ...
func (m *Material) OnRender() {}

// OnUpdate method for Sprite ...
func (m *Material) OnUpdate(window *glfw.Window, time, prevTime, elapsed float64) {}

