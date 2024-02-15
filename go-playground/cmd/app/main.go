package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}
