package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// worker читает числа из канала и выводит их, корректно завершается по отмене или закрытию канала
func worker(ctx context.Context, id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Всегда сигналим о завершении воркера
	for {
		select {
		case <-ctx.Done(): // Если пришла отмена через контекст (например, по Ctrl+C)
			fmt.Printf("worker #%d: завершаюсь...\n", id)
			return
		case num, ok := <-dataCh: // Получаем число из канала
			if !ok { // Если канал закрыт — выходим
				fmt.Printf("worker #%d: канал закрыт, выхожу...\n", id)
				return
			}
			fmt.Printf("worker #%d: %d\n", id, num)
		}
	}
}

// producer — пишет последовательные числа в канал до отмены контекста
func producer(ctx context.Context, dataCh chan<- int, done chan<- struct{}) {
	num := 1
	for {
		select {
		case <-ctx.Done(): // Получили отмену — завершаем работу
			close(done) // Сообщаем main, что producer завершён
			return
		case dataCh <- num: // Пишем число в канал
			num++
			time.Sleep(100 * time.Millisecond) // Задержка для наглядности
		}
	}
}

// handleSignals — слушает Ctrl+C (SIGINT) и вызывает cancel
func handleSignals(cancelFunc context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)   // Канал для приёма сигналов ОС
	signal.Notify(sigCh, os.Interrupt) // Подписка на SIGINT (Ctrl+C)
	<-sigCh                            // Блокируемся, ждём сигнал
	fmt.Println("\nПолучен сигнал, завершаем работу...")
	cancelFunc() // Дёргаем cancel() для завершения всей системы
}

func main() {
	nWorkers := flag.Int("n", 3, "number of workers")
	flag.Parse()

	dataCh := make(chan int) // Канал для передачи чисел от producer к воркерам
	var wg sync.WaitGroup    // WaitGroup для ожидания завершения воркеров

	ctx, cancel := context.WithCancel(context.Background()) // Контекст с возможностью отмены
	defer cancel()

	done := make(chan struct{}) // Канал для сигнала, что producer завершился

	// Запуск воркеров
	for i := 1; i <= *nWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, dataCh, &wg)
	}

	// Запуск producer и обработчика сигналов в отдельных горутинах
	go producer(ctx, dataCh, done)
	go handleSignals(cancel)

	<-done        // Ждём, пока producer завершится (после Ctrl+C)
	close(dataCh) // Закрываем канал - воркеры корректно завершатся

	wg.Wait() // Ждём завершения всех воркеров
	fmt.Println("Все воркеры остановлены.")
}
