/*
Напишите функцию, которая принимает число, делит его пополам и возвращает true в случае,
если исходное число чётное, и false, если нечетное.
Например, half(1) должна вернуть (0, false), в то время как half(2) вернет (1, true).
 */
package main

import "fmt"

func half(number int) (int,bool) {
	if number % 2 == 0 {
		return number, true
	} else {
		return number, false
	}
}

func main() {
    fmt.Print("Enter a number: ")
    var input int
    fmt.Scanf("%d", &input)
    fmt.Println(half(input))    
}