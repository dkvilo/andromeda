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
			OnLoadContext: func(andromeda *core.Andromeda) {

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

				shaderProgram := shader.Program{}
				program, err := shaderProgram.CreateProgram(string(vertexShaderData)+"\x00", string(fragmentShaderData)+"\x00")
				if err != nil {
					log.Fatalln("error while compile", err)
				}

				cubeRenderer := components.NewRenderer(&Cube, object.Load(objectSrc), &program)
				cubeRenderer.LoadAndBindTexture("github.com/dkvilo/andromeda/resources/assets/textures/art.png")

				Cube.AddComponent("Renderer", cubeRenderer)

				/**
				 * Update Projection Matrix
				 */
				cubeRenderer.Program.SetProjectionMatrix(
					mgl32.Perspective(mgl32.DegToRad(45.0),
						andromeda.Width/andromeda.Height, 0.001, 1000.0),
				)

				andromeda.EntityManager.RegisterEntity("Cube", Cube)
			},

			// Issue Entity Update
			OnUpdateContext: func(andromeda *core.Andromeda) {
				/**
				 * Update and render all Entity in the manager
				 * `Renderer` component
				 */
				for _, element := range andromeda.EntityManager.Entities {

					element.Components["Renderer"].Update(
						andromeda.GetMeta().GetTime(),
						andromeda.GetMeta().GetElapsed(),
					)
					element.Components["Renderer"].Render()
				}

			},
		},
	}
)
