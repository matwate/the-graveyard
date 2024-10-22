package buffers

import "github.com/go-gl/gl/v4.3-core/gl"

type VAO struct {
	Id uint32
}

func NewVAO() *VAO {

	v := &VAO{}
	gl.GenVertexArrays(1, &v.Id)
	return v

}

func (v *VAO) LinkAttribute(vbo VBO, index uint32, size int32, dataType uint32, stride int32, offset int) {

	vbo.Bind()
	gl.VertexAttribPointerWithOffset(index, size, dataType, false, stride, uintptr(offset))
	gl.EnableVertexAttribArray(index)
	vbo.Unbind()

}

func (v *VAO) Bind() {

	gl.BindVertexArray(v.Id)
}

func (v *VAO) Unbind() {

	gl.BindVertexArray(0)
}

func (v *VAO) Delete() {

	gl.DeleteVertexArrays(1, &v.Id)
}
