package sceneHandler

import (
	"log"
	"matwa/caobaEngine/cameras"
	scenedata "matwa/caobaEngine/sceneData"
	"runtime"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	w = 800 * 1920 / 1080
	h = 600 * 1920 / 1080

	lastX = float32(w / 2)
	lastY = float32(h / 2)

	firstMouse = true
	camera     *cameras.Camera
	deltaTime  = 0.0
	lastFrame  = 0.0

	backend imgui.Backend[imgui.GLFWWindowFlags]
)

func Loop() {

}

func RunScene(sceneName string) {
	runtime.LockOSThread()

	window := initGl()

	vendor := gl.GoStr(gl.GetString(gl.VENDOR))
	renderer := gl.GoStr(gl.GetString(gl.RENDERER))

	var lastFrame float32 = 0.0

	scene := scenedata.AvailableScenes[sceneName]()
	camera = scene.Camera[0]
	scene.SetupScene(false)

	scene.PreRender()

	for !window.ShouldClose() {
		var currentFrame float32 = float32(glfw.GetTime())
		deltaTime := currentFrame - lastFrame
		lastFrame = currentFrame

		processInput(window, deltaTime)

		gl.ClearColor(0, 0.1, 0.1, 0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)

		scene.Render()

		window.SwapBuffers()
		glfw.PollEvents()

		log.Print("FPS: ", 1/deltaTime)

	}

	defer glfw.Terminate()
	defer log.Println("Vendor: ", vendor)
	defer log.Println("Renderer: ", renderer)
}

func initGl() *glfw.Window {

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(h, w, "Caoba Engine", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetFramebufferSizeCallback(framebuffer_size_callback)
	window.SetCursorPosCallback(mouse_callback)
	window.SetScrollCallback(scroll_callback)

	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	if err := gl.Init(); err != nil {

		panic(err)

	}

	gl.Enable(gl.DEPTH_TEST)

	return window

}

func framebuffer_size_callback(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func mouse_callback(window *glfw.Window, xposIn float64, yposIn float64) {

	xpos := float32(xposIn)
	ypos := float32(yposIn)

	if firstMouse {
		lastX = xpos
		lastY = ypos
		firstMouse = false
	}

	xoffset := xpos - lastX
	yoffset := lastY - ypos

	lastX = xpos
	lastY = ypos

	camera.ProcessMouseMovement(xoffset, yoffset, true)

}

func scroll_callback(window *glfw.Window, xoffset float64, yoffset float64) {

	camera.ProcessMouseScroll(float32(yoffset))
}

func processInput(window *glfw.Window, deltaTime float32) {

	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
	if window.GetKey(glfw.KeyW) == glfw.Press {

		camera.ProcessKeyboard(cameras.Forward, deltaTime)

	}
	if window.GetKey(glfw.KeyS) == glfw.Press {

		camera.ProcessKeyboard(cameras.Backward, deltaTime)

	}
	if window.GetKey(glfw.KeyA) == glfw.Press {

		camera.ProcessKeyboard(cameras.Left, deltaTime)

	}
	if window.GetKey(glfw.KeyD) == glfw.Press {

		camera.ProcessKeyboard(cameras.Right, deltaTime)

	}

}

func RunSceneNoFlips(sceneName string) {
	runtime.LockOSThread()

	window := initGl()

	vendor := gl.GoStr(gl.GetString(gl.VENDOR))
	renderer := gl.GoStr(gl.GetString(gl.RENDERER))

	var lastFrame float32 = 0.0

	scene := scenedata.AvailableScenes[sceneName]()
	camera = scene.Camera[0]
	scene.SetupScene(true)

	for !window.ShouldClose() {
		var currentFrame float32 = float32(glfw.GetTime())
		deltaTime := currentFrame - lastFrame
		lastFrame = currentFrame

		processInput(window, deltaTime)

		gl.ClearColor(0, 0.1, 0.1, 0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		scene.Render()

		window.SwapBuffers()
		glfw.PollEvents()

		log.Print("FPS: ", 1/deltaTime)

	}

	defer glfw.Terminate()
	defer log.Println("Vendor: ", vendor)
	defer log.Println("Renderer: ", renderer)
}
