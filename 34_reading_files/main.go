package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//Copies entire file content in memory
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	//This provide more control over reading file
	//YOu open the file and then can perform various read actions on it
	//like below examples
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	//It important to close the file,
	//defer will invoke this call at the end
	defer f.Close()

	b1 := make([]byte, 20) //Allow 20 bytes to be read from file, our entire file is of 15 bytes
	n1, err := f.Read(b1)  //n1 is actual number of bytes read, n1<20, so it will read whole file
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

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
	fmt.Printf("%d bytes @ %d: ", n2, out)
	fmt.Printf("%v\n", string(b2[:n2]))

	out1, err := f.Seek(8, 0) //seeks read pointer to the 8 byte
	if err != nil {
		panic(err)
	}
	b3 := make([]byte, 7)
	n3, err := io.ReadAtLeast(f, b3, 2) // ReadAtLeast reads from f into buf b3 until it has read at least 2 bytes.
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes @ %d: %s\n", n3, out1, string(b3))

	// Seek(0, 0) rewinds reader pointer to start
	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	//This is another way to read the file
	//The bufio package implements a buffered reader
	//that may be useful both for its efficiency with many small reads
	//and because of the additional reading methods it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(10) //Peek returns the next 10 bytes without advancing the reader.
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes: %s\n", len(b4), string(b4))

}
