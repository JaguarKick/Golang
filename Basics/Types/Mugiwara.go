package main

import "fmt"

type mugiwara []string

func (m mugiwara) print() {
	for i, crew := range m {
		fmt.Println(i, crew)
	}
}

