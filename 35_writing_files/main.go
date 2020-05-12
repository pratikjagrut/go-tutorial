package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//writing bytes in file
	data := []byte("hello\nworld!\n")
	//0644 is linux file permissions, if file does not exitst it will create with thos specified permission
	err := ioutil.WriteFile("data", data, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("data2")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	str := "Go\n"
	data2 := []byte(str)
	n2, err := f.Write(data2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("Is\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() //Issue a Sync to flush writes to stable storage.

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("Fun!\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
