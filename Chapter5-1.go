package main

import "fmt"

func main() {
    fmt.Println("Числа кратные трём:")
    for i := 1; i <= 100; i++ {
       if (i % 3 == 0) {
            fmt.Println(i)
        }
    }
}