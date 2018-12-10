package main

import (
	"time"
	. ".."
	// . "github.com/naviarh/gevents"
)

var tt int64 = 0.

func empty() {
	var t0 time.Time
	var t1 time.Time
	var n = 10000000
	for i := 0; i < n; i++ {
		t0 = time.Now()
		t1 = time.Now()
		tt += t1.UnixNano() - t0.UnixNano()
	}
	tt = tt / int64(n)
	println("Empty operation: ", tt, "ns")
}

func fo(args ...interface{}) {
	return
}
func duration_fn() {
	var t0 time.Time
	var t1 time.Time
	var t int64 = 0.
	var n = 10000000
	for i := 0; i < n; i++ {
		t0 = time.Now()
		fo(12)
		t1 = time.Now()
		t += t1.UnixNano() - t0.UnixNano()
	}
	t = t / int64(n)
	println("fn: ", t, "ns")
	println("fn: ", t-tt, "ns")
}

func duration_go() {
	var t0 time.Time
	var t1 time.Time
	var t int64 = 0.
	var n = 10000000
	for i := 0; i < n; i++ {
		t0 = time.Now()
		go fo()
		t1 = time.Now()
		t += t1.UnixNano() - t0.UnixNano()
	}
	t = t / int64(n)
	println("go fn: ", t, "ns")
	println("go fn: ", t-tt, "ns")
}

func duration_gofn() {
	var t0 time.Time
	var t1 time.Time
	var t int64 = 0.
	var n = 10000000
	for i := 0; i < n; i++ {
		t0 = time.Now()
		go func() { fo() }()
		t1 = time.Now()
		t += t1.UnixNano() - t0.UnixNano()
	}
	t = t / int64(n)
	println("go func fn: ", t, "ns")
	println("go func fn: ", t-tt, "ns")
}

func forSend(args ...interface{}) {
	for {
		select {
		case <-args[0].(chan int):
			// return 3 * 27
		}
	}
}
func duration_send() {
	var ch = make(chan int)
	go forSend(ch)
	time.Sleep(time.Second)
	var t0 time.Time
	var t1 time.Time
	var t int64 = 0.
	var n = 10000000
	for i := 0; i < n; i++ {
		t0 = time.Now()
		ch <- 10
		t1 = time.Now()
		t += t1.UnixNano() - t0.UnixNano()
		// <-ch
	}
	t = t / int64(n)
	println("send: ", t, "ns")
	println("send: ", t-tt, "ns")
}

func fn(args []interface{}) {
	// print(args[0].(int))
	return
}

var event = E.AddEvent()

func duration_AddReaction() {
	// var reaction uint
	var t0 time.Time
	var t1 time.Time
	var t int64 = 0.
	var n = 100
	for i := 0; i < n; i++ {
		t0 = time.Now()
		E.AddReaction(event, fn)
		// go func() { _, _ = E.AddReaction(event, fn) }()
		t1 = time.Now()
		t += t1.UnixNano() - t0.UnixNano()
		//E.DelReaction(reaction)
	}
	//E.DelEvent(event)
	t = t / int64(n)
	println("addReaction: ", t, "ns")
	println("addReaction: ", t-tt, "ns")
}

func duration_Event() {
	// var event = E.AddEvent()
	// var reaction uint
	var t0 time.Time
	var t1 time.Time
	var t int64 = 0.
	var n = 20000
	for i := 0; i < n; i++ {
		// E.AddReaction(event, fn)
		t0 = time.Now()
		E.Event(event, i)
		// go E.Event(event, i)
		// go func() { E.Event(event, i) }()
		t1 = time.Now()
		t += t1.UnixNano() - t0.UnixNano()
		//E.DelReaction(reaction)
	}
	//E.DelEvent(event)
	t = t / int64(n)
	println("Event: ", t, "ns")
	println("Event: ", t-tt, "ns")
}

func main() {
	println()
	empty()
	duration_fn()
	duration_go()
	duration_gofn()
	duration_send()
	duration_AddReaction()
	duration_Event()
	println()
}
