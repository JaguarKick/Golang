package main

import (
	"fmt"
	"pkcs11"
	"os"
)

var p11Func = pkcs11.New("")
var session pkcs11.SessionHandle

// This function would load the Luna library.
func loadLunaLibrary() {
	p11Path:=os.Getenv("ChrystokiConfigurationPath")
	p11Func = pkcs11.New(p11Path+"\\cklog201.dll")
}

// This function checks for the result.
func checkResult(ret error, message string) {
	if(ret!=nil) {
		fmt.Printf("Problem occured during %s : %s",message,ret)
		os.Exit(0)
	}
}

// This function initiates a new logged-in session.
func connectToSlot() {
	p11Func.Initialize()
	slot,ret:=p11Func.GetSlotList(true)
	checkResult(ret,"GetSlotList")
	
	session, ret=p11Func.OpenSession(slot[0],pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
	checkResult(ret,"OpenSession")
	
	ret=p11Func.Login(session,pkcs11.CKU_USER, "M@trixc0de")
	checkResult(ret,"Login")
	
	fmt.Println("Connected to slot %d via session %d",slot[0],session)
}

// This function logs out and closes the session.
func disconnectFromSlot() {
	ret:=p11Func.Logout(session)
	checkResult(ret,"Logout")
	
	ret=p11Func.CloseSession(session)
	checkResult(ret,"CloseSession")
	
	p11Func.Finalize()
	checkResult(ret,"Finalize")
	fmt.Println("Disconnected from P11 slot.")
}

func main() {

	fmt.Println("A simple connect disconnect demo.")
	loadLunaLibrary()
	connectToSlot()
	disconnectFromSlot()
	
}