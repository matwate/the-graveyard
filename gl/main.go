package main

import (
	"log"
	"matwa/graphics-engine/camera"
	"matwa/graphics-engine/models"

	"matwa/graphics-engine/shader"
	"runtime"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	w = 800 * 1920 / 1080
	h = 600 * 1920 / 1080
)

var (
	lightSourcePos = mgl32.Vec3{1.2, 1, 2.0}

	vertices = []float32{
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,
		0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 1.0,
		0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,
		0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,

		-0.5, 0.5, 0.5, -1.0, 0.0, 0.0, 1.0, 0.0,
		-0.5, 0.5, -0.5, -1.0, 0.0, 0.0, 1.0, 1.0,
		-0.5, -0.5, -0.5, -1.0, 0.0, 0.0, 0.0, 1.0,
		-0.5, -0.5, -0.5, -1.0, 0.0, 0.0, 0.0, 1.0,
		-0.5, -0.5, 0.5, -1.0, 0.0, 0.0, 0.0, 0.0,
		-0.5, 0.5, 0.5, -1.0, 0.0, 0.0, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 0.0, 0.0, 1.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0, 0.0, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 1.0, 1.0,
		0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 1.0, 0.0,
		0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 1.0,
		0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 1.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 1.0,
	}
	cubePositions = []mgl32.Vec3{
		{0.0, 0.0, 0.0},
		{2.0, 5.0, -15.0},
		{-1.5, -2.2, -2.5},
		{-3.8, -2.0, -12.3},
		{2.4, -0.4, -3.5},
		{-1.7, 3.0, -7.5},
		{1.3, -2.0, -2.5},
		{1.5, 2.0, -2.5},
		{1.5, 0.2, -1.5},
		{-1.3, 1.0, -1.5},
	}

	indices = []uint32{
		0, 1, 3,
		1, 2, 3,
	}
	identity = mgl32.Ident4()

	deltaTime float32
	lastFrame float64

	cam = camera.NewCamera(
		mgl32.Vec3{0, 0, 3},
		mgl32.Vec3{0, 1, 0},
		-90,
		0,
	)

	lightColor   = mgl32.Vec3{1, 1, 1}
	diffuseColor = lightColor.Mul(0.5)
	ambientColor = diffuseColor.Mul(0.2)
)

func main() {

	runtime.LockOSThread()

	window := initGlfw()

	myShader := shader.NewShader("vertex.glsl", "fragment.glsl")

	myModel := models.NewModel("./TestModel.gltf")

	for !window.ShouldClose() {

		currentFrame := glfw.GetTime()
		deltaTime = float32(currentFrame - lastFrame)
		lastFrame = currentFrame

		processInput(window)
		gl.ClearColor(0.2, 0.2, 0.2, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Create an identity matrix

		myShader.Use()
		proj := mgl32.Perspective(mgl32.DegToRad(cam.Zoom), w/h, 0.1, 100)
		view := cam.GetViewMatrix()

		myShader.SetMat4("projection", proj)
		myShader.SetMat4("view", view)

		model := mgl32.Translate3D(0, 0, 0).Mul4(identity)
		myShader.SetMat4("model", model)

		myModel.Draw(myShader)

		window.SwapBuffers()
		glfw.PollEvents()

		log.Default().Println("FPS: ", 1/deltaTime)

	}

	defer glfw.Terminate()

}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(w, h, "Hello World", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	window.SetCursorPosCallback(mouseCallback)
	window.SetScrollCallback(scrollCallback)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.Enable(gl.DEPTH_TEST)
	return window
}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
	if window.GetKey(glfw.KeyW) == glfw.Press {
		cam.ProcessKeyboard(camera.Forward, deltaTime)
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		cam.ProcessKeyboard(camera.Backward, deltaTime)
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		cam.ProcessKeyboard(camera.Left, deltaTime)
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		cam.ProcessKeyboard(camera.Right, deltaTime)
	}

}

func mouseCallback(window *glfw.Window, xpos, ypos float64) {
	if cam.FirstMouse {
		cam.LastX = float32(xpos)
		cam.LastY = float32(ypos)
		cam.FirstMouse = false
	}

	xoffset := float32(xpos) - cam.LastX
	yoffset := cam.LastY - float32(ypos)

	cam.LastX = float32(xpos)
	cam.LastY = float32(ypos)

	cam.ProcessMouseMovement(xoffset, yoffset, true)
}

func scrollCallback(window *glfw.Window, xo, yoff float64) {
	cam.ProcessMouseScroll(float32(yoff))

}
