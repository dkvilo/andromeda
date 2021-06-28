package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// CreateProgram compiles shader and links
func (p Program) CreateProgram(vertexShaderSource, fragmentShaderSource string) (Program, error) {
	vertexShader, err := Compile(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return p, err
	}

	fragmentShader, err := Compile(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return p, err
	}

	p.Data = gl.CreateProgram()

	gl.AttachShader(p.Data, vertexShader)
	gl.AttachShader(p.Data, fragmentShader)
	gl.LinkProgram(p.Data)

	var status int32
	gl.GetProgramiv(p.Data, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(p.Data, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(p.Data, logLength, nil, gl.Str(log))

		return p, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return p, nil
}

// BindFragment .
func (p *Program) BindFragment(name string) {
	gl.BindFragDataLocation(p.Data, 0, gl.Str(name+"\x00"))
}

// SetupVertexAttributeArrayAndPointer .
func (p *Program) SetupVertexAttributeArrayAndPointer(name string, size, stride int32, pointerOffset int) {
	vertAttrib := uint32(gl.GetAttribLocation(p.Data, gl.Str(name+"\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, size, gl.FLOAT, false, stride, gl.PtrOffset(pointerOffset))
}

// SetProjectionMatrix .
func (p *Program) SetProjectionMatrix(projection mgl32.Mat4) {
	gl.UniformMatrix4fv(gl.GetUniformLocation(p.Data, gl.Str("projection\x00")), 1, false, &projection[0])
}

// SetModel .
func (p *Program) SetModel(model mgl32.Mat4) {
	gl.UniformMatrix4fv(gl.GetUniformLocation(p.Data, gl.Str("model\x00")), 1, true, &model[0])
}

// SetCamera .
func (p *Program) SetCamera(camera mgl32.Mat4) {
	gl.UniformMatrix4fv(gl.GetUniformLocation(p.Data, gl.Str("camera\x00")), 1, false, &camera[0])
}

// SetUniformMat4fv .
func (p *Program) SetUniformMat4fv(name string, uniformMat mgl32.Mat4, uIndex uint8, transpose bool) {
	gl.UniformMatrix4fv(gl.GetUniformLocation(p.Data, gl.Str(name+"\x00")), 1, transpose, &uniformMat[uIndex])
}

// SetFloat32 .
func (p *Program) SetFloat32(name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(p.Data, gl.Str(name+"\x00")), value)
}

// SetVec2dF32 .
func (p *Program) SetVec2dF32(name string, value mgl32.Vec2) {
	gl.Uniform2d(gl.GetUniformLocation(p.Data, gl.Str(name+"\x00")), float64(value.X()), float64(value.Y()))
}

// SetUniform1i .
func (p *Program) SetUniform1i(name string, value int32) {
	gl.Uniform1i(gl.GetUniformLocation(p.Data, gl.Str(name+"\x00")), value)
}
