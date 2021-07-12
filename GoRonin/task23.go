//Написать бинарный поиск встроенными методами языка.
package main

import "fmt"


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

main() {

}