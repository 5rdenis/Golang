//Написать быструю сортировку встроенными методами языка.
package main

import "fmt"
import "math/rand"

func qsort(a []int) []int {
  if len(a) < 2 { return a }

  left, right := 0, len(a) - 1

  // Pick a pivot
  pivotIndex := rand.Int() % len(a)

  // Move the pivot to the right
  a[pivotIndex], a[right] = a[right], a[pivotIndex]

  // Pile elements smaller than the pivot on the left
  for i := range a {
    if a[i] < a[right] {
      a[i], a[left] = a[left], a[i]
      left++
    }
  }

  // Place the pivot after the last smaller element
  a[left], a[right] = a[right], a[left]

  // Go down the rabbit hole
  qsort(a[:left])
  qsort(a[left + 1:])


  return a
}

func binarySearch (a []int, item int) int {
  left := 0
  right := len(a) - 1
  mid := len(a) / 2
  for left <= right {
    value := a[mid]
    if value == item {
        return mid
    }
    if value > item {
        right = mid - 1
        mid = (left + right) / 2
        continue
    }
    left = mid + 1
    mid = (left + right) / 2
}
  return -1
}


func main() {
  a := []int{5,10,3,-2,8,17,11,2,0,3}
  fmt.Println(qsort(a))
  fmt.Println(binarySearch(qsort(a),3))
}