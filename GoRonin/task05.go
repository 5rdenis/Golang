//Написать программу, которая будет последовательно писать значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершиться.
package main

import (
  "fmt"
  "time"
)
func main() {

  const N = 1 // N секунд по истечению которых программа завершится
  c1 := make(chan int)

  go func(ch <-chan int) {
    for v := range ch {
        time.Sleep(500*time.Millisecond)
        fmt.Println(v)
    }
  }(c1)
  go func(ch chan<- int) {
    for i := 0; ; i++ {
    ch <- i
    }
  }(c1)

  timer1 := time.NewTimer(time.Second * N)
  <-timer1.C
  close(c1)
}


