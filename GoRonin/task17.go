//Написать программу, которая в рантайме способна определить тип переменной — int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

func main(){
	// Инициализируем переменную любого типа
	var variable bool
	// Воспользуемся пакетом "reflect"
	xType := reflect.TypeOf(variable)
	// Выведем тип переменной в stdout
	fmt.Println(xType)
}