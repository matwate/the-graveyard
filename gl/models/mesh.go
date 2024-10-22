package models

import (
	"fmt"
	"matwa/caobaEngine/shaders"
	"unsafe"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Vertex struct {
	Pos       mgl32.Vec3
	Normal    mgl32.Vec3
	TexCoords mgl32.Vec2
	Tangent   mgl32.Vec3
	Bitangent mgl32.Vec3
}

type Texture struct {
	ID   uint32
	Type string
	Path string
}

type Mesh struct {
	Vertices []Vertex
	Indices  []uint32
	Textures []Texture
	VAO      uint32
	VBO      uint32
	EBO      uint32
}

func NewMesh(vertices []Vertex, indices []uint32, textures []Texture) *Mesh {

	var m Mesh
	m.Vertices = vertices
	m.Indices = indices
	m.Textures = textures

	m.Setup()
	return &m
}

func (m *Mesh) Draw(s *shaders.Shader) {

	diffuseNr := uint(1)
	specularNr := uint(1)
	normalNr := uint(1)
	heightNr := uint(1)

	for i := 0; i < len(m.Textures); i++ {

		gl.ActiveTexture(gl.TEXTURE0 + uint32(i))

		var number string
		var name string = m.Textures[i].Type

		switch name {
		case "texture_diffuse":
			number = fmt.Sprint(diffuseNr)
			diffuseNr++
		case "texture_specular":
			number = fmt.Sprint(specularNr)
			specularNr++
		case "texture_normal":
			number = fmt.Sprint(normalNr)
			normalNr++
		case "texture_height":
			number = fmt.Sprint(heightNr)
		}

		s.SetInt(name+number, int32(i))

		gl.BindTexture(gl.TEXTURE_2D, m.Textures[i].ID)

	}

	gl.BindVertexArray(m.VAO)
	gl.DrawElements(gl.TRIANGLES, int32(len(m.Indices)), gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.BindVertexArray(0)

	gl.ActiveTexture(gl.TEXTURE0)

}

func (m *Mesh) Setup() {

	gl.GenVertexArrays(1, &m.VAO)
	gl.GenBuffers(1, &m.VBO)
	gl.GenBuffers(1, &m.EBO)

	gl.BindVertexArray(m.VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, m.VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(m.Vertices)*int(unsafe.Sizeof(m.Vertices[0])), gl.Ptr(m.Vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.Indices)*int(unsafe.Sizeof(m.Indices[0])), gl.Ptr(m.Indices), gl.STATIC_DRAW)

	// Vertex Positions
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, int32(unsafe.Sizeof(m.Vertices[0])), 0)
	// Vertex Normals
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, int32(unsafe.Sizeof(m.Vertices[0])), uintptr(unsafe.Offsetof(m.Vertices[0].Normal)))
	// Vertex Texture Coords
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, int32(unsafe.Sizeof(m.Vertices[0])), uintptr(unsafe.Offsetof(m.Vertices[0].TexCoords)))
	// Vertex Tangent
	gl.EnableVertexAttribArray(3)
	gl.VertexAttribPointerWithOffset(3, 3, gl.FLOAT, false, int32(unsafe.Sizeof(m.Vertices[0])), uintptr(unsafe.Offsetof(m.Vertices[0].Tangent)))
	// Vertex Bitangent
	gl.EnableVertexAttribArray(4)
	gl.VertexAttribPointerWithOffset(4, 3, gl.FLOAT, false, int32(unsafe.Sizeof(m.Vertices[0])), uintptr(unsafe.Offsetof(m.Vertices[0].Bitangent)))
	gl.BindVertexArray(0)

}
