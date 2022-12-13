package main

import "fmt"

func main() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}
}
