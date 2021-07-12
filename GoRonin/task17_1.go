//Написать программу, которая в рантайме способна определить тип переменной — int, string, bool, channel из переменной типа interface{}.
package main

import "fmt"

func checkType(val interface{}) {
	switch val.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Int")
	case bool:
		fmt.Println("Bool")
	case struct{}:
		fmt.Println("Struct")
	case chan int:
		fmt.Println("Channel")
	default:
		fmt.Println("This type is undefined")
	}
}

func main(){
	checkType("Sd")
}