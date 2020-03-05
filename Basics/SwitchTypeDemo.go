package main

import "fmt"

func main() {
	var x interface{}
	switch i:=x.(type) {
		case nil: fmt.Printf("type of x : %T",i)
		case int: fmt.Printf("type of x : %T",i)
		case string:fmt.Printf("type of x : %T",i)
	}
}
