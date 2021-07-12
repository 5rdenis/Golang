//Дана переменная int64. Написать программу которая устанавливает i-й бит в 1 или 0.
package main

import (
  "fmt"
  "strconv"
)

func setBit(n int64, pos uint) int64 {
    n |= (1 << pos)
    return n
}

// Clears the bit at pos in n.
func clearBit(n int64, pos uint) int64 {
    return n &^ (1 << pos)
}

func main() {
    n := int64(3)
    s := strconv.FormatInt(clearBit(n, 1), 2)

    fmt.Printf("%064s", s)
}