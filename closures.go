package main

import "fmt"

func main() {
	num3 := 3

	doubleNum := func() int {
		num3 *= 2
		return num3
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
}
