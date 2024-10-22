package main

import (
	"log"
	"matwa/caobaEngine/sceneHandler"

	imgui "github.com/AllenDang/cimgui-go"
)

var backend imgui.Backend[imgui.GLFWWindowFlags]
var loadScene bool
var current_item int32
var items = []string{"bot", "backpack", "nada"}

func loop() {

	imgui.SetNextWindowSizeV(imgui.NewVec2(800, 600), imgui.CondOnce)
	imgui.BeginV("Hello, world!", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoCollapse|imgui.WindowFlagsNoMove)
	imgui.SetWindowFontScale(1.5)
	imgui.Text("This is some useful text.")

	if imgui.ComboStrarrV("Items", &current_item, items, int32(len(items)), 10) {
		log.Println("Selected item: ", items[current_item])
	}

	if imgui.Button("Load Scene") {
		loadScene = true
		backend.SetShouldClose(true)
	}

	imgui.End()

}

func afterCreateContext() {

}

func beforeDestroyContext() {

}
func main() {

	backend, _ = imgui.CreateBackend(imgui.NewGLFWBackend())
	backend.SetAfterCreateContextHook(afterCreateContext)
	backend.SetBeforeDestroyContextHook(beforeDestroyContext)
	backend.CreateWindow("hello", 800, 600)

	backend.Run(loop)

	if loadScene {
		loadScene = false
		defer sceneHandler.RunScene(items[current_item])

	} else {
		log.Println("No scene selected")
	}

}
