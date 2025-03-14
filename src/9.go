package main

import "fmt"

func main() {
	x := make([]int, 5)
	for i := range x {
		x[i] = i + 1
	}
	fmt.Println(x)
}
