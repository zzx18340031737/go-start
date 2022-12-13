package main

import "fmt"

func main(){
	var a string = "abc"
	var i interface{} = a
	var b string = i
	fmt.Println(b)

}
