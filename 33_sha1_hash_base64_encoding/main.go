package main

import (
	"crypto/sha1" //several hash functions in various crypto/* packages.
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("SHA1 hash example")
	str := "Sha1"
	//First step
	h := sha1.New() //Here we start with a new hash.
	//Second step
	h.Write([]byte(str)) //Write expects bytes.
	//Third step
	hash := h.Sum(nil) //This gets the finalized hash result as a byte slice
	fmt.Println("Byte slice of hash: ", hash)
	fmt.Printf("String: %s -> Hash: %x\n", str, hash)

	fmt.Println()

	fmt.Println("Base64 encoding exampe")
	data := "keep it safe"

	sEncode := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Data: ", data, "Base64 encode: ", sEncode)

	sDecode, _ := base64.StdEncoding.DecodeString(sEncode)
	fmt.Println("Base64 encode: ", sEncode, "Data: ", string(sDecode))

}
