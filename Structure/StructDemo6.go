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
	tbeetle := student{
		rollNo: 12,
		fName : "Tiger",
		lName : "Beetle",
		result : scores{
			english:78,
			hindi:76,
			math:89,
			science:81,
			sstudies:81,
		},
	}
	
	fmt.Println(tbeetle)
}
