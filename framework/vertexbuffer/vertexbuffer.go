package vertexbuffer

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// VertexBuffer structure
type VertexBuffer struct {
	rID uint32
}

// New - Generate VBO
func (v *VertexBuffer) New(data []float32, size int) {
	gl.GenBuffers(1, &v.rID)
	gl.BindBuffer(gl.ARRAY_BUFFER, v.rID)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(data), gl.STATIC_DRAW)
	fmt.Println(" -- Geneerated VBO ID: ", v.rID)
	fmt.Println(" -- VBO (", v.rID, ")\n\tData: ", data)
}

// Free - Delete buffer
func (v *VertexBuffer) Free() {
	gl.DeleteBuffers(1, &v.rID)
}

// Bind - buffer
func (v *VertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, v.rID)
}

// Unbind buffer
func (v *VertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

