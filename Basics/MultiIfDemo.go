package main

import "fmt"

func main() {
	var num1, num2, num3 = 45,34,57
	
	if(num1 > num2) {
		if(num1 > num3) {
			fmt.Println("Num1 is the greatest.")
		} else {
			fmt.Println("Num 3 is the greatest.")
		}
	} else if(num2 > num3) {
		fmt.Println("Num2 is the greatest.")
	} else {
		fmt.Println("Num3 is the greatest.")
	}
}
