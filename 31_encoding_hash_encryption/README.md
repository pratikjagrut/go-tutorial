# Encoding, Hashing and Encryption

To protect the data from wrongful use we tend to perform cryptography on that data and make obfuscating.

Golang provides a rich library for encoding, hashing and encrypting our data.

## Encoding

Go provides built-in support for `base64/base32` encoding/decoding.

```go
import (
    "encoding/base32"
    "encoding/base64"
)
```
Here we'll encode/decode data using `base64/base32` methods.

Here’s the string we’ll encode/decode.

```go
data := "Super Secret message"
```
Go supports both `standard` and `URL-compatible` base64 encoding/decoding. 

***This is standard base64 encoding/decoding.***

```go
sEncode := base64.StdEncoding.EncodeToString([]byte(data))
fmt.Printf("Base64 std encoding:  \t%v\n", sEncode)

sDecode, _ := base64.StdEncoding.DecodeString(sEncode)
fmt.Printf("Base64 std decode:  \t%v\n", string(sDecode))
```

Output:
```
Base64 std encoding:    U3VwZXIgU2VjcmV0IG1lc3NhZ2U=
Base64 std decode:      Super Secret message
```

***This is URL-compatible base64 encoding/decoding.***

```go
sEncode := base64.URLEncoding.EncodeToString([]byte(data))
fmt.Printf("Base64 url encode:  \t%v\n", sEncode)

sDecode, _ := base64.URLEncoding.DecodeString(sEncode)
fmt.Printf("Base64 url decode:  \t%v\n", string(sDecode))
```

Output:

```
Base64 URL encode:      U3VwZXIgU2VjcmV0IG1lc3NhZ2U=
Base64 URL decode:      Super Secret message
```

***Go supports `standard` base32 encoding/decoding.***

```go
sEncode = base32.StdEncoding.EncodeToString([]byte(data))
fmt.Printf("Base32 url encode:  \t%v\n", sEncode)

sDecode, _ = base32.StdEncoding.DecodeString(sEncode)
fmt.Printf("Base32 std decode:  \t%v\n", string(sDecode))
```

Output:

```
Base32 URL encode:      KN2XAZLSEBJWKY3SMV2CA3LFONZWCZ3F
Base32 std decode:      Super Secret message
```

## Hashing

There are several different types of hash functions are provided in the `crypto` package.

We will compute the hash for following data using SHA1 and MD5 algorithms.

```go
data := "Super Secret message"
```

### SHA1

`SHA1` hashes are frequently used to compute short identities for binary or text blobs. For example, the `git version control system` uses SHA1s extensively to identify versioned files and directories.

***Steps to compute hash:***

Import

```go
import "crypto/sha1"
```

Start with a new hash

```go
h := sha1.New()
```

Next step is to write byte stream using 

`Write(p []byte) (n int, err error)`

Write writes `len(p)` bytes from p to the underlying data stream.
It `returns the number of bytes written` from p (0 <= n <= len(p)) and `any error encountered` that caused the write to stop early.
Write must return a non-nil error if it returns n < len(p).

```go
h.Write([]byte(data))
```

Get the final has from 

`Sum(b []byte) []byte`

```go
hash := h.Sum(nil)
```

Use the `%x` format verb to convert a `hash results to a hex string`.

```go
fmt.Printf("[]byte:  \t%v\n", hash) //Hash byte slice
fmt.Printf("Hash:   \t%x\n", hash)
```

Output:

```
[]byte:     [64 210 95 11 8 99 108 69 206 246 165 112 41 75 244 123 171 42 85 193]
Hash:       40d25f0b08636c45cef6a570294bf47bab2a55c1
```

### MD5 Hash

The MD5 message-digest algorithm is a widely used hash function producing a 128-bit hash value.

***Steps to compute hash:***

Steps are very similar to SHA1 hash.

Import

```go
import "crypto/md5"
```

```go
fmt.Println("MD5 hash example")
md5 := md5.New()
md5.Write([]byte(data))
md5hash := md5.Sum(nil)
fmt.Printf("[]byte:  \t%v\n", md5hash)
fmt.Printf("Hash:   \t%x\n", md5hash)
```

Output:

```
[]byte:     [206 48 10 227 65 165 33 90 15 171 225 201 191 97 58 73]
Hash:       ce300ae341a5215a0fabe1c9bf613a49]
```

## Encryption and Decryption

`Encryption` is the process of translating plain text data (plaintext) into something that appears to be random and meaningless (ciphertext). `Decryption` is the process of converting ciphertext back to plaintext.

The `crypto` package provides strong encryption and decryption support.

### AES

The `Advanced Encryption Standard`, also known by its original name Rijndael, is a specification for the encryption of electronic data established by the U.S. National Institute of Standards and Technology in 2001

***Steps for encryption:***

Import

```go
import "crypto/aes"
import "crypto/cipher"
```

We will encrypt the following data.

```go
data := "Super Secret message"
```

Convert data to a byte slice

```go
text := []byte(data)
```

The key size should be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.

```go
key := []byte("16BytePassphrase")
```

Generate a new AES cipher using our 32-byte long key

```go
c, err := aes.NewCipher(key)
if err != nil {
    fmt.Println(err)
}
```

gcm or Galois/Counter Mode is a mode of operation
for symmetric key cryptographic block ciphers
https://en.wikipedia.org/wiki/Galois/Counter_Mode

```go
gcm, err := cipher.NewGCM(c)
if err != nil {
    fmt.Println(err)
}
```

create a new byte array the size of the nonce which must be passed to Seal and populates our nonce with a cryptographically secure random sequence.

```go
nonce := make([]byte, gcm.NonceSize())
if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
    fmt.Println(err)
}
```

Here we encrypt our text using the `Seal function`. Seal encrypts and authenticates plaintext, authenticates the additional data and appends the result to DST, returning the updated slice. The nonce must be NonceSize() bytes long and unique for all time, for a given key.

Seal returns a byte array of encryption string.

```go
ciphertext := gcm.Seal(nonce, nonce, text, nil)
fmt.Printf("Encrypted Text: \t%v\n", string(ciphertext))
```

***Steps for decryption:***

Get the Nonce Size
```go
nonceSize := gcm.NonceSize()
if len(ciphertext) < nonceSize {
    fmt.Println(err)
}
```

Decrypt ciphertext

```go
nonce, ciphertext = ciphertext[:nonceSize], ciphertext[nonceSize:]
plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("Decrypted Text: \t%v\n", string(plaintext))
```

Output:

```
Encrypted Text:     ��y^�^̫BIw�����h��v�-��Ђ��4�����_4LU\y9\�
Decrypted Text:     Super Secret message
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```