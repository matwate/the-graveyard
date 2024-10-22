package buffers

import "github.com/go-gl/gl/v4.3-core/gl"

type EBO struct {
	Id uint32
}

func NewEBO(indices []uint32) *EBO {

	ebo := &EBO{}

	gl.GenBuffers(1, &ebo.Id)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo.Id)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	return ebo
}

func (e *EBO) Bind() {

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, e.Id)

}

func (e *EBO) Unbind() {

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (e *EBO) Delete() {

	gl.DeleteBuffers(1, &e.Id)
}
