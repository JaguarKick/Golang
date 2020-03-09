package main

import "fmt"

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var num1 int = 34
	var num2 int = 45
	
	fmt.Println("Before swapping the variables");
	fmt.Printf("Num1 is %d\n",num1)
        fmt.Printf("Num2 is %d\n",num2)
	
	swap(&num1, &num2)
	
	fmt.Println("After swapping the variables")
	fmt.Printf("Num1 is %d\n",num1)
	fmt.Printf("Num2 is %d\n",num2)
}
