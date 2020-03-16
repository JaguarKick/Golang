package main

import (
	"fmt"
	"math/rand"
	)

func main() {
	rand.Seed(71)
	fmt.Printf("Random number is 	: %d\n",rand.Intn(12))
	fmt.Printf("Random number is 	: %d\n",rand.Intn(12))
	fmt.Printf("Random number is 	: %d\n",rand.Intn(12))
	fmt.Printf("Random number is 	: %d\n",rand.Intn(12))
}
