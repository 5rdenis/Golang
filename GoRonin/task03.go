//Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов
//с использованием конкурентных вычислений.
package main

import "fmt"
//import "math/rand"
import "sync"

var sum int
var count int

func f(wg *sync.WaitGroup, x int) {
    fmt.Printf("Worker %v: Started\n", x)
    defer wg.Done()
    sum += x * x
    count += 1
    fmt.Printf("Worker %v: Finished\n", x)
  //  fmt.Println("Функция после вычислений","i:",i,"x:",x,"sum:",sum, "count:",count)
}

func main() {
    var wg sync.WaitGroup
	x := [5]int{ 2, 4, 6, 8, 10 }
    for i := 0; i < 5; i++ {
        fmt.Println("Main: Starting worker", i)
        wg.Add(1)
        go f(&wg, x[i])
    }
    fmt.Println("sum:",sum)
    fmt.Println("Main: Waiting for workers to finish")
    wg.Wait()
    fmt.Println("sum:",sum)
    fmt.Println("Main: Completed")
    var input string
    fmt.Scanln(&input)
}