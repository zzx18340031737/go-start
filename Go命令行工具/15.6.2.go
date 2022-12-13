package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

const Num int = 10000

func concat_str(writer http.ResponseWriter, request *http.Request) {
	var str string

	for i := 0; i < Num; i++ {
		str = fmt.Sprintf("%s%d", str, i)
	}
}

func main() {
	http.HandleFunc("/str", concat_str)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
