package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// Reading an entire file into memory
	fmt.Println("***Reading an entire file into memory***")
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	fmt.Println()

	fmt.Println("***Reading a file with more control***")
	// This provide more control over reading file
	// YOu open the file and then can perform various read actions on it
	// like below examples
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	// It important to close the file,
	// defer will invoke this call at the end
	defer f.Close()

	b1 := make([]byte, 20) // Allow 20 bytes to be read from file, our entire file is of 15 bytes
	n1, err := f.Read(b1)  // n1 is actual number of bytes read, n1<20, so it will read whole file
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes:  \t%s\n", n1, string(b1[:n1]))

	fmt.Println()

	//You can also Seek to a known location in the file and Read from there.
	//Seek(x ,y), x = offset for the next read,
	//y = 0(from start),1(current seek point),2(from end)
	out, err := f.Seek(5, 0)
	if err != nil {
		panic(err)
	}
	b2 := make([]byte, 2) //Allow 2 bytes
	n2, err := f.Read(b2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read ptr @ %d:  \t%v\n", out, string(b2[:n2]))

	fmt.Println()

	_, err = f.Seek(8, 0) //seeks read pointer to the 8 byte
	if err != nil {
		panic(err)
	}
	b3 := make([]byte, 7)
	n3, err := io.ReadAtLeast(f, b3, 2) // ReadAtLeast reads from f into buf b3 until it has read at least 2 bytes.
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes read: \t%s\n", n3, string(b3))

	// Seek(0, 0) rewinds reader pointer to start
	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	// Reading a file in small chunks
	fmt.Println("***Reading file in chunk of 2***")
	r4 := bufio.NewReader(f) // create a new buffered reader
	b4 := make([]byte, 2)    // byte slice of length and capacity 2
	for {
		n, err := r4.Read(b4)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b4[0:n]))
	}

	fmt.Println()

	// Reading a file line by line
	fmt.Println("***Reading a file line by line***")
	f1, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	defer f1.Close()

	s := bufio.NewScanner(f1)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println()

}
