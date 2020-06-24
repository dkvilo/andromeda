package vertexarray

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func (v *VertexArray) New() {
	gl.GenVertexArrays(1, &v.rendererId)
	gl.BindVertexArray(v.rendererId)
}



