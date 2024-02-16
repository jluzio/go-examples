package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// main_error()
	exercise_errors()
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
func run() error {
	return &MyError{time.Now(), "it didn't work"}
}
func main_error() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("value %f is not valid", e)
}

func Sqrt(x float64) (float64, error) {
	r := math.Sqrt(x)
	switch {
	case math.IsNaN(r):
		return r, ErrNegativeSqrt(x)
	default:
		return r, nil
	}
}

func SqrtV0(x float64) (float64, error) {
	r := math.Sqrt(x)
	switch {
	case math.IsNaN(r):
		return r, fmt.Errorf("value %v is not valid", x)
	default:
		return r, nil
	}
}
func exercise_errors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
