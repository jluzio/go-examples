package main

import (
	"fmt"
	"math"
)

func main() {
	// main_method_struct()
	// main_method_type()
	main_method_pointer()
}

type Vertex struct {
	X, Y float64
}

// method for type Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main_method_struct() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// type alias
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main_method_type() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// only with pointer can change instance
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main_method_pointer() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
