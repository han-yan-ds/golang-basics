package main

import (
	"fmt"
	shape "golang-basics/shape"
)

func main() {
	sampleTriangle := shape.Triangle{
		Height: 1.5,
		Base:   3.0,
	}
	fmt.Print(sampleTriangle.GetArea())
}
