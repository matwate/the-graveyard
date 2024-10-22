package shaders

import (
	"os"

	"github.com/go-gl/gl/v4.3-core/gl"
)

func GetFileContents(path string) (string, error) {

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil

}

type Shader struct {
	Id uint32
}

func NewShader(vertexPath, fragmentPath string) (*Shader, error) {

	vertexCode, _ := GetFileContents(vertexPath)
	fragmentCode, _ := GetFileContents(fragmentPath)

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	csource, free := gl.Strs(vertexCode)

	gl.ShaderSource(vertexShader, 1, csource, nil)
	free()
	gl.CompileShader(vertexShader)

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	csource, free = gl.Strs(fragmentCode)

	gl.ShaderSource(fragmentShader, 1, csource, nil)
	free()
	gl.CompileShader(fragmentShader)

	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &Shader{Id: shaderProgram}, nil

}

func (s *Shader) Activate() {

	gl.UseProgram(s.Id)
}

func (s *Shader) Delete() {
	gl.DeleteProgram(s.Id)
}
