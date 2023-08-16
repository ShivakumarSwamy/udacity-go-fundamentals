package main

import "fmt"

type Student struct {
	id   uint
	name string
}

type Classroom struct {
	id          uint
	capacity    uint
	subject     string
	studentList []Student
}

func main() {
	student1 := Student{id: 1, name: "Foo"}
	student2 := Student{id: 2, name: "Bar"}

	c1 := Classroom{id: 1, capacity: 2, subject: "Biology", studentList: []Student{student1, student2}}
	fmt.Println(c1)

	c2 := new(Classroom)
	c2.id = 2
	c2.capacity = 1
	c2.subject = "Maths"
	c2.studentList = []Student{student1}

	fmt.Println(c2)
}
