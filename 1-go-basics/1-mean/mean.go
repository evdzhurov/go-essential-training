package main

import "fmt"

func main() {

	// Integer variable definition and assignment
	/*
		var x int
		var y int

		x = 1
		y = 2
	*/

	// Floating point variable definition and assignment
	/*
		var x float64
		var y float64

		x = 1.0
		y = 2.0
	*/

	// Floating point variable definition and assignment with automatic inference
	x := 1.0
	y := 2.0

	// Same as above
	// x, y := 1.0, 2.0

	// Invalid for later calculations because different types are inferred.
	// x, y := 1, 2.0

	fmt.Printf("x=%v type=%T\n", x, x)
	fmt.Printf("y=%v type=%T\n", y, y)

	mean := (x + y) / 2
	fmt.Printf("mean=%v type=%T\n", mean, mean)
}
