package main

import (
	"fmt"
	"os"
	"github.com/miekg/pkcs11"
)

var p11Func = pkcs11.New("")
var session pkcs11.SessionHandle
var priKey pkcs11.ObjectHandle
var pubKey pkcs11.ObjectHandle
var ret pkcs11.Error
var params *pkcs11.OAEPParams

// This function would load the Luna library.
func loadLunaLibrary() {
	p11Path := os.Getenv("LUNA_LIB")
	if p11Path != "" {
		p11Func = pkcs11.New(p11Path + "/libCryptoki2_64.so")
	} else {
		fmt.Println("Failed to find Luna library.")
		fmt.Println("Make sure you've set LUNA_LIB envvar pointing to the directory containing the luna library.")
		os.Exit(0)
	}
	fmt.Println("* Luna library loaded.")
}

// This function checks for the result.
func checkResult(ret error, message string) {
	if ret != nil {
		fmt.Printf("Problem occured during %s : %s\n", message, ret)
		p11Func.Finalize()
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

	ret = p11Func.Login(session, pkcs11.CKU_USER, "Partition_Password")
	checkResult(ret, "Login")
	fmt.Println("* Connected.")
}

// This function generates RSA-2048 keypair.
func generateRSAKeyPair() (pkcs11.ObjectHandle, pkcs11.ObjectHandle) {
	var modBits = 2048
	exp := []byte{0x01, 0x00, 0x00, 0x00, 0x01}
	publicTemplate := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN, false),
		pkcs11.NewAttribute(pkcs11.CKA_PRIVATE, true),
		pkcs11.NewAttribute(pkcs11.CKA_ENCRYPT, true),
		pkcs11.NewAttribute(pkcs11.CKA_VERIFY, true),
		pkcs11.NewAttribute(pkcs11.CKA_WRAP, true),
		pkcs11.NewAttribute(pkcs11.CKA_MODULUS_BITS, modBits),
		pkcs11.NewAttribute(pkcs11.CKA_PUBLIC_EXPONENT, exp),
	}

	privateTemplate := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN, false),
		pkcs11.NewAttribute(pkcs11.CKA_PRIVATE, true),
		pkcs11.NewAttribute(pkcs11.CKA_DECRYPT, true),
		pkcs11.NewAttribute(pkcs11.CKA_SIGN, true),
		pkcs11.NewAttribute(pkcs11.CKA_UNWRAP, true),
		pkcs11.NewAttribute(pkcs11.CKA_EXTRACTABLE, false),
		pkcs11.NewAttribute(pkcs11.CKA_MODIFIABLE, false),
		pkcs11.NewAttribute(pkcs11.CKA_SENSITIVE, true),
	}
	pubKey, priKey, ret := p11Func.GenerateKeyPair(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_RSA_PKCS_KEY_PAIR_GEN, nil)}, publicTemplate, privateTemplate)
	checkResult(ret, "GenerateKeyPair")
	fmt.Println("* RSA-2048 keypair generated.")
	fmt.Printf("   - Private Key : %d\n", priKey)
	fmt.Printf("   - Public Key  : %d\n", pubKey)
	return pubKey, priKey
}

// This function signs the raw data.
func signData(rawData []byte) []byte {

	ret := p11Func.SignInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_SHA256_RSA_PKCS, nil)}, priKey)
	checkResult(ret, "C_SignInit")

	signature, ret := p11Func.Sign(session, rawData)
	checkResult(ret, "C_Sign")
	fmt.Println("* Data signed.")

	return signature
}

// This function verifies the signature of decrypted data.
func verifyData(rawData []byte, signature []byte) {
	ret := p11Func.VerifyInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_SHA256_RSA_PKCS, nil)}, pubKey)
	checkResult(ret, "C_VerifyInit")

	ret = p11Func.Verify(session, rawData, signature)
	checkResult(ret, "C_Verify")
	fmt.Println("* Data verified.")
}

// This function encrypts a raw data using CKM_RSA_PKCS_OAEP mechanism.
func encryptData(rawData []byte) []byte {
	ret := p11Func.EncryptInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_RSA_PKCS_OAEP, params)}, pubKey)
	checkResult(ret, "C_EncryptInit")

	encrypted, ret := p11Func.Encrypt(session, rawData)
	checkResult(ret, "C_Encrypt")
	fmt.Println("* Data Encrypted.")
	return encrypted
}

// This function decrypts an encrypted data using CKM_RSA_PKCS_OAEP mechanism.
func decryptData(encrypted []byte) []byte {
	ret := p11Func.DecryptInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_RSA_PKCS_OAEP, params)}, priKey)
	checkResult(ret, "C_DecryptInit")

	decrypted, ret := p11Func.Decrypt(session, encrypted)
	checkResult(ret, "C_Decrypt")
	fmt.Println("* Data Decrypted.")
	return decrypted
}

// This function logs out and closes the session.
func disconnectFromSlot() {
	ret := p11Func.Logout(session)
	checkResult(ret, "Logout")

	ret = p11Func.CloseSession(session)
	checkResult(ret, "CloseSession")

	p11Func.Finalize()
	checkResult(ret, "Finalize")
	fmt.Println("* Disconnected.")
}

func main() {
	rawData := []byte("This is a simple text")
	fmt.Println("PKCS #11 Golang Sample")
	fmt.Println("-----------------------")
	fmt.Println()
	fmt.Println("This sample code performs the following :-")
	fmt.Println("  1. Signs a raw data.")
	fmt.Println("  2. Encrypts that raw data.")
	fmt.Println("  3. Decrypts the encrypted data.")
	fmt.Println("  4. Verifies the signature of decrypted data.")
	fmt.Println()
	fmt.Println()
	params = pkcs11.NewOAEPParams(pkcs11.CKM_SHA256, pkcs11.CKG_MGF1_SHA256, pkcs11.CKZ_DATA_SPECIFIED, []byte("Hello World"))
	loadLunaLibrary()
	connectToSlot()
	pubKey, priKey = generateRSAKeyPair()
	signature := signData(rawData)
	//initOAEPParams()
	encryptedData := encryptData(rawData)
	decryptedData := decryptData(encryptedData)
	verifyData(decryptedData, signature)
	disconnectFromSlot()
	fmt.Println("-- THE END --")
}