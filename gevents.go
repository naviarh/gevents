/*
 * Gevents
 * Version 2
 * https://github.com/naviarh/gevents
 *
 * Event Model for Go
 *
 * Copyright (c) 2018 Navi-Arh
 * Licensed under the MIT license
 * https://raw.githubusercontent.com/naviarh/gevents/master/LICENSE
 *
 */

package gevents

type Interface struct {
	SizeOUT   int
	SizeIN    int
	in        chan interface{}
	out       chan interface{}
	OUT       chan<- interface{}
	IN        <-chan interface{}
	reactions int
}

func (e *Interface) worker() {
	var in interface{}
	for {
		in = <-e.in
		for i := 0; i < e.reactions; i++ {
			e.out <- in
		}
	}
}

func (e *Interface) Init(args ...int) {
	e.SizeIN = 100
	if len(args) > 0 {
		if args[0] > 0 {
			e.SizeIN = args[0]
		}
	}
	e.SizeOUT = 0
	if len(args) > 1 {
		if args[1] > 0 {
			e.SizeIN = args[1]
		}
	}
	e.reactions = 0
	e.in = make(chan interface{}, e.SizeOUT)
	e.out = make(chan interface{}, e.SizeIN)
	e.IN = e.out
	e.OUT = e.in
	go e.worker()
}

func (e *Interface) Add(n ...int) bool {
	if e.reactions < int(^uint(0)>>1) {
		if len(n) > 0 {
			e.reactions += n[0]
		} else {
			e.reactions++
		}
	} else {
		return false
	}
	return true
}

func (e *Interface) Del(n ...int) bool {
	if e.reactions > 0 {
		if len(n) > 0 {
			e.reactions -= n[0]
		} else {
			e.reactions--
		}
	} else {
		return false
	}
	return true
}

type Bool struct {
	SizeOUT   int
	SizeIN    int
	in        chan bool
	out       chan bool
	OUT       chan<- bool
	IN        <-chan bool
	reactions int
}

func (e *Bool) worker() {
	var in bool
	for {
		in = <-e.in
		for i := 0; i < e.reactions; i++ {
			e.out <- in
		}
	}
}

func (e *Bool) Init(args ...int) {
	e.SizeIN = 100
	if len(args) > 0 {
		if args[0] > 0 {
			e.SizeIN = args[0]
		}
	}
	e.SizeOUT = 0
	if len(args) > 1 {
		if args[1] > 0 {
			e.SizeIN = args[1]
		}
	}
	e.reactions = 0
	e.in = make(chan bool, e.SizeOUT)
	e.out = make(chan bool, e.SizeIN)
	e.IN = e.out
	e.OUT = e.in
	go e.worker()
}

func (e *Bool) Add(n ...int) bool {
	if e.reactions < int(^uint(0)>>1) {
		if len(n) > 0 {
			e.reactions += n[0]
		} else {
			e.reactions++
		}
	} else {
		return false
	}
	return true
}

func (e *Bool) Del(n ...int) bool {
	if e.reactions > 0 {
		if len(n) > 0 {
			e.reactions -= n[0]
		} else {
			e.reactions--
		}
	} else {
		return false
	}
	return true
}

type Int struct {
	SizeOUT   int
	SizeIN    int
	in        chan int
	out       chan int
	OUT       chan<- int
	IN        <-chan int
	reactions int
}

func (e *Int) worker() {
	var in int
	for {
		in = <-e.in
		for i := 0; i < e.reactions; i++ {
			e.out <- in
		}
	}
}

func (e *Int) Init(args ...int) {
	e.SizeIN = 100
	if len(args) > 0 {
		if args[0] > 0 {
			e.SizeIN = args[0]
		}
	}
	e.SizeOUT = 0
	if len(args) > 1 {
		if args[1] > 0 {
			e.SizeIN = args[1]
		}
	}
	e.reactions = 0
	e.in = make(chan int, e.SizeOUT)
	e.out = make(chan int, e.SizeIN)
	e.IN = e.out
	e.OUT = e.in
	go e.worker()
}

func (e *Int) Add(n ...int) bool {
	if e.reactions < int(^uint(0)>>1) {
		if len(n) > 0 {
			e.reactions += n[0]
		} else {
			e.reactions++
		}
	} else {
		return false
	}
	return true
}

func (e *Int) Del(n ...int) bool {
	if e.reactions > 0 {
		if len(n) > 0 {
			e.reactions -= n[0]
		} else {
			e.reactions--
		}
	} else {
		return false
	}
	return true
}

type String struct {
	SizeOUT   int
	SizeIN    int
	in        chan string
	out       chan string
	OUT       chan<- string
	IN        <-chan string
	reactions int
}

func (e *String) worker() {
	var in string
	for {
		in = <-e.in
		for i := 0; i < e.reactions; i++ {
			e.out <- in
		}
	}
}

func (e *String) Init(args ...int) {
	e.SizeIN = 100
	if len(args) > 0 {
		if args[0] > 0 {
			e.SizeIN = args[0]
		}
	}
	e.SizeOUT = 0
	if len(args) > 1 {
		if args[1] > 0 {
			e.SizeIN = args[1]
		}
	}
	e.reactions = 0
	e.in = make(chan string, e.SizeOUT)
	e.out = make(chan string, e.SizeIN)
	e.IN = e.out
	e.OUT = e.in
	go e.worker()
}

func (e *String) Add(n ...int) bool {
	if e.reactions < int(^uint(0)>>1) {
		if len(n) > 0 {
			e.reactions += n[0]
		} else {
			e.reactions++
		}
	} else {
		return false
	}
	return true
}

func (e *String) Del(n ...int) bool {
	if e.reactions > 0 {
		if len(n) > 0 {
			e.reactions -= n[0]
		} else {
			e.reactions--
		}
	} else {
		return false
	}
	return true
}
