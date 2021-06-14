//Напишите функцию с переменным числом параметров, которая находит наибольшее число в списке
package main

import "fmt"

func max(x...int)int{
	var max int
	for _ , d := range x {
		if d > max {
			max = d
		}
	}
	return max 
}

func main() {
    fmt.Println(max(98,93,77,82,83,1,2,3,100))
}
