package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := map[uint]string{
		1: "Calculus",
		2: "Biology",
		3: "Chemistry",
		4: "Computer Science",
		5: "Communications",
		6: "English",
		7: "Cantonese",
	}

	for id, course := range courses {
		if strings.HasPrefix(course, "C") {
			fmt.Println(id, course)
		}
	}

	courses[4] = "Algorithms"
	courses[8] = "Spanish"
	fmt.Println(courses)

	delete(courses, 1)
	fmt.Println(courses)
}
