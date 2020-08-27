package main

import (
	"fmt"
	"math"
)

type A struct {
	x float64
	y float64
}

func (a A) abs() float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y)
}

func main() {
	a := A{
		3,
		4,
	}

	fmt.Println(a.abs())
}
