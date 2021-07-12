//Чем завершится данная программа?
package main

import "sync"
import "fmt"

func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}


/*
package main

import "sync"
import "fmt"

func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg *sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(&wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
*/
//The deadlock happens because each worker gets a copy of the original "WaitGroup" variable.
//When workers execute wg.Done() it has no effect on the "WaitGroup" variable in the main goroutine.
