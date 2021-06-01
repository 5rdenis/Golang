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