//Что выводит данная программа и почему?
package main

import "fmt"

func update(p *int) {
  b := 2
  fmt.Println(p)
  p = &b
  fmt.Println(p)
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}

//я всего лишь меняю адрес записанный в переменную,
// но не меняю содержимое которое лежит по адресу