package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawCatmullRomInterpolating(points []rl.Vector2, thick float32, color rl.Color) {

	// Take the second point , reflect it across the first point, use that as the first point
	// Then the second to last point , reflect it across the last point, use that as the last point

	SecondPoint := points[1]
	FirstPoint := rl.Vector2{
		X: SecondPoint.X + (SecondPoint.X - points[0].X),
		Y: SecondPoint.Y + (SecondPoint.Y - points[0].Y),
	}
	LastPoint := points[len(points)-2]
	LastPoint = rl.Vector2{
		X: LastPoint.X + (LastPoint.X - points[len(points)-1].X),
		Y: LastPoint.Y + (LastPoint.Y - points[len(points)-1].Y),
	}

	NewSpline := []rl.Vector2{FirstPoint}
	NewSpline = append(NewSpline, points...)
	NewSpline = append(NewSpline, LastPoint)

	rl.DrawSplineCatmullRom(NewSpline, thick, color)
}
