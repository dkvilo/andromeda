package vertexbuffer

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func (v *VertexBuffer) New(data []float32, size int) {
	gl.GenBuffers(1, &v.rendererId)
	gl.BindBuffer(gl.ARRAY_BUFFER, v.rendererId)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(data), gl.STATIC_DRAW)
	fmt.Println(" -- Geneerated VBO ID: ", v.rendererId)
	fmt.Println(" -- VBO (", v.rendererId, ")\n\tData: ", data)
}
