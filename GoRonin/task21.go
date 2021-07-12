package main

import (
	"fmt"
	"sync"
)

// Функция будет печатать данные в stdout
func printer( i interface{},wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	// Воспользуемся sync.WaitGroup для ожидания завершения нескольких горутин
	wg:=&sync.WaitGroup{}

	// Создадим массив с данными
	nums := [...]int{1,2,3,5,6,6,4,3,1,2,3,4,5,6,7,2,3,1}
	// Итерируемся по массиву
	for _, num := range nums {
		wg.Add(1)
		go printer(num,wg)
	}
	// Ждем  выполнения работы всех горутин
	wg.Wait()
}