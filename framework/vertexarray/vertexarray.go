package vertexarray

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// VertexArray structure
type VertexArray struct {
	rID uint32
}

// New - Generate VAO
func (v *VertexArray) New() {
	gl.GenVertexArrays(1, &v.rID)
	gl.BindVertexArray(v.rID)
	fmt.Println(" -- Geneerated VAO ID: ", v.rID)
	fmt.Println(" -- VAO Binding (", v.rID, ")")
}

// Free - delete array
func (v *VertexArray) Free() {
	gl.DeleteVertexArrays(1, &v.rID)
}

// Bind - array
func (v *VertexArray) Bind() {
	gl.BindVertexArray(v.rID)
}

// Unbind - array
func (v *VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

