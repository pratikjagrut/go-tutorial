package main

import "fmt"

// Constant 
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

func main() {
	fmt.Println("Integer range on this computer")
	fmt.Println("MinUint:", MinUint)
	fmt.Println("MaxUint:", MaxUint)
	fmt.Println("MinInt:", MinInt)
	fmt.Println("MaxInt:", MaxInt)
}