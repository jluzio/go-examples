package main

import "fmt"

func main() {
	main_generics()
	main_generic_types()
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main_generics() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

type Holder[T any] struct {
	val T
}

func print_details[T any](holder Holder[T]) {
	fmt.Printf("holder = %v|%T || val = %v|%T\n", holder, holder, holder.val, holder.val)
}

func main_generic_types() {
	// List represents a singly-linked list that holds
	// values of any type.
	var holderInt = Holder[int]{10}
	var holderString = Holder[string]{"test"}

	print_details(holderInt)
	print_details(holderString)
}
