//Что выведет данная программа и почему?
package main

import "fmt"

func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}

//you are passing a copy of the slice not the original slice.
/*Map and slice values behave like pointers: they are descriptors that contain pointers to the underlying map or slice data. 
Copying a map or slice value doesn't copy the data it points to.
 Copying an interface value makes a copy of the thing stored in the interface value. 
 If the interface value holds a struct, copying the interface value makes a copy of the struct.
  If the interface value holds a pointer, copying the interface value makes a copy of the pointer, but again not the data it points to.
*/