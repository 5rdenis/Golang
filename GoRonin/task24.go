//Создать слайс с предварительно выделенными 100 элементами.
package main

import "fmt"
//import "unsafe"
//import "reflect"

func main() {
	s := make([]int, 0, 100)
	//hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
    //data := *(*[3]int)(unsafe.Pointer(hdr.Data))
	fmt.Println(s)
	for i:=0;i<20;i++ {
		s = append(s,1)
		fmt.Println(cap(s))
	}
	fmt.Println(len(s),s)
	
}

//https://stackoverflow.com/questions/36706843/how-to-get-the-underlying-array-of-a-slice-in-go
