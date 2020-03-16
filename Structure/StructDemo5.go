package main

import "fmt"

type scores struct {
	english int
	hindi int
	math int
	science int
	sstudies int
}

type student struct {
	rollNo int
	fName string
	lName string
	result scores
}

func main() {
	var tbeetle student
	fmt.Println(tbeetle)
}
