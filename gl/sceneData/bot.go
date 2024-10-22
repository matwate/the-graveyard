package scenedata

import (
	"matwa/caobaEngine/cameras"
	"matwa/caobaEngine/scenes"

	"github.com/go-gl/mathgl/mgl32"
)

func BotSceneRenderingLogic(s *scenes.Scene) {
	shader := s.Shaders[0]
	shader.Use()

	shader.SetVec4("lightColor", 1.0, 1.0, 1.0, 1.0)
	shader.SetVec3("viewPos", DefaultCamera.Position)

	projection := mgl32.Perspective(mgl32.DegToRad(DefaultCamera.Zoom), float32(w)/float32(h), 0.1, 100.0)
	view := DefaultCamera.GetViewMatrix()
	shader.SetMat4("projection", projection)
	shader.SetMat4("view", view)

	var modelMat mgl32.Mat4
	modelMat = mgl32.Translate3D(0.0, 0.0, 0)
	modelMat = mgl32.Scale3D(1, 1, 1).Mul4(modelMat)
	shader.SetMat4("model", modelMat)

	model := s.Models[0]
	model.Draw(shader)
}

func GetBotScene() scenes.Scene {

	scene := scenes.NewScene(
		[]string{"./resources/bot/d2f0cff60afc40f5afe79156ec7db657.obj"},
		[]string{"vertex.vs"},
		[]string{"fragment.fs"},
		[]mgl32.Vec3{{0.0, 0.0, 0.0}},
		[]*cameras.Camera{DefaultCamera},
	)

	scene.SetRenderingLogic(BotSceneRenderingLogic)
	scene.SetPreRender(func(s *scenes.Scene) {})

	return *scene
}
