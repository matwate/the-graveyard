package buffers

import "github.com/go-gl/gl/v4.3-core/gl"

type Drawable struct {
	vertexPos       []float32
	vertexNormals   []float32
	vertexTexCoords []float32
	vertexData      []float32
	posSize         uint32
	normSize        uint32
	texSize         uint32
	vao             uint32
	vbo             uint32
	ebo             uint32
}

func NewDrawable(vertexPos []float32, vertexNormals []float32, vertexTexCoords []float32, posSize uint32, normSize uint32, texSize uint32) *Drawable {
	return &Drawable{
		vertexPos:       vertexPos,
		vertexNormals:   vertexNormals,
		vertexTexCoords: vertexTexCoords,
		posSize:         posSize,
		normSize:        normSize,
		texSize:         texSize,
	}
}

func (d *Drawable) Vao() (uint32, uint32, uint32) {
	d.mergeVertices()
	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(d.vertexData)*4, gl.Ptr(d.vertexData), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(d.vertexData)*4, gl.Ptr(d.vertexData), gl.STATIC_DRAW)

	d.makeAttribPointers()

	d.vao = vao
	d.vbo = vbo
	d.ebo = ebo

	return vao, vbo, ebo

}
func (d *Drawable) mergeVertices() []float32 {

	result := make([]float32, 0)
	aIndex, bIndex, cIndex := 0, 0, 0

	for {
		// Merge chunks from a
		for i := uint32(0); i < d.posSize && aIndex < len(d.vertexPos); i++ {
			result = append(result, d.vertexPos[aIndex])
			aIndex++
		}

		// Merge chunks from b
		for i := uint32(0); i < d.normSize && bIndex < len(d.vertexNormals); i++ {
			result = append(result, d.vertexNormals[bIndex])
			bIndex++
		}

		// Merge chunks from c
		for i := uint32(0); i < d.texSize && cIndex < len(d.vertexTexCoords); i++ {
			result = append(result, d.vertexTexCoords[cIndex])
			cIndex++
		}

		// Exit if all slices have been processed
		if aIndex == len(d.vertexPos) && bIndex == len(d.vertexNormals) && cIndex == len(d.vertexTexCoords) {
			break
		}
	}
	d.vertexData = result

	return result
}

func (d *Drawable) makeAttribPointers() {
	//Pos attrib
	gl.VertexAttribPointerWithOffset(0, int32(d.posSize), gl.FLOAT, false, int32(d.posSize+d.normSize+d.texSize), 0)
	gl.EnableVertexAttribArray(0)

	//Norm attrib
	gl.VertexAttribPointerWithOffset(0, int32(d.normSize), gl.FLOAT, false, int32(d.posSize+d.normSize+d.texSize), uintptr(d.posSize))
	gl.EnableVertexAttribArray(1)

	//Tex attrib
	gl.VertexAttribPointerWithOffset(0, int32(d.texSize), gl.FLOAT, false, int32(d.posSize+d.normSize+d.texSize), uintptr(d.posSize+d.normSize))
	gl.EnableVertexAttribArray(2)
}

func (d *Drawable) Draw() {
	gl.BindVertexArray(d.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(d.vertexPos)/3))
}
