package vertexbuffer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VertexBuffer struct {
	rendererId uint32
}

func (v *VertexBuffer) Free() {
	gl.DeleteBuffers(1, &v.rendererId)
}

func (v *VertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, v.rendererId)
}

func (v *VertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

