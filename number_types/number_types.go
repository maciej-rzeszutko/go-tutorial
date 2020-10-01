package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	var i uint8 = 255
	var f float32 = 5.6
	fmt.Println(float32(i) + f)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// The expression T(v) converts the value v to the type T.
	var x, y int = 3, 4
	var fl float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(fl)
	fmt.Println(x, y, z)

}
