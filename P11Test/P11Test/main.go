package main

// Required packages
import (
	"fmt" // you know what this package is required for..
	"os" // required for reading environment variables.
	"log" // logging
	"pkcs11" // pkcs11 wrapper for go.
)

func main() {

	fmt.Println("P11Test")
	fmt.Println("--------")
	fmt.Println("This sample application will do the following :-")
	fmt.Println("\tC_Initialize > C_OpenSession > C_Login")
	fmt.Println("\tC_Logout > C_CloseSession > C_Finalize")

	p11Lib:=os.Getenv("P11_LIB") // P11_LIB is the path where Luna Client library is located.
	if(p11Lib!="") {
		fmt.Println("Luna Client installation found.")
	} else {
		log.Fatal("Failed to find Luna Client.")
	}
	
	isCklogEnabled:=os.Getenv("CKLOG") // This is optional..
	if(isCklogEnabled=="YES") {
		fmt.Println("Cklog will be enabled.\n")
		p11Lib=p11Lib+"/libcklog2.so"
	} else {
		p11Lib=p11Lib+"/libCryptoki2_64.so"
	}
	fmt.Println("P11Lib	: ", p11Lib)
	
	p11Func:=pkcs11.New(p11Lib)
	ret:=p11Func.Initialize()
	if(ret!=nil) {
		panic(ret)
	}
	
	slots, ret := p11Func.GetSlotList(true)
	if(ret!=nil) {
		panic(ret)
	}
	
	session, ret := p11Func.OpenSession(slots[0],pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
	if(ret!=nil) {
		panic(ret)
	}
	fmt.Printf("New session %d open on slot %d\n",session, slots[0])

	ret = p11Func.Login(session, pkcs11.CKU_USER, "M@trixc0de")
	if(ret!=nil) {
		panic(ret)
	}

	ret = p11Func.Logout(session)
	if(ret!=nil) {
		panic(ret)
	}

	ret = p11Func.CloseSession(session)
	if(ret!=nil) {
		panic(ret)
	}

	ret = p11Func.Finalize()
	if(ret!=nil) {
		panic(ret)
	}
	fmt.Println("Finalized...This app will now exit.")
}
