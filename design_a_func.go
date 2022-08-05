package main

import "fmt"

func main() {
	listNums := []float64{1, 2, 3, 4, 5}

	fmt.Println("Sum :", addThemUp(listNums))
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}

	return sum
}
