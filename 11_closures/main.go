//Go functions may be closures.
//A closure is a function value that references variables from outside its body.
//The function may access and assign to the referenced variables;
//in this sense the function is "bound" to the variables.

package main

import "fmt"

func getLimit() func() int {
	limit := 10
	return func() int {
		limit--
		return limit
	}
}

func main(){
	//we can call it by assigning it to a varaible
	showLimit := getLimit()
	fmt.Println(showLimit())
	fmt.Println(showLimit())
	fmt.Println(showLimit())

	fmt.Println()

	//each time we assign it to a new varaible limit resets
	showLimit1 := getLimit()
	fmt.Println(showLimit1())
	fmt.Println(showLimit1())
	fmt.Println(showLimit1())
}