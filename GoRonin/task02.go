//Написать программу, которая конкурентно рассчитает значение квадратов
//значений взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import "fmt"

func f(x int) {
    x *= x
    fmt.Println(x)
}

func main() {
	x := [5]int{ 2, 4, 6, 8, 10 }
    for i := 0; i < 5; i++ {
        go f(x[i])
    }
    var input string
    fmt.Scanln(&input)
}



/* for _,value := range x {
	go f(value)
}
*/