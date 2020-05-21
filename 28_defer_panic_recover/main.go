package main

import (
	"fmt"
)

func main() {
	fmt.Println("Calling firstFunc from main")
	defer func() {
		fmt.Println("Cleaning function ...")
	}()
	firstFunc()
	fmt.Println("Returned normally from firstFunc.")
}

func firstFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in firstFunc: ", err)
		}
	}()

	fmt.Println("Calling secondFunc from firstFunc")
	secondFunc(0)
	fmt.Println("Returned normally from secondFunc.")
}

func secondFunc(i int) {
	var arr [2]int
	arr[i] = i + 1
	defer fmt.Printf("Defer -> Index: %d, Value: %d\n", i, arr[i])
	fmt.Printf("Index: %d, Value: %d\n", i, arr[i])
	secondFunc(i + 1)
}
