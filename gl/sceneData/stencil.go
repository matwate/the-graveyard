package scenedata

import (
	"matwa/caobaEngine/cameras"
	"matwa/caobaEngine/scenes"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func NormalDrawCall(s *scenes.Scene) {
	shader := s.Shaders[0]
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

func ScaledUpDrawCall(s *scenes.Scene) {
	shader := s.Shaders[0]
	shader.SetVec4("lightColor", 1.0, 1.0, 1.0, 1.0)
	shader.SetVec3("viewPos", DefaultCamera.Position)

	projection := mgl32.Perspective(mgl32.DegToRad(DefaultCamera.Zoom), float32(w)/float32(h), 0.1, 100.0)
	view := DefaultCamera.GetViewMatrix()
	shader.SetMat4("projection", projection)
	shader.SetMat4("view", view)

	var modelMat mgl32.Mat4
	modelMat = mgl32.Translate3D(0.0, 0.0, 0)
	modelMat = mgl32.Scale3D(2, 2, 2).Mul4(modelMat)
	shader.SetMat4("model", modelMat)

	model := s.Models[0]
	model.Draw(shader)
}

func GetStencilTestingScene() {
	scene := scenes.NewScene(
		[]string{"./resources/sword/TinyTruck.obj"},
		[]string{"vertex.vs"},
		[]string{"fragment.fs", "stencil.fs"},
		[]mgl32.Vec3{{0.0, 0.0, 0.0}},
		[]*cameras.Camera{DefaultCamera},
	)

	scene.SetPreRender(func(s *scenes.Scene) {
		gl.Enable(gl.STENCIL_TEST)
	})

	scene.SetRenderingLogic(
		func(s *scenes.Scene) {
			gl.StencilOp(gl.KEEP, gl.KEEP, gl.REPLACE)
			gl.StencilFunc(gl.ALWAYS, 1, 0xFF)
			gl.StencilMask(0xFF)
			s.Shaders[0].Use()
			NormalDrawCall(s)
			gl.StencilFunc(gl.NOTEQUAL, 1, 0xFF)
			gl.StencilMask(0x00)
			gl.Disable(gl.DEPTH_TEST)
			s.Shaders[1].Use()
			ScaledUpDrawCall(s)
			gl.Enable(gl.DEPTH_TEST)
		},
	)

}
