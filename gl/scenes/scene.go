package scenes

import (
	"matwa/caobaEngine/cameras"
	"matwa/caobaEngine/models"
	"matwa/caobaEngine/shaders"

	"github.com/go-gl/mathgl/mgl32"
)

type Scene struct {
	ModelPaths     []string
	VertexPaths    []string
	FragmentPaths  []string
	LightPositions []mgl32.Vec3
	Camera         []*cameras.Camera
	Models         []*models.Model
	Shaders        []*shaders.Shader
	renderingLogic func(s *Scene)
	preRender      func(s *Scene)
}

func NewScene(ModelPaths []string, VertexPaths []string, FragmentPaths []string, LightPos []mgl32.Vec3, Camera []*cameras.Camera) *Scene {
	var S Scene

	S.ModelPaths = ModelPaths
	S.VertexPaths = VertexPaths
	S.FragmentPaths = FragmentPaths
	S.LightPositions = LightPos
	S.Camera = Camera

	return &S
}

func (S *Scene) SetRenderingLogic(renderingLogic func(s *Scene)) {
	S.renderingLogic = renderingLogic
}

func (S *Scene) SetPreRender(preRender func(s *Scene)) {
	S.preRender = preRender
}

func (S *Scene) PreRender() {
	S.preRender(S)
}

func (S *Scene) Render() {
	S.renderingLogic(S)
}

func (S *Scene) SetupScene(noFlips bool) {

	S.setupModels(noFlips)
	S.setupShaders()

}

func (S *Scene) setupModels(noFlips bool) {

	for _, path := range S.ModelPaths {

		S.Models = append(S.Models, models.NewModel(path, false, noFlips))
	}

}

func (S *Scene) setupShaders() {
	if len(S.VertexPaths) != len(S.FragmentPaths) {
		// Handle error: slices must have the same length
		return
	}

	for i := 0; i < len(S.VertexPaths); i++ {
		vertexShaderPath := S.VertexPaths[i]
		fragmentShaderPath := S.FragmentPaths[i]

		S.Shaders = append(S.Shaders, shaders.NewShader(vertexShaderPath, fragmentShaderPath))

	}
}
