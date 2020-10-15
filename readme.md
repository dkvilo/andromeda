# Andromeda Renderer

Core Renderer of andromeda Game Engine written in go, using Modern OpenGL API

## Run

```sh
  # install module deps
  $ go mod download

  # run example scenes from sandbox layer
  $ make build && make run
```

## Framework usage
```go

package sandbox

import (
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"runtime"

	"github.com/andrebq/gas"
	"github.com/dkvilo/andromeda/engine/components"
	. "github.com/dkvilo/andromeda/engine/core"
	"github.com/dkvilo/andromeda/engine/entity"
	"github.com/dkvilo/andromeda/engine/object"
	"github.com/dkvilo/andromeda/engine/shader"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	runtime.LockOSThread();
}

var (
	// MyApp Example App
	MyApp Example = Example{
		Andromeda {
			Width:  1080,
			Height: 720,
			Title:  "Andromeda 3d Sandbox",
			
			// Issue OnLoad 
			SetupOnLoadContext: func(andromeda *Andromeda) {

				vertexShaderSrc, err := gas.Abs("github.com/dkvilo/andromeda/engine/assets/shaders/reflection/shader.vert")
				if err != nil {
					log.Fatalln("vertex shader path found")
				}

				fragmentShaderSrc, err := gas.Abs("github.com/dkvilo/andromeda/engine/assets/shaders/reflection/shader.frag")
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
					Index: 0,
					Position: mgl32.Vec3{3, 3, 3},
				}

				cubeShape := components.NewShape(&Cube)
				cubeShape.RawShape = object.Load(objectSrc)

				shaderProgram := shader.Program{}
				program, err := shaderProgram.CreateProgram(string(vertexShaderData) + string("\x00"), string(fragmentShaderData) + "\x00")
				if err != nil {
					log.Fatalln("error while compile", err)
				}

				cubeShape.Program = &program;

				Cube.AddComponent("Shape", cubeShape)
				Cube.OnStart()
				
				cubeShape.Program.SetProjectionMatrix(
					mgl32.Perspective(mgl32.DegToRad(52.0),
					andromeda.Width / andromeda.Height, 0.1, 100.0),
				)
				
				cubeShape.LoadAndBindTexture("github.com/dkvilo/andromeda/resources/assets/textures/art.png");
				andromeda.EntityManager.RegisterEntity("Cube", Cube)
			},

			// Issue Entity Update
			SetupOnUpdateContext: func(andromeda *Andromeda) {				
				andromeda.EntityManager.GetEntities()["Cube"].Components["Shape"].OnUpdate(
					andromeda.GetWindow(),
					andromeda.GetMeta().GetTime(),
					andromeda.GetMeta().GetPreviousTime(),
					andromeda.GetMeta().GetElapsed(),
				);
			},

			// Issue Draw Call
			SetupOnRenderContext: func(andromeda *Andromeda) {
				andromeda.EntityManager.GetEntities()["Cube"].Components["Shape"].OnRender();
			},
		},
	}
)


```

