package main

import "fmt"

type Baseball struct {
	mass         int
	acceleration int
}

type Football struct {
}

type Force interface {
	getForce() int
}

func (b Baseball) getForce() int {
	return b.mass * b.acceleration
}

func (f Football) getForce() int {
	return 50
}

func compareForce(a, b Force) bool {
	return a.getForce() > b.getForce()
}

func main() {
	b1 := Baseball{mass: 2, acceleration: 20}
	f1 := Football{}

	fmt.Println(compareForce(b1, f1))
}
