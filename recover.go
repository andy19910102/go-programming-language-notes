package main

import "fmt"

func main() {
	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))
}

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}
