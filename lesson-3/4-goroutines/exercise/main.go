package main

import (
	"fmt"
	"time"
)

func printSlice(data []string) {

	for _, element := range data {
		fmt.Println(element)
	}
}

func main() {
	startTime := time.Now()
	colorNames := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	colorCodes := []string{"#FF0000", "#FF7F00", "#FFFF00", "#00FF00", "#0000FF", "#4B0082", "#9400D3"}

	go printSlice(colorNames)
	printSlice(colorCodes)

	duration := time.Since(startTime)

	fmt.Println("Duration of execution: " + duration.String())

	time.Sleep(time.Second)
}
