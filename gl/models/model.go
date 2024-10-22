package models

import (
	"encoding/binary"
	"math"
	"matwa/graphics-engine/shader"
	"os"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/qmuntal/gltf"
)

type Model struct {
	data              *gltf.Document
	Meshes            []Mesh
	translationMeshes []mgl32.Vec3
	rotationMeshes    []mgl32.Quat
	scaleMeshes       []mgl32.Vec3
	matricesMeshes    []mgl32.Mat4
	file              []byte
	loadedTexNames    []string
	loadedTex         []Texture
	buffer            []byte
}

func NewModel(path string) *Model {
	m := new(Model)

	// Load the gltf file
	doc, err := gltf.Open(path)
	if err != nil {
		panic(err)

	}
	m.file = []byte(path)
	m.buffer = m.getData()
	m.data = doc
	m.traverseNode(0, mgl32.Ident4())
	return m

}

func (m *Model) Draw(s *shader.Shader) {
	for i := 0; i < len(m.Meshes); i++ {
		m.Meshes[i].Draw(s)
	}
}

func (m *Model) loadMesh(meshInd uint32) {

	posAccInd := m.data.Meshes[meshInd].Primitives[0].Attributes["POSITION"]
	normalAccInd := m.data.Meshes[meshInd].Primitives[0].Attributes["NORMAL"]
	texCoordAccInd := m.data.Meshes[meshInd].Primitives[0].Attributes["TEXCOORD_0"]
	indicesAccInd := m.data.Meshes[meshInd].Primitives[0].Indices

	posVec := m.getFloats(m.data.Accessors[posAccInd], []byte{0})
	positions := m.groupFloatsVec3(posVec)

	normalVec := m.getFloats(m.data.Accessors[normalAccInd], []byte{0})
	normals := m.groupFloatsVec3(normalVec)

	texCoordVec := m.getFloats(m.data.Accessors[texCoordAccInd], []byte{0})
	texCoords := m.groupFloatsVec2(texCoordVec)

	vertices := m.assembleVertices(positions, normals, texCoords)

	indices := m.getIndices(m.data.Accessors[int32(*indicesAccInd)], []byte{0})

	textures := m.getTextures()

	m.Meshes = append(m.Meshes, Mesh{
		Vertices: vertices,
		Indices:  indices,
		Textures: textures,
	})

}

func (m *Model) traverseNode(nextNode uint32, matrix mgl32.Mat4) {

	node := m.data.Nodes[nextNode]

	translation := mgl32.Vec3{0, 0, 0}
	transValues := node.Translation
	if transValues != [3]float64{} {
		translation = mgl32.Vec3{float32(transValues[0]), float32(transValues[1]), float32(transValues[2])}
	}

	rotation := mgl32.Quat{1, mgl32.Vec3{0, 0, 0}}
	rotValues := node.Rotation
	if rotValues != [4]float64{} {
		rotation = mgl32.Quat{float32(rotValues[3]), mgl32.Vec3{float32(rotValues[0]), float32(rotValues[1]), float32(rotValues[2])}}

	}

	scale := mgl32.Vec3{1, 1, 1}
	scaleValues := node.Scale
	if scaleValues != [3]float64{} {
		scale = mgl32.Vec3{float32(scaleValues[0]), float32(scaleValues[1]), float32(scaleValues[2])}

	}

	matNode := mgl32.Ident4()
	matValues := node.Matrix
	if matValues != [16]float64{} {
		matNode = mgl32.Mat4{
			float32(matValues[0]), float32(matValues[1]), float32(matValues[2]), float32(matValues[3]),
			float32(matValues[4]), float32(matValues[5]), float32(matValues[6]), float32(matValues[7]),
			float32(matValues[8]), float32(matValues[9]), float32(matValues[10]), float32(matValues[11]),
			float32(matValues[12]), float32(matValues[13]), float32(matValues[14]), float32(matValues[15]),
		}

	}

	trans := mgl32.Translate3D(translation[0], translation[1], translation[2])
	rot := rotation.Mat4()
	sca := mgl32.Scale3D(scale[0], scale[1], scale[2])

	matNextNode := matrix.Mul4(matNode).Mul4(trans).Mul4(rot).Mul4(sca)

	if node.Mesh != nil {
		m.translationMeshes = append(m.translationMeshes, translation)
		m.rotationMeshes = append(m.rotationMeshes, rotation)
		m.scaleMeshes = append(m.scaleMeshes, scale)
		m.matricesMeshes = append(m.matricesMeshes, matNextNode)

		m.loadMesh(*node.Mesh)
	}

	if node.Children != nil {
		for _, child := range node.Children {
			m.traverseNode(child, matNextNode)
		}

	}
}

