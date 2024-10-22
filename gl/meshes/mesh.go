package mesh

import (
	"matwa/graphics-engine/buffers"
	camera "matwa/graphics-engine/cameras"
	shaders "matwa/graphics-engine/shader"
	"matwa/graphics-engine/vertex"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Mesh struct {
	vertices []vertex.Vertex
	indices  []uint32
	textures []Texture
	vao      buffers.VAO
}

func NewMesh(vertices []vertex.Vertex, indices []uint32, textures []Texture) *Mesh {

	m := &Mesh{}

	m.vertices = vertices
	m.indices = indices
	m.textures = textures

	vao := buffers.NewVAO()
	m.vao = *vao

	veretexData := []float32{}

	for _, v := range m.vertices {
		veretexData = append(veretexData, v.Data()...)

	}

	vbo := buffers.NewVBO(veretexData, len(m.vertices)*4*8)

	ebo := buffers.NewEBO(m.indices)

	m.vao.LinkAttribute(*vbo, 0, 3, 8*4, gl.FLOAT, 0)
	m.vao.LinkAttribute(*vbo, 1, 3, 8*4, gl.FLOAT, 3*4)
	m.vao.LinkAttribute(*vbo, 2, 3, 8*4, gl.FLOAT, 6*4)
	m.vao.LinkAttribute(*vbo, 3, 2, 8*4, gl.FLOAT, 9*4)

	m.vao.Unbind()
	vbo.Unbind()
	ebo.Unbind()

	return m
}

func (m *Mesh) Draw(

	s *shaders.Shader,
	c *camera.Camera,
	mat mgl32.Mat4,
	trans mgl32.Vec3,
	rot mgl32.Quat,
	scale mgl32.Vec3,

) {
	s.Activate()
	m.vao.Bind()

	numDiffuseTextures := 0
	numSpecularTextures := 0

	for i, texture := range m.textures {
		
		num := string(0)

}
