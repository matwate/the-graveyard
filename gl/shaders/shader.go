package shaders

import (
	"os"

	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	ID uint32
}

func NewShader(vertexPath string, fragmentPath string) *Shader {

	var vxCode []byte
	var frCode []byte

	vxCode, err := os.ReadFile(vertexPath)
	if err != nil {
		panic(err)
	}
	frCode, err = os.ReadFile(fragmentPath)
	if err != nil {
		panic(err)
	}

	vertex := gl.CreateShader(gl.VERTEX_SHADER)
	fragment := gl.CreateShader(gl.FRAGMENT_SHADER)

	cVStr, free := gl.Strs(string(vxCode) + "\x00")
	gl.ShaderSource(vertex, 1, cVStr, nil)
	free()

	cFStr, free := gl.Strs(string(frCode) + "\x00")
	gl.ShaderSource(fragment, 1, cFStr, nil)
	free()

	gl.CompileShader(vertex)
	checkCompileErrors(vertex, "VERTEX")
	gl.CompileShader(fragment)
	checkCompileErrors(fragment, "FRAGMENT")

	id := gl.CreateProgram()
	gl.AttachShader(id, vertex)
	gl.AttachShader(id, fragment)
	gl.LinkProgram(id)
	checkCompileErrors(id, "PROGRAM")

	gl.DeleteShader(vertex)
	gl.DeleteShader(fragment)

	return &Shader{
		ID: id,
	}

}

func (s *Shader) Use() {
	gl.UseProgram(s.ID)
}

func (s *Shader) SetBool(name string, value bool) {

	var a int32
	if value {
		a = 1
	} else {
		a = 0

	}
	gl.Uniform1i(gl.GetUniformLocation(s.ID, gl.Str(name+"\x00")), a)
}

func (s *Shader) SetInt(name string, value int32) {

	gl.Uniform1i(gl.GetUniformLocation(s.ID, gl.Str(name+"\x00")), value)
}

func (s *Shader) SetFloat(name string, value float32) {

	gl.Uniform1f(gl.GetUniformLocation(s.ID, gl.Str(name+"\x00")), value)
}

func (s *Shader) SetVec2(name string, value mgl32.Vec2) {
	gl.Uniform2fv(gl.GetUniformLocation(s.ID, gl.Str(name+"\x00")), 1, &value[0])
}

func (s *Shader) SetVec3(name string, value mgl32.Vec3) {
	gl.Uniform3fv(gl.GetUniformLocation(s.ID, gl.Str(name+"\x00")), 1, &value[0])
}

func (s *Shader) SetVec4(name string, x float32, y float32, z float32, w float32) {
	gl.Uniform4fv(gl.GetUniformLocation(s.ID, gl.Str(name+"\x00")), 1, &[]float32{x, y, z, w}[0])
}

func (s *Shader) SetMat2(name string, value mgl32.Mat2) {
	gl.UniformMatrix2fv(gl.GetUniformLocation(s.ID, gl.Str(name+"\x00")), 1, false, &value[0])
}
func (s *Shader) SetMat3(name string, value mgl32.Mat3) {
	gl.UniformMatrix3fv(gl.GetUniformLocation(s.ID, gl.Str(name+"\x00")), 1, false, &value[0])
}
func (s *Shader) SetMat4(name string, value mgl32.Mat4) {
	gl.UniformMatrix4fv(gl.GetUniformLocation(s.ID, gl.Str(name+"\x00")), 1, false, &value[0])
}

func (s *Shader) Delete() {
	gl.DeleteProgram(s.ID)
}

func checkCompileErrors(shader uint32, shaderType string) {
	var success int32
	var infoLog [1024]byte
	if shaderType != "PROGRAM" {
		gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success)
		if success == 0 {
			gl.GetShaderInfoLog(shader, 1024, nil, &infoLog[0])
			panic("ERROR::SHADER_COMPILATION_ERROR of type: " + shaderType + "\n" + string(infoLog[:]))
		}
	} else {
		gl.GetProgramiv(shader, gl.LINK_STATUS, &success)
		if success == 0 {
			gl.GetProgramInfoLog(shader, 1024, nil, &infoLog[0])
			panic("ERROR::PROGRAM_LINKING_ERROR of type: " + shaderType + "\n" + string(infoLog[:]))
		}
	}
}
