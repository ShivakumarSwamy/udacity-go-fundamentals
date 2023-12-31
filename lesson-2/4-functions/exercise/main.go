package main

import "fmt"

func getRectangleArea(width, length int) string {
	area := width * length
	if area < 50 {
		return fmt.Sprintf("The area is %d, which is less than 50", area)
	}
	return fmt.Sprintf("The area is %d, which is greater than or equal to 50", area)
}

func main() {
	fmt.Println(getRectangleArea(5, 2))
	fmt.Println(getRectangleArea(5, 20))
}
