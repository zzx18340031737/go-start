package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is index")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
