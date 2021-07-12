//Что выведет программа данная программа?
package main

import "fmt"

func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}

//GO VARIABLE SHADOWING
//The statement n := 1 declares a new variable
//which shadows the original n throughout the scope of the if statement.