package main

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// 1. Остановка по условию (atomic-флаг)
// Плюсы: просто, быстро. Минусы: нужен atomic или sync, неудобно для большого числа воркеров

func goroutineFlag(stopFlag *int32) {
	for {
		if atomic.LoadInt32(stopFlag) == 1 {
			fmt.Println("[flag] Горутина остановлена по флагу (atomic).")
			return
		}
		fmt.Println("[flag] Работаю...")
		time.Sleep(300 * time.Millisecond)
	}
}

// 2. Остановка через канал
// Плюсы: потокобезопасно, удобно для многих горутин. Минус: канал можно закрыть только один раз

func goroutineChan(stopChan <-chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("[chan] Горутина остановлена через канал.")
			return
		default:
			fmt.Println("[chan] Работаю...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// 3. Остановка через context.Context
// Плюсы: стандарт Go, поддерживает иерархии и таймауты. Минус: чуть сложнее синтаксис

func goroutineContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[context] Горутина остановлена через context.")
			return
		default:
			fmt.Println("[context] Работаю...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// 4. Принудительное завершение через runtime.Goexit
// Плюсы: мгновенное завершение. Минусы: сбрасывает defers, не для обычных задач

func goroutineGoexit() {
	for i := 0; i < 3; i++ {
		fmt.Println("[Goexit] Работаю...", i)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("[Goexit] Завершаю через Goexit (экстренно).")
	runtime.Goexit() // Немедленно завершает текущую горутину
}

// 5. Остановка по таймеру (time.After)
// Плюсы: удобно для автозавершения по времени. Минус: только для фиксированных сценариев

func goroutineTimer() {
	timer := time.After(1 * time.Second)
	for {
		select {
		case <-timer:
			fmt.Println("[timer] Горутина завершена по таймеру.")
			return
		default:
			fmt.Println("[timer] Работаю...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// 6. Остановка по закрытию канала (range)
// Плюсы: лаконично для producer-consumer паттерна. Минус: подходит не для всех задач

func goroutineRangeChan(dataCh <-chan int) {
	for n := range dataCh {
		fmt.Println("[range-chan] Принято значение:", n)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("[range-chan] Горутина завершена после закрытия канала.")
}

// 7. Остановка через panic/recover
// Плюсы: можно ловить фатальные ошибки. Минусы: не для штатного завершения, усложняет отладку

func goroutinePanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[panic/recover] Горутина завершилась по panic:", r)
		}
	}()
	for i := 0; i < 3; i++ {
		fmt.Println("[panic/recover] Работаю...", i)
		time.Sleep(200 * time.Millisecond)
	}
	panic("что-то пошло не так") // Имитация экстренного завершения
}

func main() {
	// ---------- 1. Остановка по флагу ----------
	var stopFlag int32
	go goroutineFlag(&stopFlag)
	time.Sleep(1 * time.Second)
	atomic.StoreInt32(&stopFlag, 1)
	time.Sleep(1 * time.Second) // Ждём завершения горутины

	// ---------- 2. Остановка через канал ----------
	stopChan := make(chan struct{})
	go goroutineChan(stopChan)
	time.Sleep(1 * time.Second)
	close(stopChan)
	time.Sleep(1 * time.Second) // Ждём завершения горутины

	// ---------- 3. Остановка через context.Context ----------
	ctx, cancel := context.WithCancel(context.Background())
	go goroutineContext(ctx)
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second) // Ждём завершения горутины

	// ---------- 4. Принудительное завершение через runtime.Goexit ----------
	go goroutineGoexit()
	time.Sleep(2 * time.Second)

	// ---------- 5. Остановка по таймеру (time.After) ----------
	go goroutineTimer()
	time.Sleep(2 * time.Second)

	// ---------- 6. Остановка по закрытию канала (range) ----------
	dataCh := make(chan int)
	go goroutineRangeChan(dataCh)
	for i := 1; i <= 3; i++ {
		dataCh <- i
	}
	close(dataCh)
	time.Sleep(1 * time.Second)

	// ---------- 7. Остановка через panic/recover ----------
	go goroutinePanicRecover()
	time.Sleep(2 * time.Second)

	fmt.Println("main завершён.")
}
