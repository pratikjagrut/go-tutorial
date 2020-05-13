//Go's basic types are

// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte (alias for uint8)
// rune (alias for int32) (represents a Unicode code point)
// float32 float64
// complex64 complex128

package main

import "fmt"

func main() {
	var i int // Varaible decalration
	i = 10    // initialization
	j := 20   // Shorthand for declaring and initializing
	fmt.Println("i + j = ", i+j)

	var k = "golang" // String
	fmt.Println(k)

	var a, b float32 = 1.45, 2.98 // Float32
	fmt.Printf("a = %v, b = %v -> a + b = %v\n", a, b, a+b)

	var c = true // Boolean
	fmt.Println(!c)

	var x, y, z = 0.867 + 0.5i, 4.65, true
	fmt.Println(x) // complex128
	fmt.Println(y)
	fmt.Println(z)

}
