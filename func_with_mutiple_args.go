package main

import "fmt"

func main() {
	fmt.Println(subtractThem(1, 2, 3, 4, 5))
}

func subtractThem(args ...int) int {
	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}
