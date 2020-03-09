package main

import "fmt"

func calc(a,b,c,d,e float64) (float64, float64) {
	total:=a+b+c+d+e
	avg:=total/5
	return total,avg
}

func main() {
	eng:=56.0
	hin:=45.0
	math:=46.0
	sci:=63.0
	sst:=56.0
	total, avg := calc(eng, hin, math, sci, sst)
	fmt.Printf("Total	: %f\n",total)
	fmt.Printf("Average	: %f\n",avg)
}

