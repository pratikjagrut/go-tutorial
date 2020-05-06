//A map maps keys to values.
package main

import "fmt"

func main() {
	//Create empty map
	m := make(map[int]string) //make(map[key-type]val-type)

	m[1] = "go"
	m[2] = "is"
	m[3] = "fun"

	fmt.Println(m)

	//another way of declaring map
	n := map[string]int{"foo": 0, "bar": 1}
	fmt.Println(n)

	//To know if key exist in map
	_, k := m[4] //_ is for ignoring the value, k is bool for key existance
	fmt.Println(k)
}
