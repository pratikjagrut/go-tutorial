package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("First name: ")
	firstName, _ := reader.ReadString('\n') //Reads unitl \n char is found

	scanner := bufio.NewScanner(os.Stdin)
	var lastName string
	fmt.Println("Last name: ")
	for scanner.Scan() {
		lastName = scanner.Text()
		break //otherwise scan will ask infinite time
	}

	fmt.Printf("Hello, %v %v", lastName, firstName)
}
