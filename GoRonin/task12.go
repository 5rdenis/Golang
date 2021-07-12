//Что выводит данная программа и почему?
package main

import "fmt"

func update(p *int) {
  b := 2
  fmt.Println("b",&b)
  fmt.Println("p",p, *p)
  p = &b
  fmt.Println("new p",p, *p)
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(p,*p)
  update(p)
  a = 2
  fmt.Println(p,*p)
}
