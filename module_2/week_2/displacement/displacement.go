package main

import (
	"fmt"
)

func GenDisplaceFn(a, v, d float64) func (float64) float64 {
	fn := func(t float64) float64 {
		return float64(1/2 * a * (t * t)) + (v * t) + d
	}
	return fn
}

func main() {
	var acceleration, velocity, displacement float64

	fmt.Println("Please enter acceleration: ")
	fmt.Scan(&acceleration)
	fmt.Println("Please enter initial velocity: ")
	fmt.Scan(&velocity)
	fmt.Println("Please enter initial displacement: ")
	fmt.Scan(&displacement)

	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}
