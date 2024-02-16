package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func main() {
	// main_imports()
	// main_functions()
	// main_namedParams()
	// main_vars()
	// main_basicTypes()
	// main_zeroValues()
	// main_typeConversions()
	// main_typeInference()
	// main_constants()
	main_numericConstants()
}

func main_imports() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

func add(x int, y int) int {
	return x + y
}
func main_functions() {
	fmt.Println(add(42, 13))
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main_namedParams() {
	fmt.Println(split(17))
}

func main_vars() {
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	k := 3
	fmt.Println(i, j, c, python, java, k)
}

func main_basicTypes() {
	/*
			bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // alias for uint8

		rune // alias for int32
		     // represents a Unicode code point

		float32 float64

		complex64 complex128

		The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
	*/

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func main_zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func main_typeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func main_typeInference() {
	v := "42" // change me!
	fmt.Printf("v is of type %T\n", v)
}

func main_constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", math.Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func main_numericConstants() {
	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
