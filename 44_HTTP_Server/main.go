package main

import (
	"fmt"
	"net/http"
)

//ping function responds with pong
func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

//It gives headers information
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	//We register our handlers on server routes using the http.HandleFunc
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Please hit endpoint /ping or /header")
	})
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/header", headers)

	//we call the ListenAndServe with the port and a handler
	http.ListenAndServe(":8081", nil) //nil tells it to use the default router we’ve just set up

	//TO run: go run main.go
	//In other terminal run following, or hit those URL in any browser
	//curl http:/localhost:8081/
	//curl http:/localhost:8081/hello
	//curl http:/localhost:8081/headers
}
