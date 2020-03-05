package main

import "fmt"

func main() {
	mugiwara:=[]string{"Luffy","Zoro","Sanji","Nami","Usoup","Franky"}

	fmt.Println(mugiwara)
	mugiwara = append(mugiwara,"Robin")
	mugiwara = append(mugiwara,"Tony")
	mugiwara = append(mugiwara,"Brookes")
	fmt.Println(mugiwara)
}
