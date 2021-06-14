//Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. (3, 6, 9, …)
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