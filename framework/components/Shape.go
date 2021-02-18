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
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// Shape Structure
type Shape struct {
	RawShape       []float32
	Vao            vertexarray.VertexArray
	Vbo            vertexbuffer.VertexBuffer
	Program        *shader.Program
	Container      *entity.Entity
}

// NewShape constructor
func NewShape(container *entity.Entity) *Shape {
	return &Shape{
		Container:      container,
	}
}

var angle = 0.1
var xPos, yPos float64
var model = mgl32.Ident4()

var camera = Camera{}

// Start func
func (s *Shape) Start() {

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	// Create And bind
	s.Vbo.New(s.RawShape, len(s.RawShape)*4)

	// Create and bind
	s.Vao.New()

	gl.UseProgram(s.Program.Data)
	s.Program.SetModel(model)

	s.BindTexture()
}

// LoadAndBindTexture .
func (s *Shape) LoadAndBindTexture(path string) {
	textureSrc, err := gas.Abs(path)
	if err != nil {
		log.Fatalln("texture path is not valid")
	}

	// Load the texture
	textureID, err := texture.Create(textureSrc, s.Container.Index)
	if err != nil {
		log.Fatalln(err)
	}

	s.Program.SetUniform1i("tex", int32(s.Container.Index))
	s.Program.BindFragment("outputColor")

	s.Program.SetupVertexAttributeArrayAndPointer("vert", 3, 5*4, 0)
	s.Program.SetupVertexAttributeArrayAndPointer("vertTexCoord", 2, 5 * 4, 3 * 4)

	gl.ActiveTexture(gl.TEXTURE0 + s.Container.Index)
	gl.BindTexture(gl.TEXTURE_2D, textureID)
}

// BindTexture .
func (s *Shape) BindTexture() {

	s.Program.SetUniform1i("tex", int32(s.Container.Index))
	s.Program.BindFragment("outputColor")

	s.Program.SetupVertexAttributeArrayAndPointer("vert", 3, 5 * 4, 0)
	s.Program.SetupVertexAttributeArrayAndPointer("vertTexCoord", 2, 5 * 4, 3 * 4)

	gl.ActiveTexture(gl.TEXTURE0 + s.Container.Index)
	gl.BindTexture(gl.TEXTURE_2D, 1)
}

// Render method for Sprite ...
func (s *Shape) Render() {
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(s.RawShape)*3))
}

// Update method for Sprite ...
func (s *Shape) Update(window *glfw.Window, time, prevTime, elapsed float64) {

	s.Program.SetFloat32("time", float32(time))

	camera.Update(window, time, prevTime, elapsed)

	angle += elapsed * 0.5

	// Head Rotation Direaction
	model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, -1, 0})

	camera.LookAt(
		mgl32.Vec3{s.Container.Position.X(), s.Container.Position.Y(), s.Container.Position.Z()},
		mgl32.Vec3{0, 0, 0},
		mgl32.Vec3{0, 1, 0},
	)

	s.Program.SetCamera(camera.GetPosition())
	s.Program.SetUniformMat4fv("model", model, 0, true)
}

// Ready method for Sprite ...
func (s *Shape) Ready() {
	s.Start()
}
