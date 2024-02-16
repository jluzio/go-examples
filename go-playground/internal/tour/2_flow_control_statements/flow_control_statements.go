package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// main_for()
	// main_forCont()
	// main_while()
	// main_if()
	// main_ifElse_shortStatement()
	// main_exerciseFunctions()
	// main_switch1()
	// main_switch2()
	main_defer()
}

func main_for() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func main_forCont() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func main_while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func main_forever() {
	for {
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main_if() {
	fmt.Println(sqrt(2), sqrt(-4))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
func main_ifElse_shortStatement() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func Sqrt(x float64) float64 {
	z := float64(1)
	bestZ := z
	bestX := math.Pow(z, 2)
	for i := 1; i <= 10; i++ {
		var currentX = math.Pow(z, 2)
		if math.Abs(x-currentX) < math.Abs(x-bestX) {
			bestZ = z
			bestX = currentX
		}
		z -= (z*z - x) / (2 * z)
	}
	return bestZ
}

func main_exerciseFunctions() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

func main_switch1() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func main_switch2() {
	t := time.Now()
	// only matches one case
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 12:
		fmt.Println("Good morning! (will never match)")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	case t.Hour() < 17:
		fmt.Println("Good afternoon (will never match).")
	default:
		fmt.Println("Good evening.")
	}
}

func main_defer() {
	// it's like try/finally in Java

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	defer fmt.Println("world")
	fmt.Println("hello")
}
