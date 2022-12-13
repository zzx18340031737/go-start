package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is an error")
	var err2 error
	fmt.Println(err.Error())
	fmt.Println(err2)
}
