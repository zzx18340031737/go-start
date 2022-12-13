package main

import (
	"net/http"
	_ "net/http/pprof"
)

func ping(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("{\"message\":\"pong\"}"))
}

func main() {
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8080", nil)
}
