package main

import "fmt"

func main() {
	//Create empty map
	m := make(map[int]string) //make(map[key-type]val-type)
	fmt.Println(m)
	m[1] = "go"
	m[2] = "is"
	m[3] = "fun"

	fmt.Println(m)

	//another way of declaring map
	n := map[string]int{"foo": 0, "bar": 1}
	fmt.Println(n)

	//Delete an element
	delete(m, 2)
	fmt.Println(m)

	//To know if key exist in map
	_, k := m[2] //_ is for ignoring the value, k is bool for key existance
	fmt.Println(k)
}
