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

func (s student) print() {
	fmt.Printf("First name		: %s\n",s.fName)
	fmt.Printf("Last name		: %s\n",s.lName)
	fmt.Printf("Roll No		: %d\n",s.rollNo)
}

func (s scores) print() {
	fmt.Printf("English		: %d\n",s.english)
	fmt.Printf("Hindi		: %d\n",s.hindi)
	fmt.Printf("Math		: %d\n",s.math)
	fmt.Printf("Science		: %d\n",s.science)
	fmt.Printf("S. Studies		: %d\n",s.sstudies)
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
	
	tbeetle.print()
	tbeetle.result.print()
}
