package main

import (
	"fmt"
)

func addingFive(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number, "+ 5 =", number+5)
	}
}
func main() {
	fmt.Println("Variadic functions in Go are when you can have any amount of trailing arguments")

	addingFive(1, 2, 3)
	addingFive(1)
	array_to_add := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	addingFive(array_to_add...)

}
