package main

import (
	"fmt"
	"math"
)

type ErrNegSqrt float64

func (e ErrNegSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}

func main() {
	res, err := sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
