//A common use of panic is to abort if a function returns an error value
//that we don’t know how to (or want to) handle.
package main

import "os"

func main() {

	panic("something went unexpectedly wrong")

	_, err := os.Open("/tmp/file")
	if err != nil {
		panic(err)
	}
}
