package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

// producer пишет числа в канал до истечения timeout
func producer(dataCh chan<- int, timeout <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done() // Гарантируем сигнал завершения
	for num := 1; ; num++ {
		select {
		case <-timeout:
			close(dataCh) // Закрываем канал
			fmt.Println("Producer завершает работу...")
			return
		case dataCh <- num: // Пишем число в канал
		}
	}
}

// consumer читает числа из канала до его закрытия
func consumer(dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Гарантируем сигнал завершения
	for n := range dataCh {
		fmt.Println("Получено значение:", n)
	}
	fmt.Println("Consumer завершает работу...")
}

func main() {
	// Флаг delay определяет продолжительность работы программы (секунды)
	delay_seconds := flag.Int("delay", 3, "время работы программы в секундах")
	flag.Parse()

	dataCh := make(chan int)                                           // Канал передачи данных между producer и consumer
	timeout := time.After(time.Duration(*delay_seconds) * time.Second) // Таймер завершения

	var wg sync.WaitGroup
	wg.Add(2) // Ожидаем завершения двух горутин

	go producer(dataCh, timeout, &wg)
	go consumer(dataCh, &wg)

	wg.Wait() // Ждём, пока обе горутины завершатся
	fmt.Println("Программа завершена.")
}
