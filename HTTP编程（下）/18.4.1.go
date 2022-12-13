package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is index")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello World!")
}

func timeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})

}

func main() {
	//创建一个路由转发器
	mux := http.NewServeMux()
	//给hello函数注册timeHandler中间件
	mux.Handle("/hello", timeHandler(http.HandlerFunc(hello)))
	mux.Handle("/", timeHandler(http.HandlerFunc(index)))
	http.ListenAndServe(":8080", mux)
}
