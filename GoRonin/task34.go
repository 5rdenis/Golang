//Написать программу, которая проверяет, что все символы в строке уникальные.
package main

import "fmt"

func unique(arr string) bool {
    m := make(map[rune]bool)
    for _, i := range arr {
        _, ok := m[i]
        if ok {
            return false
        }

        m[i] = true
    }

    return true
}

func main() {
    fmt.Println(unique("123lksadiopj4"))
}