package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Server listening......")
	defer log.Println("Server served")
	select {
	case <-time.After(3 * time.Second):
		fmt.Fprint(w, "Welcome to go world")
	case <-ctx.Done():
		log.Println(ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}
