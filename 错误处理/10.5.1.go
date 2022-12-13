package main

import (
	"errors"
	"fmt"
)

func div(dividend int, divisor int) (int, error) {
	if divisor == 0 { //除数为0则返回错误
		return 0, errors.New("divisor is zero")
	}
	return dividend / divisor, nil
}

func main() {
	res1, err := div(4, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("4/2 =", res1)
	}
	res2, err := div(1, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("1/0 =", res2)
	}

}
