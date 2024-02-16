package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func main() {
	// main_pointers()
	// main_structs()
	// main_structs_pointers()
	// main_struct_literals()
	// main_arrays()
	// main_slices()
	// main_slices_are_references_or_views()
	// main_slice_literals()
	// main_slice_defaults()
	// main_slice_len_capacity()
	// main_slice_nil()
	// main_slice_builder()
	// main_slice_of_slices()
	// main_slice_append()
	// main_for_range()
	// main_for_range_2()
	// exercise_slices()
	// main_map()
	// main_map_literal()
	// main_map_mutate()
	// exercise_maps()
	// main_function_as_value()
	// main_function_closure()
	exercise_functions()
}

func main_pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type VertexInt struct {
	X int
	Y int
}

func main_structs() {
	v := VertexInt{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
}

func main_structs_pointers() {
	v := VertexInt{1, 2}

	v2 := v

	p := &v
	p.X = 1e9

	v2.Y = 3

	fmt.Println(v)
	fmt.Println(v2)
}

func main_struct_literals() {
	var (
		v1 = VertexInt{1, 2}  // has type Vertex
		v2 = VertexInt{X: 1}  // Y:0 is implicit
		v3 = VertexInt{}      // X:0 and Y:0
		p  = &VertexInt{1, 2} // has type *Vertex
	)
	fmt.Println(v1, p, v2, v3)
}

func main_arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// array then a slice literal
	var b = []string{"Hello", "World"}
	fmt.Println(b)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func main_slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func main_slices_are_references_or_views() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func main_slice_literals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func main_slice_defaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func main_slice_len_capacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main_slice_nil() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func main_slice_builder() {
	a := make([]int, 5)
	printSliceWithMake("a", a)

	b := make([]int, 0, 5)
	printSliceWithMake("b", b)

	c := b[:2]
	printSliceWithMake("c", c)

	d := c[2:5]
	printSliceWithMake("d", d)
}
func printSliceWithMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main_slice_of_slices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func main_slice_append() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func main_for_range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func main_for_range_2() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx, dy)
	bitmap := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		bitmap[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			bitmap[y][x] = uint8((x + y) / 2)
		}
	}

	fmt.Println(bitmap)
	return bitmap
}

// check in browser with "data:image/png;base64," + <result without IMAGE:>
func exercise_slices() {
	pic.Show(Pic)
}

type VertexFloat struct {
	Lat, Long float64
}

var m map[string]VertexFloat

func main_map() {
	m = make(map[string]VertexFloat)
	m["Bell Labs"] = VertexFloat{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func main_map_literal() {
	var m = map[string]VertexFloat{
		"Bell Labs": VertexFloat{40.68433, -74.39967},
		"Google":    VertexFloat{37.42202, -122.08408},
	}
	fmt.Println(m)

	var m2 = map[string]VertexFloat{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)
}

func main_map_mutate() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	var wordMap = make(map[string]int)
	var words = strings.Split(s, " ")
	for _, word := range words {
		wordMap[word]++
	}
	return wordMap
}
func exercise_maps() {
	wc.Test(WordCount)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func main_function_as_value() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main_function_closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	numbers := []int{0, 1}
	i := 0
	return func() int {
		var v, x, y int
		if i <= 1 {
			v = numbers[i]
		} else {
			x, y = numbers[i-2], numbers[i-1]
			v = x + y
			numbers = append(numbers, v)
		}
		i++
		return v
	}
}
func exercise_functions() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
