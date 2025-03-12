package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	fmt.Println("The random number is", x)
}
