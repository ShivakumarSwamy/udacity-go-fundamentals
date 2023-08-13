package main

import "fmt"

func main() {
	// var product string = "T-Shirt"
	// var product = "T-Shirt"
	// var cost = 220

	product := "T-Shirt"
	cost := 220

	fmt.Println("product's value is:", product)
	fmt.Printf("product's type is: %T\n", product)

	cost = 15

	fmt.Println("cost's value is:", cost)
	fmt.Printf("cost's type is: %T\n", cost)

	// datatype
	// bool: true or false
	// string
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64, uintptr
	// byte or uint8
	// rune or int32
	// float32 or float64
	// complex64 and complex128
}