func (m *Model) getData() []byte {
	var bytesText []byte
	var uri string = m.data.Buffers[0].URI

	var fileStr = m.file
	var fileDir = fileStr[:strings.LastIndex(string(fileStr), "/")+1]

	bytesText, err := os.ReadFile(string(fileDir) + uri)
	if err != nil {
		panic(err)
	}

	return bytesText

}

func (m *Model) getFloats(acr *gltf.Accessor, data []byte) []float32 {
	var floatVec []float32

	bufferViewInd := acr.BufferView
	count := acr.Count
	accByteOffset := acr.ByteOffset
	typing := acr.Type

	bufferView := m.data.BufferViews[int32(*bufferViewInd)]
	byteOffset := bufferView.ByteOffset

	var numPerVertex uint32

	switch typing {
	case gltf.AccessorScalar:
		numPerVertex = 1
	case gltf.AccessorVec2:
		numPerVertex = 2
	case gltf.AccessorVec3:
		numPerVertex = 3
	case gltf.AccessorVec4:
		numPerVertex = 4
	}

	var begOfData uint32 = uint32(byteOffset + accByteOffset)
	var lenOfData uint32 = uint32(count * numPerVertex * 4)

	for i := begOfData; i < begOfData+lenOfData; i += 4 {
		floatVec = append(floatVec, math.Float32frombits(binary.LittleEndian.Uint32(data[i:i+4])))
	}

	return floatVec

}

func (m *Model) getIndices(acr *gltf.Accessor, data []byte) []uint32 {
	var indices []uint32

	var bufferViewInd = acr.BufferView
	var count = acr.Count
	var accByteOffset = acr.ByteOffset
	var typing = acr.ComponentType

	bufferView := m.data.BufferViews[int32(*bufferViewInd)]
	byteOffset := bufferView.ByteOffset

	var begOfData uint32 = uint32(byteOffset + accByteOffset)
	switch typing {
	case gltf.ComponentUint:
		for i := begOfData; i < begOfData+uint32(count*4); i += 4 {
			indices = append(indices, binary.LittleEndian.Uint32(data[i:i+4]))

		}
	case gltf.ComponentUshort:
		for i := begOfData; i < begOfData+uint32(count*2); i += 2 {
			indices = append(indices, uint32(binary.LittleEndian.Uint16(data[i:i+2])))

		}
	case gltf.ComponentShort:
		for i := begOfData; i < begOfData+uint32(count*2); i += 2 {
			indices = append(indices, uint32(int32(binary.LittleEndian.Uint16(data[i:i+2]))))
		}
	}

	return indices
}

func (m *Model) getTextures() []Texture {
	var textures []Texture

	var fileStr = string(m.file)

	var fileDir = fileStr[:strings.LastIndex(fileStr, "/")+1]

	for _, image := range m.data.Images {

		var texPath = image.URI

		var skip bool = false

		for _, texName := range m.loadedTexNames {
			if texName == texPath {
				skip = true
				break
			}

		}

	}

	return textures

}

func (m *Model) assembleVertices(pos []mgl32.Vec3, normal []mgl32.Vec3, texCoord []mgl32.Vec2) []Vertex {
	var vertices []Vertex

	for i := 0; i < len(pos); i++ {
		vertex := Vertex{
			Position: pos[i],
			Normal:   normal[i],
			TexCoord: texCoord[i],
		}
		vertices = append(vertices, vertex)

	}

	return vertices
}

func (m *Model) groupFloatsVec2(data []float32) []mgl32.Vec2 {
	var vectors []mgl32.Vec2

	for i := 0; i < len(data); i += 2 {
		vectors = append(vectors, mgl32.Vec2{data[i], data[i+1]})

	}

	return vectors
}

func (m *Model) groupFloatsVec3(data []float32) []mgl32.Vec3 {
	var vectors []mgl32.Vec3

	for i := 0; i < len(data); i += 3 {
		vectors = append(vectors, mgl32.Vec3{data[i], data[i+1], data[i+2]})
	}

	return vectors
}

func (m *Model) groupFloatsVec4(data []float32) []mgl32.Vec4 {
	var vectors []mgl32.Vec4

	for i := 0; i < len(data); i += 4 {
		vectors = append(vectors, mgl32.Vec4{data[i], data[i+1], data[i+2], data[i+3]})
	}

	return vectors
}
