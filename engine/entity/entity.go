package entity

import (
	"fmt"

	"github.com/imdario/mergo"
)

// Component Interface
type Component interface {
	OnUpdate()
}

// Entity holds entities
type Entity struct {
	Components map[string]Component
	Active bool
}

// OnUpdate method for blah ...
func (ent *Entity) OnUpdate() {
	for key, value := range ent.Components {
		fmt.Println("key", key)
		fmt.Println("Update On", value)
		// value.OnUpdate()
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
	
	temp = map[string] Component {
		string(componentName): component,
	}

	if err := mergo.Merge(&ent.Components, temp); err != nil {
	}



	return ent
}

func (ent *Entity) GetComponent(componentName string) Component {
	// for _, comp := range ent.Components {
		for key, value := range ent.Components {
			fmt.Println("list", key, value)
			if key == componentName {
				return value
			} else {
				panic(componentName + " component not found on entity")
			}
		}
	// }
	return nil
}

