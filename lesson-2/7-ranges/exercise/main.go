package main

import "fmt"

func reduce(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(reduce([]int{0, 1, 1, 2, 3, 5, 8, 13, 21}))
}
