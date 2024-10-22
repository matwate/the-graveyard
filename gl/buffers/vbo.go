package buffers

import "github.com/go-gl/gl/v4.3-core/gl"

type VBO struct {
	Id uint32
}

func NewVBO(vertices []float32, size int) *VBO {

	vbo := &VBO{}

	gl.GenBuffers(1, &vbo.Id)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.Id)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(vertices), gl.STATIC_DRAW)

	return vbo
}

func (v *VBO) Bind() {

	gl.BindBuffer(gl.ARRAY_BUFFER, v.Id)

}

func (v *VBO) Unbind() {

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

}

func (v *VBO) Delete() {

	gl.DeleteBuffers(1, &v.Id)

}
