package main

import (
	shape "golang-basics/shape"
)

func main() {
	sampleTriangle := shape.Triangle{
		Height: 1.5,
		Base:   3.0,
	}
	shape.PrintArea(sampleTriangle)
}
