package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const HARDCODED_GROUND = 450

/*
Madeline is a fighting game character, it has the ability to dash once
before she touches the ground, she can jump (so far)
*/
type Madeline struct {
	Position     rl.Vector2
	Velocity     rl.Vector2
	MaxMoveSpeed float32

	CanJump bool
	// Madeline's dash ability
	DashCount int
	DashSpeed float32
}

func main() {
	rl.InitWindow(1600, 900, "Madeline prototype")
	rl.SetTargetFPS(60)

	m := Madeline{
		Position:     rl.NewVector2(100, 450),
		MaxMoveSpeed: 5,
		DashSpeed:    15,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Madeline prototype", 10, 10, 20, rl.DarkGray)
		m.Update()
		rl.EndDrawing()
	}

}

func (m *Madeline) Update() {
	m.HandleInput()
	// Update position
	m.UpdatePhysics()

	m.Draw()
}

func (m *Madeline) HandleInput() {
	//Move with wasd, jump with space, dash with j
	// It is a platform fighter so no up or down
	if rl.IsKeyDown(rl.KeyW) {
		// Madeline is looking up
	} else if rl.IsKeyDown(rl.KeyS) {
		// Madeline is looking down
	}
	if rl.IsKeyDown(rl.KeyA) {
		m.Velocity.X = -1 * m.MaxMoveSpeed
	} else if rl.IsKeyDown(rl.KeyD) {
		m.Velocity.X = 1 * m.MaxMoveSpeed
	} else {
		m.Velocity.X = 0
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		// Jump
		if m.CanJump {
			m.Velocity.Y = -10
			m.CanJump = false
		}
	}

	if rl.IsKeyPressed(rl.KeyJ) {
		// Dash
		if m.DashCount < 1 {
			m.DashCount++
		}
	}

}

func (m *Madeline) UpdatePhysics() {
	// Update position
	if m.DashCount > 0 {
		m.Velocity.X = m.DashSpeed * rl.Vector2Normalize(m.Velocity).X
	}

	m.Position = rl.Vector2Add(m.Position, m.Velocity)

	// Gravity
	if m.Position.Y < HARDCODED_GROUND {
		m.Velocity.Y += 0.5
		if m.Velocity.Y > 10 {
			m.Velocity.Y = 10
		}

		if m.DashCount > 0 {
			m.Velocity.Y = 0
		}
	} else {
		m.Velocity.Y = 0
		m.Position.Y = HARDCODED_GROUND
		m.CanJump = true
	}

	if math.Abs(float64(m.Velocity.X)) > float64(m.MaxMoveSpeed) {
		m.Velocity.X -= 0.2 * rl.Vector2Normalize(m.Velocity).Xp0
	}

	if m.Velocity.X < m.MaxMoveSpeed {
		m.DashCount = 0
	}
}

func (m *Madeline) Draw() {
	rl.DrawRectangle(int32(m.Position.X), int32(m.Position.Y), 50, 75, rl.Red)
	// Debug madelines stats
	rl.DrawText(fmt.Sprintf("Position: %v", m.Position), 10, 30, 20, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("Velocity: %v", m.Velocity), 10, 50, 20, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("CanJump: %v", m.CanJump), 10, 70, 20, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("DashCount: %v", m.DashCount), 10, 90, 20, rl.DarkGray)
	rl.DrawFPS(1570, 10)
}
