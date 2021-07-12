//Реализовать конкурентную запись в map.
package main

import "sync"

type myMap struct {
    sync.Map
}

func main() {
   m := &myMap{}
   m[`foo`] = 1

   var wg sync.WaitGroup

   wg.Add(2)
   go func() {
      for i := 0; i < 1000; i++  {
         m[`foo`]++
      }
   }()
   go func() {
      for i := 0; i < 1000; i++  {
         m[`foo`]++
      }
   }()
   wg.Wait()
}


//syncmap rwmutex