package textures

import (
	shaders "matwa/graphics-engine/shader"

	"github.com/go-gl/gl/v4.3-core/gl"
	"neilpa.me/go-stbi"
)

type Texture struct {
	Id     uint32
	typing string
}

func NewTexture(image string, texType string, slot int) *Texture {

	t := &Texture{}

	t.typing = texType

	img, err := stbi.Load(image)
	if err != nil {
		panic(err)

	}

	gl.GenTextures(1, &t.Id)
	gl.ActiveTexture(gl.TEXTURE0 + uint32(slot))
	gl.BindTexture(gl.TEXTURE_2D, t.Id)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(img.Rect.Dx()), int32(img.Rect.Dy()), 0, gl.RGBA, gl.UNSIGNED_INT, gl.Ptr(img.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return t

}

func (t *Texture) textUnit(s shaders.Shader, unifrom string, unit int32) {
	texUni := gl.GetUniformLocation(s.Id, gl.Str(unifrom+"\x00"))
	s.Activate()
	gl.Uniform1i(texUni, unit)
}

func (t *Texture) Bind() {

	gl.BindTexture(gl.TEXTURE_2D, t.Id)
}

func (t *Texture) Unbind() {

	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (t *Texture) Delete() {

	gl.DeleteTextures(1, &t.Id)
}
