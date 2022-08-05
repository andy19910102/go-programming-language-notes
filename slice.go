package main

import "fmt"

func main() {
	numSlice := []int{5, 4, 3, 2, 1}
	numSlice2 := numSlice[3:5]
	fmt.Println("numSlice2[0] = ", numSlice2[0])
	fmt.Println("numSlice2[1] = ", numSlice2[1])
	// fmt.Println("numSlice2[2] = ", numSlice2[2])
	// panic: runtime error: index out of range [2] with length 2

	fmt.Println("numSlice[:2] =", numSlice[:2]) // [5, 4]
	fmt.Println("numSlice[2:] =", numSlice[2:]) // [3 2 1]

	numSlice3 := make([]int, 5, 10)
	copy(numSlice3, numSlice)
	fmt.Println("numSlice3=", numSlice3)

	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println("numSlice3=", numSlice3)
}
