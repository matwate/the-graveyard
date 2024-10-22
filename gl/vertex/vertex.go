package vertex

import "github.com/go-gl/mathgl/mgl32"

type Vertex struct {
	Position  mgl32.Vec3
	Normal    mgl32.Vec3
	TexCoords mgl32.Vec2
}

func NewVertex(position, normal mgl32.Vec3, texCoords mgl32.Vec2) *Vertex {

	v := &Vertex{
		Position:  position,
		Normal:    normal,
		TexCoords: texCoords,
	}

	return v

}

func (v *Vertex) Data() []float32 {

	data := []float32{
		v.Position.X(), v.Position.Y(), v.Position.Z(),
		v.Normal.X(), v.Normal.Y(), v.Normal.Z(),
		v.TexCoords.X(), v.TexCoords.Y(),
	}

	return data
}
