package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// ********* switch, example *********
	//Basic switch, multiple expresssions can be seperated by commas
	A := "A"
	fmt.Print(A, " for ")
	switch A {
	case "a", "A":
		fmt.Println("apple")
	case "b", "B":
		fmt.Println("ball")
	default:
		fmt.Println("Unknown letter")
	}

	//Prints which OS you're using
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	//Switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//A type switch compares types. This can be used to know type of value
	datatype := func(p interface{}) {
		switch p.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Unknown datatype")
		}
	}
	datatype(true)
	datatype(1)
	datatype("go")

}
