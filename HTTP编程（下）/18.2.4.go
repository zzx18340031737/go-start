package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	http.HandleFunc("/", index)

	//设置静态目录
	fsh := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("../../go-start-one/", fsh))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
