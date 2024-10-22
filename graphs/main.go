package main

import (
	"gonum.org/v1/plot"
)

func main() {
	p := plot.New()
	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

}
