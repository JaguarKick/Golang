package main

import "fmt"

func main() {
	fmt.Println(greet())
	fmt.Println(greet())
	fmt.Println(greet())
	fmt.Println(greet())
}

func greet() string {
	return "Hello World"
}
