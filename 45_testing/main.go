package main

import "fmt"

func sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func greet(name string) string {
	if len(name) > 0 {
		return fmt.Sprintf("Hello, %s", name)
	}
	return fmt.Sprintf("Hello, World!")
}

func main() {
	fmt.Println(greet("Jacky"))
	fmt.Println(greet(""))
	fmt.Println("Sum of 4 + 5 + 3 + 6 + 8 + 9: ", sum(4, 5, 3, 6, 8, 9))
}
