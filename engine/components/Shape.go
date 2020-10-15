package components

import (
	"log"

	"github.com/andrebq/gas"
	"github.com/dkvilo/andromeda/engine/entity"
	"github.com/dkvilo/andromeda/engine/shader"
	"github.com/dkvilo/andromeda/engine/texture"
	"github.com/dkvilo/andromeda/engine/vertexarray"
	"github.com/dkvilo/andromeda/engine/vertexbuffer"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// Shape Structure
type Shape struct {
	RawShape []float32
	Vao vertexarray.VertexArray
	Vbo vertexbuffer.VertexBuffer
	Program *shader.Program
	Container *entity.Entity
	PointerOffset int
	PointerOffset2 int
}

// NewShape constructor
func NewShape(container *entity.Entity) *Shape {
	return &Shape{
		Container: container,
		PointerOffset: 0,
		PointerOffset2: 12,
	}
}

var angle = 0.1
var xPos, yPos float64
var model = mgl32.Ident4()

var camera = Camera {}

var cameraUp = Transform {
	Label: "Camera Up Vector",
	Rotation: mgl32.Vec3{0, 1, 0},
	Position: mgl32.Vec3{0, 0, 0},
}

var cameraEye = Transform {
	Label: "Camera Eye",
	Position: mgl32.Vec3{3, 3, 3},
}

var cameraCenter = Transform {
	Label: "Camera Center",
	Position: mgl32.Vec3{0, 0, 0},
}

// Start func
func (s *Shape) Start() {

	s.PointerOffset = 0
	s.PointerOffset2 = 12

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	
	s.Vbo.New(s.RawShape, len(s.RawShape) * 4) 
	s.Vbo.Bind()

	s.Vao.New()
	s.Vao.Bind()

	gl.UseProgram(s.Program.Data)
	s.Program.SetModel(model)
	camera.LookAt(cameraEye.Position, cameraCenter.Position, cameraUp.Position)
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

	s.Program.SetupVertexAttributeArrayAndPointer("vert", 3, 5 * 4, s.PointerOffset)
	s.Program.SetupVertexAttributeArrayAndPointer("vertTexCoord", 2, 5 * 4, s.PointerOffset2)

	gl.ActiveTexture(gl.TEXTURE0 + s.Container.Index)
	gl.BindTexture(gl.TEXTURE_2D, textureID)
}

// BindTexture . 
func (s *Shape) BindTexture() {

	s.Program.SetUniform1i("tex", int32(s.Container.Index))
	s.Program.BindFragment("outputColor")

	s.Program.SetupVertexAttributeArrayAndPointer("vert", 3, 5 * 4, s.PointerOffset)
	s.Program.SetupVertexAttributeArrayAndPointer("vertTexCoord", 2, 5 * 4, s.PointerOffset2)

	gl.ActiveTexture(gl.TEXTURE0 + s.Container.Index)
	gl.BindTexture(gl.TEXTURE_2D, 1)
}

// OnRender method for Sprite ...
func (s *Shape) OnRender() {
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(s.RawShape) * 3))
}

// OnUpdate method for Sprite ...
func (s *Shape) OnUpdate(window *glfw.Window, time, prevTime, elapsed float64) {
	
	s.Program.SetFloat32("time", float32(time))

	camera.OnUpdate(window, time, prevTime, elapsed);
	
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

// OnStart method for Sprite ...
func (s *Shape) OnStart() {
	s.Start()
}
