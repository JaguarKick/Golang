package main

import "fmt"

type student struct {
	rollNo int
	fName string
	lName string
}

func main() {
	var tbeetle student
	fmt.Printf("%+v\n",tbeetle)
}
