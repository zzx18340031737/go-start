package main

import (
	"fmt"
	"log"
	"net/http"
)

func myfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/hello", myfunc)
	log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil))
}
