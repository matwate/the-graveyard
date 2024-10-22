package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Framecount int

var s rl.Shader

var GlobalPath []rl.Vector2 = []rl.Vector2{}

func main() {
	rl.InitWindow(1600, 900, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)
	circleImage := rl.GenImageColor(20, 20, rl.Blank)
	rl.ImageDrawCircle(circleImage, 5, 5, 5, rl.White)
	circleTexture := rl.LoadTextureFromImage(circleImage)

	MakeRandomAttractors(3)
	MakeRandomParticles(10)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		UpdateParticles()

		DrawParticles(circleTexture)
		DrawAttractors()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
		Framecount++
	}

	rl.CloseWindow()

}

type Attractor struct {
	Position rl.Vector2
	Mass     float32
	force    rl.Vector2
}

type Particle struct {
	Attractor
	Velocity rl.Vector2
	Accel    rl.Vector2
	forces   rl.Vector2
	path     []rl.Vector2
}

var particles []*Particle = []*Particle{}
var attractors []*Attractor = []*Attractor{}

func MakeRandomAttractors(n int) {
	for i := 0; i < n; i++ {
		attractors = append(attractors, &Attractor{rl.NewVector2(float32(rl.GetRandomValue(0, 1600)), float32(rl.GetRandomValue(0, 900))), 30, rl.NewVector2(0, 0)})
	}
}

func SetDefaultAttractor() {
	attractors = append(attractors, &Attractor{rl.NewVector2(800, 450), 30, rl.NewVector2(0, 0)})
}

func MakeRandomParticles(n int) {
	for i := 0; i < n; i++ {
		particles = append(particles, &Particle{Attractor{rl.NewVector2(float32(rl.GetRandomValue(0, 1600)), float32(rl.GetRandomValue(0, 900))), 0.2, rl.NewVector2(10, 20)}, rl.NewVector2(10, 0), rl.NewVector2(0, 0), rl.NewVector2(0, 0), []rl.Vector2{}})
	}
}

func MakeParticlesCT(n int) {
	// This is just to demostrate chaos theory, a bunch of particles really close to each other
	for i := 0; i < n; i++ {
		particles = append(particles, &Particle{
			Attractor{
				rl.NewVector2(float32(800)+float32(0.2*float32(i)), 450),
				0.2,
				rl.NewVector2(0, 0),
			},
			rl.NewVector2(10, 0),
			rl.NewVector2(0, 0),
			rl.NewVector2(0, 0),
			[]rl.Vector2{},
		})
	}
}

func MakeParticlesPerPixel() {
	for x := 0; x < 1600; x++ {
		for y := 0; y < 900; y++ {
			particles = append(particles, &Particle{Attractor{rl.NewVector2(float32(x), float32(y)), 0.2, rl.NewVector2(0, 0)}, rl.NewVector2(10, 0), rl.NewVector2(0, 0), rl.NewVector2(0, 0), []rl.Vector2{}})
		}
	}
}

func (p *Particle) Update(attractors []*Attractor, dt float32) {

	p.Accel = rl.Vector2Scale(p.forces, 1/p.Mass)

	p.Velocity = rl.Vector2Add(p.Velocity, p.Accel)
	p.Velocity = rl.Vector2Scale(p.Velocity, 0.010)
	v := rl.Vector2Scale(p.Velocity, dt)
	a := rl.Vector2Scale(p.Accel, 0.5*dt*dt)
	p.Position = rl.Vector2Add(p.Position, rl.Vector2Add(v, a))

	for _, a := range attractors {
		force := rl.Vector2Subtract(a.Position, p.Position)
		force = rl.Vector2Normalize(force)

		force = rl.Vector2Scale(force, a.Mass/p.Mass)
		p.forces = rl.Vector2Add(p.forces, force)
		a.force = force

	}

	// Every 5 frames pos is saved
	if Framecount%2 == 0 {
		p.path = append(p.path, p.Position)
		GlobalPath = append(GlobalPath, p.Position)
	}

}

func UpdateParticles() {
	for _, p := range particles {
		p.Update(attractors, rl.GetFrameTime())

	}
}

func DrawParticles(tex rl.Texture2D) {
	for _, p := range particles {
		rl.DrawTexture(tex, int32(p.Position.X), int32(p.Position.Y), LerpColor(rl.Blue, rl.White, rl.Vector2Length(p.Velocity)/750))
	}
}

func DrawAttractors() {
	for _, a := range attractors {
		rl.DrawCircleV(a.Position, 10, rl.Red)
	}
}

func DrawPath() {
	for _, p := range particles {
		if len(p.path) > 3 {
			DrawCatmullRomInterpolating(p.path, 2, rl.White)
		}

		if len(p.path) > 30 {
			p.path = p.path[1:]
		}
	}

}

func LerpColor(c1, c2 rl.Color, t float32) rl.Color {
	if t > 1 {
		t = 1
	}
	return rl.NewColor(
		uint8(float32(c1.R)*(1-t)+float32(c2.R)*t),
		uint8(float32(c1.G)*(1-t)+float32(c2.G)*t),
		uint8(float32(c1.B)*(1-t)+float32(c2.B)*t),
		255,
	)
}

// Here starts the gravity fractal code
/*
	Explanation:
	We're going to use a shader to determine for each pixel of the screen
	what attractor is going to end up closer to it after a certain number of iterations.

	Correction:
	We was going to use a shader but we're going to use the CPU instead

*/
func CalculateNearestAttractor(p *Particle, attractors []*Attractor, iterations int) *Attractor {
	var nearest *Attractor
	var minDistance float32 = 1000000
	for range iterations {
		p.Update(attractors, 1/60.0)
	}

	for _, a := range attractors {
		distance := rl.Vector2Distance(p.Position, a.Position)
		minDistance := min(minDistance, distance)
		if distance == minDistance {
			nearest = a
		}
	}
	return nearest
}
