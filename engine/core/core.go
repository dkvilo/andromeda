package core

import (
	"fmt"
	"github.com/andrebq/gas"
	"github.com/dkvilo/andromeda/engine/object"
	"github.com/dkvilo/andromeda/engine/shader"
	"github.com/dkvilo/andromeda/engine/texture"
	"github.com/dkvilo/andromeda/engine/vertexarray"
	"github.com/dkvilo/andromeda/engine/vertexbuffer"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"io/ioutil"
	"log"
	"math"
	"reflect"
)

type Andromeda struct {
	Width, Height float32
	Title string
	window *glfw.Window
	err error
	running bool
	StaticFps uint32
	tick uint32
	eventType uint32
	time float32
}

func (andromeda *Andromeda) init() *Andromeda{

	if andromeda.err = glfw.Init(); andromeda.err != nil {
		log.Fatalf("GLFW Init: %s", andromeda.err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	if andromeda.window, andromeda.err = glfw.CreateWindow(int(andromeda.Width), int(andromeda.Height), andromeda.Title, nil, nil); andromeda.err !=  nil {
		log.Fatalf("GLFW CreateWindiw: %s", andromeda.err)
	}

	andromeda.window.MakeContextCurrent()

	if andromeda.err = gl.Init(); andromeda.err != nil {
		log.Fatalf("GL Initi: %s", andromeda.err)
	}

	fmt.Println("GL Version:", gl.GoStr(gl.GetString(gl.VERSION)))

	andromeda.running = true

	return andromeda
}

func (andromeda *Andromeda) masterRenderer() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.7, 0.6, 0.8, 1)
}

func (andromeda *Andromeda) Run() *Andromeda {
	andromeda.init()
	andromeda.loadScene()

	defer func() {
		andromeda.window.Destroy()
		glfw.Terminate()
	} ()

	return andromeda
}

func sizeOf(t interface{}) uintptr {
	return reflect.TypeOf(t).Size()
}

var xPos, yPos float64

func (andromeda *Andromeda) loadScene() {

	objectSrc, err := gas.Abs("github.com/dkvilo/andromeda/resources/assets/objects/cube.andromeda")
	if err != nil {
		panic(err)
	}

	objectVertices := object.Load(objectSrc)

	vertexShaderSrc, err := gas.Abs("github.com/dkvilo/andromeda/engine/assets/shaders/shader.vsh")
	if err != nil {
		log.Fatalln("vertex shader path found")
	}

	fragmentShaderSrc, err := gas.Abs("github.com/dkvilo/andromeda/engine/assets/shaders/shader.fsh")
	if err != nil {
		log.Fatalln("fragment shader path found")
	}

	vertexShaderData, err := ioutil.ReadFile(vertexShaderSrc)
	if err != nil {
		log.Fatalln("vertex shader not found")
	}

	fragmentShaderData, err := ioutil.ReadFile(fragmentShaderSrc)
	if err != nil {
		log.Fatalln("fragment shader not found")
	}

	// Configure the vertex and fragment shaders
	program, err := shader.Program(string(vertexShaderData) + string("\x00"), string(fragmentShaderData) + "\x00")
	if err != nil {
		log.Fatalln("error while compile", err)
	}

	gl.UseProgram(program)

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), andromeda.Width / andromeda.Height, 0.1, 100.0)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	cubeVertexArray := vertexarray.VertexArray{}
	cubeVertexArray.New()

	cubeVertexBuffer := vertexbuffer.VertexBuffer{}
	cubeVertexBuffer.New(objectVertices, len(objectVertices) * 4)

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5 * 4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5 * 4, gl.PtrOffset(3*4))

	textureSrc, err := gas.Abs("github.com/dkvilo/andromeda/resources/assets/textures/art.png")
	if err != nil {
		log.Fatalln("texture path is not valid")
	}

	// Load the texture
	textureData, err := texture.Create(textureSrc)
	if err != nil {
		log.Fatalln(err)
	}

	angle := 0.0
	previousTime := glfw.GetTime()
	
	for !andromeda.window.ShouldClose() {

		andromeda.window.SetInputMode(glfw.CursorMode, glfw.CursorHidden);

		xPos, yPos = andromeda.window.GetCursorPos()

		andromeda.masterRenderer()

		// Update
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time

		// Configure global settings
		gl.Enable(gl.DEPTH_TEST)
		gl.DepthFunc(gl.LESS)
		gl.ClearColor(float32(math.Cos(time)), float32(math.Sin(time)), float32(math.Cos(time)), 0.1)

		texVertTime := float32(gl.GetUniformLocation(program, gl.Str("time\x00")))
		gl.Uniform1f(int32(texVertTime), float32(time))

		camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})

		angle += elapsed * 0.5

		if andromeda.window.GetKey(glfw.KeyR) == 1 {
			model = mgl32.HomogRotate3D(float32(xPos / 100 ), mgl32.Vec3{0, -1, 0})
			camera = mgl32.LookAtV(mgl32.Vec3{float32(xPos / 100), float32(yPos / 100), 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
		} else {
			model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, -1, 0})
		}

		cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
		gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

		if andromeda.window.GetKey(glfw.KeyL) == 1 {
			gl.Enable(gl.BLEND)
			gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		} else {
			gl.Disable(gl.BLEND)
		}

		// Render
		gl.UseProgram(program)
		gl.UniformMatrix4fv(modelUniform, 1, true, &model[0])

		cubeVertexArray.Bind()

		gl.ActiveTexture(gl.TEXTURE0)

		gl.BindTexture(gl.TEXTURE_2D, textureData)
		gl.DrawArrays(gl.TRIANGLES, 0, 6 * 2 * 3)

		if andromeda.window.GetKey(glfw.KeyQ) == 1 {
			break
		}

		andromeda.window.SwapBuffers()
		glfw.PollEvents()
	}

}




