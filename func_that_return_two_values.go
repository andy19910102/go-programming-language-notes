package main

import "fmt"

func main() {
	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)
}

func next2Values(number int) (int, int) {
	return number + 1, number + 2
}
