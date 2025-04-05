package main

import "fmt"

func main() {
    // Example usage of Go's loop and conditionals
    var i int = 10;
    if i < 5 {
        fmt.Println("i is less than 5")
    } else {
        for i <= 10 {
            if i == 8 {
                break
            }
            fmt.Println(i)
        }
    }
}
