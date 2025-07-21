package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

// воркеры читают из dataCh и выводят в stdout
// каждый воркер — отдельная горутина, которая ждёт данные из канала
func worker(id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()           // сообщаем WaitGroup о завершении работы воркера
	for num := range dataCh { // читаем данные из канала, пока он не будет закрыт
		fmt.Printf("worker #%d: %d\n", id, num)
	}
}

func main() {
	// получаем количество воркеров через флаг -n (по умолчанию 3)
	nWorkers := flag.Int("n", 3, "number of workers")
	flag.Parse()

	dataCh := make(chan int) // канал для передачи данных от производителя к воркерам
	var wg sync.WaitGroup    // WaitGroup для отслеживания завершения всех воркеров

	// запускаем указанное количество воркеров
	for i := 1; i <= *nWorkers; i++ {
		wg.Add(1) // увеличиваем счётчик ожидаемых воркеров
		go worker(i, dataCh, &wg)
	}

	// горутина-producer пишет данные в канал
	go func() {
		num := 1
		for {
			dataCh <- num // отправляем число в канал
			num++
			time.Sleep(100 * time.Millisecond) // чтобы не спамить в консоль
		}
	}()

	wg.Wait() // ожидаем завершения всех воркеров
}
