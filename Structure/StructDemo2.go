package main

import "fmt"

type student struct {
	rollNo int
	fName string
	lName string
}

func main() {
	sam:= student{rollNo:34,fName:"Sam",lName:"Paul"}
	fmt.Printf("Student Name	: %s %s\n",sam.fName,sam.lName)
	fmt.Printf("Roll No		: %d\n",sam.rollNo)
}
