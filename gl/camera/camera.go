package camera

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	yaw         = -90.0
	pitch       = 0.0
	speed       = 2.5
	sensitivity = 0.1
	zoom        = 45.0
)

const (
	Forward = iota
	Backward
	Left
	Right
)

type Camera struct {
	Pos         mgl32.Vec3
	Front       mgl32.Vec3
	Up          mgl32.Vec3
	Right       mgl32.Vec3
	WorldUp     mgl32.Vec3
	Yaw         float32
	Pitch       float32
	MoveSpeed   float32
	Sensitivity float32
	Zoom        float32
	LastX       float32
	LastY       float32
	FirstMouse  bool
}

func NewCamera(pos mgl32.Vec3, up mgl32.Vec3, yaw float32, pitch float32) *Camera {
	c := &Camera{
		Pos:         pos,
		WorldUp:     up,
		Yaw:         yaw,
		Pitch:       pitch,
		Front:       mgl32.Vec3{0.0, 0.0, -1.0},
		MoveSpeed:   speed,
		Sensitivity: sensitivity,
		Zoom:        zoom,
		Up:          up,
		Right:       mgl32.Vec3{},
		LastX:       800.0 / 2.0,
		LastY:       600.0 / 2.0,
		FirstMouse:  true,
	}
	return c
}

func NewCameraScalars(posX, posY, posZ, upX, upY, upZ, yaw, pitch float32) *Camera {
	return NewCamera(mgl32.Vec3{posX, posY, posZ}, mgl32.Vec3{upX, upY, upZ}, yaw, pitch)
}

func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	PF := c.Pos.Add(c.Front)
	return mgl32.LookAt(c.Pos[0], c.Pos[1], c.Pos[2], PF[0], PF[1], PF[2], c.Up[0], c.Up[1], c.Up[2])

}

func (c *Camera) ProcessKeyboard(direction int, deltaTime float32) {
	velocity := c.MoveSpeed * deltaTime
	if direction == Forward {
		c.Pos = c.Pos.Add(c.Front.Mul(velocity))
	}
	if direction == Backward {
		c.Pos = c.Pos.Sub(c.Front.Mul(velocity))
	}
	if direction == Left {
		c.Pos = c.Pos.Sub(c.Right.Mul(velocity))
	}
	if direction == Right {
		c.Pos = c.Pos.Add(c.Right.Mul(velocity))
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

	c.updateCameraVectors()
}

func (c *Camera) ProcessMouseScroll(yoffset float32) {
	c.Zoom -= yoffset
	if c.Zoom < 1.0 {
		c.Zoom = 1.0
	}
	if c.Zoom > 45.0 {
		c.Zoom = 45.0
	}
}

func (c *Camera) updateCameraVectors() {
	var front mgl32.Vec3
	front[0] = float32(math.Cos(float64(mgl32.DegToRad(float32(c.Yaw)))) * math.Cos(float64(mgl32.DegToRad(float32(c.Pitch)))))
	front[1] = float32(math.Sin(float64(mgl32.DegToRad(float32(c.Pitch)))))
	front[2] = float32(math.Sin(float64(mgl32.DegToRad(float32(c.Yaw)))) * math.Cos(float64(mgl32.DegToRad(float32(c.Pitch)))))
	c.Front = front.Normalize()
	c.Right = c.Front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
}
