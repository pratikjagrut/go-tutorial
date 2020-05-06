//This program illustrate if else and switch control statements

package main

import (
	"fmt"
	"time"
)

func main() {
	// ********* if, example *********
	if true {
		fmt.Println("True")
	}

	// if else, example

	if 5%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd nubmer")
	}

	// if else ladder

	i := 10
	if i == 0 {
		fmt.Println("It's 0")
	} else if i < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Positive number")
	}

	// ********* switch, example *********
	//Basic switch, multiple expresssions can be seperated by commas
	j := "A"
	fmt.Print(j, " for ")
	switch j {
	case "a", "A":
		fmt.Println("apple")
	case "b", "B":
		fmt.Println("ball")
	case "c", "C":
		fmt.Println("cat")
	default:
		fmt.Println("Unknown letter")
	}

	// this switch can be used insteaf of if/else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//A type switch compares types instead of values.
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
