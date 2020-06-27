package main

import (
	"fmt"
	"os"
	"pkcs11"
)

var p11Func = pkcs11.New("")
var session pkcs11.SessionHandle
var objHandle pkcs11.ObjectHandle
var dataToSign = "Hello World"

// This function would load the Luna library.
func loadLunaLibrary() {
	p11Path := os.Getenv("ChrystokiConfigurationPath")
	p11Func = pkcs11.New(p11Path + "\\cklog201.dll")
}

// This function checks for the result.
func checkResult(ret error, message string) {
	if ret != nil {
		fmt.Printf("Problem occured during %s : %s", message, ret)
		os.Exit(0)
	}
}

// This function initiates a new logged-in session.
func connectToSlot() {
	p11Func.Initialize()
	slot, ret := p11Func.GetSlotList(true)
	checkResult(ret, "GetSlotList")

	session, ret = p11Func.OpenSession(slot[0], pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
	checkResult(ret, "OpenSession")

	ret = p11Func.Login(session, pkcs11.CKU_USER, "M@trixc0de")
	checkResult(ret, "Login")

	fmt.Println("Connected to slot %d via session %d", slot[0], session)
}

// This function logs out and closes the session.
func disconnectFromSlot() {
	ret := p11Func.Logout(session)
	checkResult(ret, "Logout")

	ret = p11Func.CloseSession(session)
	checkResult(ret, "CloseSession")

	p11Func.Finalize()
	checkResult(ret, "Finalize")
	fmt.Println("Disconnected from P11 slot.")
}

func generateAESKey() {
	var ret error
	des3KeyAttributes := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN, false),
		pkcs11.NewAttribute(pkcs11.CKA_ENCRYPT, true),
		pkcs11.NewAttribute(pkcs11.CKA_DECRYPT, true),
		pkcs11.NewAttribute(pkcs11.CKA_SENSITIVE, true),
		pkcs11.NewAttribute(pkcs11.CKA_PRIVATE, true),
		pkcs11.NewAttribute(pkcs11.CKA_VALUE_LEN, 32),
	}

	objHandle, ret = p11Func.GenerateKey(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_AES_KEY_GEN, nil)}, des3KeyAttributes)
	checkResult(ret, "GenerateKey")
	fmt.Printf("AES-256 key generated as handle %d\n", objHandle)
}

func signData() {
	var signature []byte
	p11Func.SignInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_AES_GMAC, nil)}, objHandle)
	signature = p11Func.Sign(session, dataToSign)
}

func main() {

	fmt.Println("DES-3 Key Demo.\n")
	loadLunaLibrary()
	connectToSlot()
	generateAESKey()
	signData()
	disconnectFromSlot()
}
