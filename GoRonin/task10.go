//Дана последовательность температурных колебаний (-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5).
//Объединить данный значения в группы с шагов в 10 градусов. Последовательность в подмножностве не важна.
package main 

import "fmt"
//import "math"

func getNum(number float32) int {
	if number < 0 {
		return int(number/10-1) * 10
	} else {
		return int(number/10) * 10
	}
}

func main() {
    temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 1.0, 0.5, 0.0 , 10.0, -10.0, -5.5, -9.9, 2.0, -15.0}
    tempGroups := make(map[int][]float32)
    for _, temp := range temperatures {
    	//rounded := math.Trunc(temp/10) * 10
    	rounded := getNum(temp)
    	tempGroups[rounded] = append(tempGroups[rounded],temp)
    }
    fmt.Printf("%v\n",temperatures)
    fmt.Println(tempGroups)
}