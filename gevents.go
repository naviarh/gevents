package main

import "time"

type events struct {
	SizeOUT   int        // размер буферов канала событий
	SizeIN    int        // размер буферов канала рассылок
	intIN     chan int   // ресивер воркера (канал событий)
	intOUT    chan int   // сендлер воркера (канал рассылок)
	OUT       chan<- int // воркер пользователя (канал событий)
	IN        <-chan int // ресивер пользователя (канал рассылок)
	reactions int        // количество рассылок событий
}

// Воркер событий
func (e *events) worker() {
	for {
		in := <-e.intIN
		for i := 0; i < e.reactions; i++ {
			e.intOUT <- in
		}
	}
}

// Инициализация целочисленной событийности
func (e *events) Int(args ...interface{}) {
	sizeIN := 100 // дефолтный размер буфера канала рассылок
	if len(args) > 0 {
		if args[0].(int) > 0 {
			sizeIN = args[0].(int)
		}
	}
	sizeOUT := 0 // дефолтный размер буфера канала событий
	if len(args) > 1 {
		if args[1].(int) > 0 {
			sizeIN = args[1].(int)
		}
	}
	e.reactions = 0
	e.intIN = make(chan int, sizeOUT)
	e.intOUT = make(chan int, sizeIN)
	e.IN = e.intOUT
	e.OUT = e.intIN
	go e.worker()
}

// Инкрементор рассылок события
func (e *events) Increment() {
	e.reactions++
}

// Декрементор рассылок события
func (e *events) Decrement() {
	e.reactions--
}

func main() {
	// Экземпляр событийности
	var E events
	E.Int()
	for i := 0; i < 100000; i++ {
		E.Increment()
		go func() {
			for {
				select {
				case ch := <-E.IN:
					print(" <", ch)
				}
			}
		}()
	}
	go func() {
		for i := 0; i < 5; i++ {
			print(" >", i)
			E.OUT <- i
		}
	}()
	time.Sleep((time.Minute))
	println()
}
