package main

import (
	"log"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("./static/"))
	log.Fatal(http.ListenAndServe(":8080", h))
}
