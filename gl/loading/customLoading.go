package loading

import (
	"encoding/binary"
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

func LoadFile(path string) {

	doc, _ := gltf.Open(path)
	scene := doc.Scenes[0]

	for _, nodeIndex := range scene.Nodes {
		ParseNode(doc.Nodes[nodeIndex], doc)
	}
}

func ParseNode(node *gltf.Node, doc *gltf.Document) {

	for _, childIndex := range node.Children {
		ParseNode(doc.Nodes[childIndex], doc)
	}

	if node.Mesh != nil {
		meshIndex := int32(*node.Mesh)
		ParseMesh(doc.Meshes[meshIndex], doc)
	}
}

func ParseMesh(mesh *gltf.Mesh, doc *gltf.Document) {

	for _, primitive := range mesh.Primitives {
		pos, norm, tex, ind := ParsePrimitive(primitive, doc)
		fmt.Println(pos, norm, tex, ind)
	}
}

func ParsePrimitive(primitive *gltf.Primitive, doc *gltf.Document) (pos []mgl32.Vec3, norm []mgl32.Vec3, tex []mgl32.Vec2, ind []uint32) {

	accessor := doc.Accessors[primitive.Attributes["POSITION"]]
	bufferView := doc.BufferViews[int32(*accessor.BufferView)]
	positions, err := modeler.ReadBufferView(doc, bufferView)
	if err != nil {
		fmt.Printf("could not read buffer view: %v", err)
	}

	accessor = doc.Accessors[primitive.Attributes["NORMAL"]]
	bufferView = doc.BufferViews[int32(*accessor.BufferView)]
	normals, err := modeler.ReadBufferView(doc, bufferView)
	if err != nil {
		fmt.Printf("could not read buffer view: %v", err)
	}

	accessor = doc.Accessors[primitive.Attributes["TEXCOORD_0"]]
	bufferView = doc.BufferViews[int32(*accessor.BufferView)]
	texCoords, err := modeler.ReadBufferView(doc, bufferView)
	if err != nil {
		fmt.Printf("could not read buffer view: %v", err)
	}

	indices := doc.Accessors[int32(*primitive.Indices)]
	bufferViewIndices := doc.BufferViews[int32(*indices.BufferView)]
	bufferIndices := doc.Buffers[bufferViewIndices.Buffer]
	bufferIndicesData := make([]uint32, len(bufferIndices.Data)/4)

	var idxs = make([]uint32, len(bufferIndices.Data)/4)

	for i := 0; i < len(bufferIndices.Data); i += 4 {
		idxs[i/4] = binary.LittleEndian.Uint32(bufferIndices.Data[i : i+4])
	}
	idxs, _ = modeler.ReadIndices(doc, indices, bufferIndicesData)

	pos = vector3Ize(positions)
	norm = vector3Ize(normals)
	tex = vector2Ize(texCoords)
	ind = idxs

	return
}

func vector2Ize(v []byte) []mgl32.Vec2 {
	vecs := make([]mgl32.Vec2, len(v)/8)
	for i := 0; i < len(v); i += 8 {
		vecs[i/8] = mgl32.Vec2{float32(binary.LittleEndian.Uint32(v[i : i+4])), float32(binary.LittleEndian.Uint32(v[i+4 : i+8]))}
	}
	return vecs
}

func vector3Ize(v []byte) []mgl32.Vec3 {
	vecs := make([]mgl32.Vec3, len(v)/12)
	for i := 0; i < len(v); i += 12 {
		vecs[i/12] = mgl32.Vec3{float32(binary.LittleEndian.Uint32(v[i : i+4])), float32(binary.LittleEndian.Uint32(v[i+4 : i+8])), float32(binary.LittleEndian.Uint32(v[i+8 : i+12]))}
	}
	return vecs
}

func vector4Ize(v []byte) []mgl32.Vec4 {
	vecs := make([]mgl32.Vec4, len(v)/16)
	for i := 0; i < len(v); i += 16 {
		vecs[i/16] = mgl32.Vec4{float32(binary.LittleEndian.Uint32(v[i : i+4])), float32(binary.LittleEndian.Uint32(v[i+4 : i+8])), float32(binary.LittleEndian.Uint32(v[i+8 : i+12])), float32(binary.LittleEndian.Uint32(v[i+12 : i+16]))}
	}
	return vecs
}
