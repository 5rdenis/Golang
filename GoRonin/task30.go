//Удалить i-ый элемент из слайса.
package main

import (
    "fmt"
)

func RemoveIndex(s []int, index int) []int {
	if index > len(s)-1 || index < -len(s) {
		fmt.Println("There's no such index in current slice")
		return nil 
	} else if index < 0 {
        index = len(s) + index
	}
    return append(s[:index], s[index+1:]...)
}

func main() {
    all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]
    for i:=-12;i<12;i++ {
    	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    	all = RemoveIndex(all, i)
    	if all != nil {fmt.Println(all)} 
    }
}
