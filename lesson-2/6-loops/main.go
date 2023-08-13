package main

import "fmt"

func main() {

	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	x := 0 // initialization

	for x <= 20 { // condition
		fmt.Println(x)
		x += 2 // increment
	}
}
