package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println("i =", i)
		i++ // i = i +1
	}
	fmt.Println("===============")

	for j := 0; j < 5; j++ {
		fmt.Println("j =", j)
	}
}
