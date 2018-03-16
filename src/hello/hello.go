package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

func main() {
	fmt.Println("hello, world")
	fmt.Println("My favorite number is: ", rand.Intn(10))
	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	a, b := swap("hello", "go")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var c, python, java = true, false, "no!"
	k := 3
	fmt.Println(i, j, c, python, java, k)

	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	fmt.Println(x, y, uint(f))
}
