package components

import (
	"log"

	"github.com/andrebq/gas"
	"github.com/dkvilo/andromeda/framework/entity"
	"github.com/dkvilo/andromeda/framework/shader"
	"github.com/dkvilo/andromeda/framework/texture"
	"github.com/dkvilo/andromeda/framework/vertexarray"
	"github.com/dkvilo/andromeda/framework/vertexbuffer"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Renderer Structure
type Renderer struct {
	RawRenderer []float32
	Vao         vertexarray.VertexArray
	Vbo         vertexbuffer.VertexBuffer
	Program     *shader.Program
	Container   *entity.Entity
}

// NewRenderer constructor
func NewRenderer(container *entity.Entity, data []float32, program *shader.Program) *Renderer {

	instance := Renderer{
		Container:   container,
		RawRenderer: data,
		Program:     program,
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	// Create And bind
	instance.Vbo.New(data, len(data)*4)

	// Create and bind
	instance.Vao.New()

	gl.UseProgram(instance.Program.Data)
	instance.Program.SetModel(model)

	return &instance
}

var angle = 0.1
var model = mgl32.Ident4()

var camera = Camera{}

// LoadAndBindTexture .
func (r *Renderer) LoadAndBindTexture(path string) {
	textureSrc, err := gas.Abs(path)
	if err != nil {
		log.Fatalln("texture path is not valid")
	}

	// Load the texture
	textureID, err := texture.Create(textureSrc)
	if err != nil {
		log.Fatalln(err)
	}

	r.Program.SetUniform1i("tex", int32(0))
	r.Program.BindFragment("outputColor")

	r.Program.SetupVertexAttributeArrayAndPointer("vert", 3, 5*4, 0)
	r.Program.SetupVertexAttributeArrayAndPointer("vertTexCoord", 2, 5*4, 3*4)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, textureID)
}

// BindTexture .
func (r *Renderer) UnbindTexture() {
	gl.BindTexture(gl.TEXTURE_2D, 1)
}

// Render method for Renderer ...
func (r *Renderer) Render() {
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(r.RawRenderer)*4))
}

// Update method for Renderer ...
func (r *Renderer) Update(time, elapsed float64) {

	r.Program.SetFloat32("time", float32(time))

	camera.Update(time, elapsed)

	angle += elapsed * 0.5

	// Head Rotation Direaction
	model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, -1, 0})

	camera.LookAt(
		r.Container.Position,
		mgl32.Vec3{0, 0, 0},
		mgl32.Vec3{0, 1, 0},
	)

	r.Program.SetCamera(camera.GetPosition())
	r.Program.SetUniformMat4fv("model", model, 0, true)
}
