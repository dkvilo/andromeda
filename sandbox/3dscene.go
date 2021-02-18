package sandbox

import (
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"runtime"

	"github.com/andrebq/gas"
	"github.com/dkvilo/andromeda/framework/components"
	core "github.com/dkvilo/andromeda/framework/core"
	"github.com/dkvilo/andromeda/framework/entity"
	"github.com/dkvilo/andromeda/framework/object"
	"github.com/dkvilo/andromeda/framework/shader"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	runtime.LockOSThread()
}

var (
	// Scene3d Example App
	Scene3d Example = Example{
		core.Andromeda{
			Width:  1080,
			Height: 720,
			Title:  "Andromeda 3d Sandbox",

			// Issue OnLoad
			SetupOnLoadContext: func(andromeda *core.Andromeda) {

				vertexShaderSrc, err := gas.Abs("github.com/dkvilo/andromeda/framework/assets/shaders/reflection/shader.vert")
				if err != nil {
					log.Fatalln("vertex shader path found")
				}

				fragmentShaderSrc, err := gas.Abs("github.com/dkvilo/andromeda/framework/assets/shaders/reflection/shader.frag")
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

				objectSrc, err := gas.Abs("github.com/dkvilo/andromeda/resources/assets/objects/cube.andromeda")
				if err != nil {
					panic(err)
				}

				Cube := entity.Entity{
					Index:    0,
					Position: mgl32.Vec3{3, 3, 3},
				}

				cubeShape := components.NewShape(&Cube)
				cubeShape.RawShape = object.Load(objectSrc)

				shaderProgram := shader.Program{}
				program, err := shaderProgram.CreateProgram(string(vertexShaderData)+"\x00", string(fragmentShaderData)+"\x00")
				if err != nil {
					log.Fatalln("error while compile", err)
				}

				cubeShape.Program = &program

				Cube.AddComponent("Shape", cubeShape)
				Cube.Ready()

				cubeShape.LoadAndBindTexture("github.com/dkvilo/andromeda/resources/assets/textures/water.png")

				cubeShape.Program.SetProjectionMatrix(
					mgl32.Perspective(mgl32.DegToRad(62.0),
						andromeda.Width/andromeda.Height, 0.001, 1000.0),
				)
				
				andromeda.EntityManager.RegisterEntity("Cube", Cube)
			},

			// Issue Entity Update
			SetupOnUpdateContext: func(andromeda *core.Andromeda) {
				andromeda.EntityManager.GetEntities()["Cube"].Components["Shape"].Update(
					andromeda.GetWindow(),
					andromeda.GetMeta().GetTime(),
					andromeda.GetMeta().GetPreviousTime(),
					andromeda.GetMeta().GetElapsed(),
				)
			},

			// Issue Draw Call
			SetupOnRenderContext: func(andromeda *core.Andromeda) {
				andromeda.EntityManager.GetEntities()["Cube"].Components["Shape"].Render()
			},
		},
	}
)
