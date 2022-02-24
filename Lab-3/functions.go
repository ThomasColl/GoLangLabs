package main

import (
	"fmt"
)

func plusFive(number int) int {
	return number + 5
}
func minusFive(number int) int {
	return number - 5
}

func plusAndMinusFive(number int) (int, int) {
	return plusFive(number), minusFive(number)
}

func main() {
	fmt.Println("Functions in Go are interesting because you can return single or multiple values")

	fmt.Println(plusFive(10))
	fmt.Println(minusFive(7))
	fmt.Println(plusAndMinusFive(6))
}
