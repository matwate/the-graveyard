package cameras

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	Forward = iota
	Backward
	Left
	Right
)

const (
	yaw         = -90.0
	pitch       = 0.0
	speed       = 2.5
	sensitivity = 0.1
	zoom        = 45.0
)

type Camera struct {
	Position    mgl32.Vec3
	Front       mgl32.Vec3
	Up          mgl32.Vec3
	Right       mgl32.Vec3
	WorldUp     mgl32.Vec3
	Yaw         float32
	Pitch       float32
	MovSped     float32
	Sensitivity float32
	Zoom        float32
}

func NewCamera(
	pos mgl32.Vec3,
	up mgl32.Vec3,
	yaw float32,
	pitch float32,
	front mgl32.Vec3,
	movSped float32,
	sens float32,
	zoom float32,

) *Camera {

	c := &Camera{}
	c.Position = pos
	c.WorldUp = up
	c.Yaw = yaw
	c.Pitch = pitch
	c.Front = front
	c.MovSped = movSped
	c.Sensitivity = sens
	c.Zoom = zoom
	c.UpdateCameraVectors()
	return c
}

func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(c.Position, c.Position.Add(c.Front), c.Up)
}

func (c *Camera) ProcessKeyboard(direction int, deltaTime float32) {

	velocity := c.MovSped * deltaTime

	if direction == Forward {
		c.Position = c.Position.Add(c.Front.Mul(velocity))
	}
	if direction == Backward {
		c.Position = c.Position.Sub(c.Front.Mul(velocity))
	}
	if direction == Left {
		c.Position = c.Position.Sub(c.Right.Mul(velocity))
	}
	if direction == Right {
		c.Position = c.Position.Add(c.Right.Mul(velocity))
	}
}

func (c *Camera) ProcessMouseMovement(xoffset, yoffset float32, constrainPitch bool) {
	xoffset *= c.Sensitivity
	yoffset *= c.Sensitivity

	c.Yaw += xoffset
	c.Pitch += yoffset

	if constrainPitch {
		if c.Pitch > 89.0 {
			c.Pitch = 89.0
		}
		if c.Pitch < -89.0 {
			c.Pitch = -89.0
		}
	}

	c.UpdateCameraVectors()
}

func (c *Camera) ProcessMouseScroll(yoffset float32) {
	if c.Zoom >= 1.0 && c.Zoom <= 45.0 {
		c.Zoom -= yoffset
	}
	if c.Zoom <= 1.0 {
		c.Zoom = 1.0
	}
	if c.Zoom >= 45.0 {
		c.Zoom = 45.0
	}
}

func (c *Camera) UpdateCameraVectors() {
	front := mgl32.Vec3{
		float32(math.Cos(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch)))),
		float32(math.Sin(float64(mgl32.DegToRad(c.Pitch)))),
		float32(math.Sin(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch)))),
	}

	c.Front = front.Normalize()
	c.Right = c.Front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
}
