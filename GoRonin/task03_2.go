//Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов
//с использованием конкурентных вычислений.
package main

import "fmt"
//import "time"
//import "math/rand"

var sum int
var count int

func f(i, x int) {
    count++
    sum += x * x
    fmt.Println("Функция после вычислений","i:",i,"x:",x,"sum:",sum)
    if count == 5 {
    	fmt.Println(sum)
    }
}

func main() {
	x := [5]int{ 2, 4, 6, 8, 10 }
    for i := 0; i < 5; i++ {
        go f(i, x[i])
    }
    var input string
    fmt.Scanln(&input)
}