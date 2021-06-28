package entity

import (
	"fmt"
	"reflect"

	"github.com/imdario/mergo"
)

type Manager struct {
	Entities map[string]Entity
}

func (m *Manager) RegisterEntity(entityName string, ent Entity) Entity {
	temp := m.Entities
	for key := range temp {
		if key == entityName {
			panic(fmt.Errorf("Entity is Already registered in manager %s", key))
		}
	}

	temp = map[string]Entity{
		string(entityName): ent,
	}

	if err := mergo.Merge(&m.Entities, temp); err != nil {
	}
	return ent
}

// GetEntities .
func (m *Manager) GetEntities() map[string]Entity {
	return m.Entities
}

// GetEntity .
func (m *Manager) GetEntity(entityName string) Entity {
	for key := range m.Entities {
		if reflect.TypeOf(key) == reflect.TypeOf(entityName) {
			return m.Entities[entityName]
		}
	}
	return Entity{}
}
