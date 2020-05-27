package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Issue an HTTP GET request to a server
	resp, err := http.Get("https://golang.org/doc/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // Close Body at the end

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	fmt.Println()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ { // Prints 10 lines from resp.Body
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil { // checks for error
		panic(err)
	}
}
