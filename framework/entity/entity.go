package entity

import (
	"fmt"
	"reflect"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/imdario/mergo"
)

// Component Interface
type Component interface {
	Update(time, elapsed float64)
	Render()
}

// Entity that holds entities
type Entity struct {
	Components map[string]Component
	Active     bool
	Index      uint32
	Position   mgl32.Vec3
	Rotation   mgl32.Mat3
}

// SetPosition setter
func (ent *Entity) SetPosition(pos mgl32.Vec3) {
	ent.Position = pos
}

// Update method for blah ...
func (ent *Entity) Update(time, elapsed float64) {
	for _, component := range ent.Components {
		component.Update(time, elapsed)
	}
}

func (ent *Entity) Render() {
	for _, component := range ent.Components {
		component.Render()
	}
}

// AddComponent attaches new component to the manager
func (ent *Entity) AddComponent(componentName string, component Component) *Entity {

	temp := ent.Components

	for key := range temp {
		if key == componentName {
			panic(fmt.Errorf("Component Already exists on this Object %s", key))
		}
	}

	temp = map[string]Component{
		string(componentName): component,
	}

	if err := mergo.Merge(&ent.Components, temp); err != nil {
	}

	return ent
}

// GetComponentList .
func (ent *Entity) GetComponentList() map[string]Component {
	return ent.Components
}

// GetComponent .
func (ent *Entity) GetComponent(componentName string) Component {
	for key := range ent.Components {
		if reflect.TypeOf(key) == reflect.TypeOf(componentName) {
			return ent.Components[componentName]
		}
	}
	return nil
}
