//Какой результат выполнения данного кода и почему?
package main

import "fmt"

func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}

/*We can modify the data of the slice in the literal function,
 however if the pointer to the data changes for any reason or the slice metadata is modified,
  the change can be partially or no visible at all to the outside function.
For example, if the slice gets allocated again, a new location of the memory is used;
 even if the values are the same, the slice points to a new location and therefore no modification of the values will be visible outside,
  since the slices are pointing to two different localtions (the pointer in the slice copy got overwritten).
 */