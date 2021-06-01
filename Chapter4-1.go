//Напишите программу, переводящую температуру из градусов Фаренгейта в градусы Цельсия
package main

import "fmt"

func main() {
    fmt.Print("Enter a number to convert F to C: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := ((input - 32) * 5/9)

    fmt.Println(output)    
}