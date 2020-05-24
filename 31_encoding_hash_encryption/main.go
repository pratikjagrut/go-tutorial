package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"io"
)

//several hash functions in various crypto/* packages.

func main() {
	data := "Super Secret message" //this data will be hashed and encode by various algorithms

	fmt.Println("Base64 encoding example")
	//base64 ecoding and decoding
	sEncode := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("Base64 std encoding:  \t%v\n", sEncode)

	sDecode, _ := base64.StdEncoding.DecodeString(sEncode)
	fmt.Printf("Base64 std decode:  \t%v\n", string(sDecode))

	fmt.Println()

	sEncode = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Printf("Base64 url encode:  \t%v\n", sEncode)

	sDecode, _ = base64.URLEncoding.DecodeString(sEncode)
	fmt.Printf("Base64 url decode:  \t%v\n", string(sDecode))

	fmt.Println()

	fmt.Println("Base32 encoding example")
	//base32 ecoding and decoding
	sEncode = base32.StdEncoding.EncodeToString([]byte(data))
	fmt.Printf("Base32 url encode:  \t%v\n", sEncode)

	sDecode, _ = base32.StdEncoding.DecodeString(sEncode)
	fmt.Printf("Base32 std decode:  \t%v\n", string(sDecode))

	fmt.Println()

	fmt.Println("SHA1 hash example")
	//First step
	h := sha1.New() //Here we start with a new hash.
	//Second step
	h.Write([]byte(data)) //Write expects bytes.
	//Third step
	hash := h.Sum(nil) //This gets the finalized hash result as a byte slice
	fmt.Printf("[]byte:  \t%v\n", hash)
	fmt.Printf("Hash:   \t%x\n", hash)

	fmt.Println()

	fmt.Println("MD5 hash example")
	md5 := md5.New()
	md5.Write([]byte(data))
	md5hash := md5.Sum(nil)
	fmt.Printf("[]byte:  \t%v\n", md5hash)
	fmt.Printf("Hash:   \t%x\n", md5hash)

	fmt.Println()

	fmt.Println("Encryption and decryption")
	//Encryption
	text := []byte(data)
	// The key size should be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
	key := []byte("16BytePassphrase")
	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}
	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		fmt.Println(err)
	}
	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	ciphertext := gcm.Seal(nonce, nonce, text, nil)
	fmt.Printf("Encrypted Text: \t%v\n", string(ciphertext))

	//Decryption
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext = ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Decrypted Text: \t%v\n", string(plaintext))
}
