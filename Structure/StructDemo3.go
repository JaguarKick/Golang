package main

import "fmt"

type student struct {
	rollNo int
	fName string
	lName string
}

func main() {
	var tbeetle student
	tbeetle.rollNo = 71
	tbeetle.fName = "Tiger"
	tbeetle.lName = "Beetle"
	fmt.Printf("Student Name	: %s %s\n",tbeetle.fName,tbeetle.lName)
	fmt.Printf("Roll No		: %d\n",tbeetle.rollNo)
}
