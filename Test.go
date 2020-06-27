package main

import (
	"fmt"
	"pkcs11"
)

func main() {
	p11 := pkcs11.New("C:\\Program Files\\SafeNet\\LunaClient\\cklog201.dll")
	fmt.Printf("Type : %T", p11)
}
