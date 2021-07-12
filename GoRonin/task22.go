//Написать быструю сортировку встроенными методами языка.
package main

import (
	"fmt"
)

func partition(a []int, low, high int) ([]int, int) {
	pivot := a[high]
	i := low
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high] = a[high], a[i]
	return a, i
}

func sort(a []int, left, right int) []int {
	if left < right {
		var p int
		a, p = partition(a, left, right)
		a = sort(a, left, p-1)
		a = sort(a, p+1, right)
	}
    
	return a
}

func quickSort(arr []int) []int {
    sort(arr, 0, len(arr) - 1)

    return arr
}

func main() {
	list := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}
	
	quickSort(list)
	fmt.Println(list)
}