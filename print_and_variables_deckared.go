// Runing this file, by runing this command in terminal
// go run herewego.go
package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	// uint8 : unsigned 8-bit integers (0 to 255)
	// uint16 : unsigned 8-bit integers (0 to 65535)
	// uint32 : unsigned 8-bit integers (0 to 4294967295)
	// uint64 : unsigned 8-bit integers (0 to 18446744073709551615)
	// int8 : signed 8-bit integers (-128 to 127)
	// int16 : signed 8-bit integers (-32768 to 32767)
	// int32 : signed 8-bit integers (-2147483648 to 2147483647)
	// int64 : signed 8-bit integers (-9223372036854775808 to 9223372036854775807)

	var age int = 40
	var favNum float64 = 1.6180339
	randNum := 1
	fmt.Println(age, " ", favNum) // 40   1.6180339
	fmt.Println(randNum)          // 1

	var numOne = 1.000
	var num99 = .9999

	fmt.Println(numOne - num99) // 9.999999999998899e-05

	fmt.Println("6 + 4 =", 6+4)
	fmt.Println("6 - 4 =", 6-4)
	fmt.Println("6 * 4 =", 6*4)
	fmt.Println("6 / 4 =", 6/4)
	fmt.Println("6 % 4 =", 6%4)
}
