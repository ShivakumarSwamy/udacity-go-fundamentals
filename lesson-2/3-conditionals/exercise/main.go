package main

import "fmt"

func main() {
	number := -8
	numberCheck(number)
	number = 28
	numberCheck(number)
	number = 280000
	numberCheck(number)
}

func numberCheck(number int) {
	if number < 0 {
		fmt.Println(number, "is negative")
	} else if number > 0 && number < 100 {
		fmt.Println(number, "is positive")
	} else {
		fmt.Println(number, "is positive and is a large number!")
	}
}
