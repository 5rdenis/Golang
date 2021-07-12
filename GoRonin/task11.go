//Написать пересечение двух неупорядоченных массивов.
package main

import (
    "fmt"
)

func getIntersection(a []int, b []int) ([]int) {
    m := make(map[int]int)

    for _, k := range a {
        m[k] += 1
    }

    for _, k := range b {
        m[k] += 2
    }

    result := []int{}
        for k, v := range m {
        	if v == 3 {
                result = append(result, k)
        	}
        }
    return result
}

func main() {
    a := []int{1,2,3,4,5}
    b := []int{4,5,6,3,2}

    fmt.Println("Пересечение a и b:", getIntersection(a, b))
}
