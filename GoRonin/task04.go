//Реализовать набор из N воркеров, которые читают из канала произвольные данные и выводят в stdout.
//Данные в канал пишутся из главного потока.
//Необходима возможность выбора кол-во воркеров при старте, а также способ завершения работы всех воркеров.
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan interface{},wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "данные из канала: ", j)
	}
}

const n = 5 // Выбор кол-ва воркеров при старте
func main() {
	wg:=&sync.WaitGroup{}
	jobs := make(chan interface{}, 100)

	for w := 0; w < n; w++ {
		wg.Add(1)
		go worker(w, jobs,wg)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	// Срез с числами
	numbers := []int{10,12,17,13,15}
	//Отправляем числа в канал
    for _, squared := range numbers {
    	jobs <- squared * squared
    }
	close(jobs)
	wg.Wait()
}
