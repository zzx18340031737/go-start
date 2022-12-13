package main

import "fmt"

func main() {
	const (
		a = iota          //0
		b                 //沿用上一行iota，此时iota = 1，b=1
		c = "Hello world" //iota +=1,iota=2
		d                 //沿用上一行的"Hello world"，d = "Hello world"，iota += 1，iota=3
		e = iota
	)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
