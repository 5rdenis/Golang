//Реализовать конкурентную запись в map.
package main

import (
    "fmt"
    "sync"
    "time"
)

var m = make(map[string]string)
var l = sync.Mutex{}

func main() {
    go func() {
        for {
            l.Lock()
            m["x1"] = "foo"
            l.Unlock()
        }
    }()
    go func() {
        for {
            l.Lock()
            m["x"] = "foo1"
            l.Unlock()
        }
    }()

    time.Sleep(1 * time.Second)
    fmt.Println(m["x"])
}