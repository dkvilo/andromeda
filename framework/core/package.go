package core

import (
	"github.com/dkvilo/andromeda/framework/entity"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// AndromedaContextCallback .
type AndromedaContextCallback func(andromeda *Andromeda)

// Andromeda Basic Structure
type Andromeda struct {
	Width, Height                  float32
	Title                          string
	window                         *glfw.Window
	EntityManager                  entity.Manager
	OnLoadContext, OnUpdateContext AndromedaContextCallback
	Frameless                      bool
	running                        bool
	err                            error
	meta                           Meta
}

// GetWindow Getter
func (andromeda *Andromeda) GetWindow() *glfw.Window {
	return andromeda.window
}

// Meta dataset for andromeda internal
type Meta struct {
	Elapsed      float64
	PreviousTime float64
	Time         float64
}

// GetMeta Getter
func (andromeda *Andromeda) GetMeta() *Meta {
	return &andromeda.meta
}

// GetTime .
func (meta *Meta) GetTime() float64 {
	return meta.Time
}

// GetElapsed .
func (meta *Meta) GetElapsed() float64 {
	return meta.Elapsed
}

// GetPreviousTime .
func (meta *Meta) GetPreviousTime() float64 {
	return meta.PreviousTime
}

func (andromeda *Andromeda) updateMetaOnRender() {
	andromeda.meta.PreviousTime = andromeda.meta.Time
	andromeda.meta.Time = glfw.GetTime()
	andromeda.meta.Elapsed = andromeda.meta.Time - andromeda.meta.PreviousTime
}
