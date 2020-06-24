package vertexarray

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VertexArray struct {
	rendererId uint32
}

func (v *VertexArray) Free() {
	gl.DeleteVertexArrays(1, &v.rendererId)
}

func (v *VertexArray) Bind() {
	gl.BindVertexArray(v.rendererId)
}

func (v *VertexArray) Unbind() {
	gl.BindVertexArray(0)
}
