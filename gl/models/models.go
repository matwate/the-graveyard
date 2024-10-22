package models

import (
	"fmt"
	"log"
	"matwa/caobaEngine/shaders"
	"strings"

	"github.com/anthonynsimon/bild/transform"
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/raedatoui/assimp"
	"neilpa.me/go-stbi"
)

type Model struct {
	loaded_tex      []Texture
	meshses         []Mesh
	dir             string
	gammaCorrection bool
	textureFlips    bool
}

func NewModel(dir string, gammaCorrection bool, noFlips bool) *Model {
	var m Model
	m.dir = dir
	m.gammaCorrection = gammaCorrection
	m.textureFlips = noFlips
	m.LoadModel(m.dir)
	return &m
}

func (m *Model) Draw(s *shaders.Shader) {
	for i := 0; i < len(m.meshses); i++ {
		m.meshses[i].Draw(s)
	}
}

func (m *Model) LoadModel(path string) {

	aiScene := assimp.ImportFile(path, uint(assimp.Process_Triangulate)|uint(assimp.Process_FlipUVs))
	if aiScene == nil {
		panic("Error loading model")
	}

	m.dir = path[:strings.LastIndex(path, "/")]

	m.ProcessNode(aiScene.RootNode(), aiScene)
}

func (m *Model) ProcessNode(node *assimp.Node, scene *assimp.Scene) {

	for i := 0; i < len(node.Meshes()); i++ {
		aiMesh := scene.Meshes()[node.Meshes()[i]]
		m.meshses = append(m.meshses, m.ProcessMesh(aiMesh, scene))
	}

	for i := 0; i < len(node.Children()); i++ {
		m.ProcessNode(node.Children()[i], scene)
	}

}

func (m *Model) ProcessMesh(aiMesh *assimp.Mesh, scene *assimp.Scene) Mesh {
	var v []Vertex
	var idxs []uint32
	var t []Texture

	// Processing vertices
	log.Println("Processing vertices")
	for i := 0; i < len(aiMesh.Vertices()); i++ {
		var vertex Vertex
		vertex.Pos = mgl32.Vec3{aiMesh.Vertices()[i].X(), aiMesh.Vertices()[i].Y(), aiMesh.Vertices()[i].Z()}
		log.Printf("Processing vertex %d", i)
		if len(aiMesh.Normals()) > 0 {
			vertex.Normal = mgl32.Vec3{aiMesh.Normals()[i].X(), aiMesh.Normals()[i].Y(), aiMesh.Normals()[i].Z()}
		} else {
			log.Printf("No normals found for vertex %d", i)
		}
		log.Printf("Processing texture coordinates for vertex %d", i)

		if len(aiMesh.TextureCoords(0)) > 0 {
			vertex.TexCoords = mgl32.Vec2{aiMesh.TextureCoords(0)[i].X(), aiMesh.TextureCoords(0)[i].Y()}

			log.Printf("Processing tangents and bitangents for vertex %d", i)

			if len(aiMesh.Tangents()) > 0 {
				vertex.Tangent = mgl32.Vec3{aiMesh.Tangents()[i].X(), aiMesh.Tangents()[i].Y(), aiMesh.Tangents()[i].Z()}
			} else {
				log.Printf("No tangents found for vertex %d", i)
			}

			log.Printf("Processing bitangents for vertex %d", i)

			if len(aiMesh.Bitangents()) > 0 {
				vertex.Bitangent = mgl32.Vec3{aiMesh.Bitangents()[i].X(), aiMesh.Bitangents()[i].Y(), aiMesh.Bitangents()[i].Z()}
			} else {
				log.Printf("No bitangents found for vertex %d", i)
			}
			log.Printf("Vertex %d: %v", i, vertex)
		} else {
			log.Printf("No texture coordinates found for vertex %d", i)
			vertex.TexCoords = mgl32.Vec2{0.0, 0.0}
		}

		log.Printf("Vertex %d: %v", i, vertex)
		v = append(v, vertex)
	}

	// Processing faces
	for i := 0; i < len(aiMesh.Faces()); i++ {
		face := aiMesh.Faces()[i]
		a := face.CopyIndices()
		log.Printf("Processing face %d with %d indices", i, face.NumIndices())

		for j := 0; j < int(face.NumIndices()); j++ {
			if j >= len(a) {
				log.Printf("Index out of range for face %d: index %d, length %d", i, j, len(a))
			} else {
				idxs = append(idxs, a[j])
			}
		}
	}

	// Accessing materials
	if aiMesh.MaterialIndex() < len(scene.Materials()) {
		material := scene.Materials()[aiMesh.MaterialIndex()]
		log.Printf("Processing material index %d", aiMesh.MaterialIndex())

		diffuseMaps := m.LoadMaterialTextures(material, assimp.TextureType(1), "texture_diffuse")
		t = append(t, diffuseMaps...)

		specularMaps := m.LoadMaterialTextures(material, assimp.TextureType(2), "texture_specular")
		t = append(t, specularMaps...)

		normalMaps := m.LoadMaterialTextures(material, assimp.TextureType(6), "texture_normal")
		t = append(t, normalMaps...)

		heightMaps := m.LoadMaterialTextures(material, assimp.TextureType(5), "texture_height")
		t = append(t, heightMaps...)
	} else {
		log.Printf("Material index %d out of range", aiMesh.MaterialIndex())
	}

	return *NewMesh(v, idxs, t)
}

func (m *Model) LoadMaterialTextures(material *assimp.Material, texType assimp.TextureType, typeName string) []Texture {

	var textures []Texture

	for i := 0; i < int(material.GetMaterialTextureCount(texType)); i++ {

		str, _, _, _, _, _, _, _ := material.GetMaterialTexture(texType, int(i))
		var skip bool = false

		for j := 0; j < len(m.loaded_tex); j++ {

			if m.loaded_tex[j].Path == str {
				textures = append(textures, m.loaded_tex[j])
				skip = true
				break
			}
		}

		if !skip {
			var texture Texture
			if m.textureFlips {
				texture.ID = TextureFromFile(str, m.dir, m.gammaCorrection, false)
			} else {
				texture.ID = TextureFromFile(str, m.dir, m.gammaCorrection, true)
			}

			texture.Type = typeName
			texture.Path = str
			textures = append(textures, texture)
			m.loaded_tex = append(m.loaded_tex, texture)
		}

	}
	return textures
}

func TextureFromFile(path, directory string, gamma bool, flipVert bool) uint32 {

	filename := directory + "/" + path

	var texture uint32
	gl.GenTextures(1, &texture)

	data, err := stbi.Load(filename)
	if err != nil {
		panic(err)
	}
	if data == nil || data.Pix == nil {
		panic("Failed to load texture")
	}

	if flipVert {
		data = transform.FlipV(data)
		if data == nil {
			panic("Failed to load texture")
		}
	}

	gl.BindTexture(gl.TEXTURE_2D, texture)
	if err := gl.GetError(); err != gl.NO_ERROR {
		panic(fmt.Sprintf("OpenGL error after BindTexture: %v", err))
	}
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(data.Rect.Dx()), int32(data.Rect.Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(data.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	if err := gl.GetError(); err != gl.NO_ERROR {
		panic(fmt.Sprintf("OpenGL error after GenerateMipmap: %v", err))
	}

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	return texture

}
