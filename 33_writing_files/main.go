package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Writing String to a file")
	// The create function creates a file named string.txt with 0666 mode.
	f, err := os.Create("string.txt")
	if err != nil {
		panic(err)
	}

	// It important to close the file,
	// defer will invoke this call at the end
	defer f.Close()

	// Write string to file
	n, err := f.WriteString("Go is fun!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d byte wrote to string.txt\n", n)

	fmt.Println()

	// Writing bytes to a file
	fmt.Println("Writing bytes to a file")
	f2, err := os.Create("byte.txt")
	if err != nil {
		panic(err)
	}

	defer f2.Close()

	data := []byte{71, 111, 32, 105, 115, 32, 102, 117, 110, 33}
	n2, err := f2.Write(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d byte wrote to byte.txt\n", n2)

	fmt.Println()

	// Writing file line by line
	fmt.Println("Writing file line by line")
	f3, err := os.Create("lineByLine.txt")
	if err != nil {
		panic(err)
	}

	defer f3.Close()

	text := []string{
		"Go is a statically typed,",
		"compiled programming language designed at Google",
		"by Robert Griesemer, Rob Pike, and Ken Thompson.",
	}

	n3 := 0 // Number of bytes

	for _, str := range text {
		n3, err = fmt.Fprintln(f3, str)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%d byte wrote to byte.txt\n", n3)
}
