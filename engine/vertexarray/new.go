package vertexarray

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func (v *VertexArray) New() {
	gl.GenVertexArrays(1, &v.rendererId)
	gl.BindVertexArray(v.rendererId)
	fmt.Println(" -- Geneerated VAO ID: ", v.rendererId)
	fmt.Println(" -- VAO Binding (", v.rendererId, ")")
}
