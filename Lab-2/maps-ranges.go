package main

import (
	"fmt"
)

func main() {
	new_map := make(map[string]int)

	new_map["Alpha"] = 1
	new_map["Omega"] = 999

	fmt.Println("The map is:", new_map)

	value1 := new_map["Alpha"]
	fmt.Println(value1)

	new_map["Sigma"] = 5
	fmt.Println(new_map)

	delete(new_map, "Sigma")
	fmt.Println(new_map)

	var is_there bool

	_, is_there = new_map["Sigma"]

	fmt.Println(is_there)

	var num int
	var has_value bool

	num, has_value = new_map["Omega"]

	fmt.Println(has_value, num)

	//RANGES

	number_array := []int{2, 3, 4}
	sum := 0
	for _, num := range number_array {
		sum += num
	}
	fmt.Println("The sum of the numbers is:", sum)
}
