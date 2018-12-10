package gevents

// Структура Multi
type Multi struct {
	SizeOUT   int                // размер буферов канала событий
	SizeIN    int                // размер буферов канала рассылок
	in        chan interface{}   // ресивер воркера (канал событий)
	out       chan interface{}   // сендлер воркера (канал рассылок)
	OUT       chan<- interface{} // воркер пользователя (канал событий)
	IN        <-chan interface{} // ресивер пользователя (канал рассылок)
	reactions int                // количество рассылок событий
}

// Воркер событий Multi
func (e *Multi) worker() {
	var in interface{}
	for {
		in = <-e.in
		for i := 0; i < e.reactions; i++ {
			e.out <- in
		}
	}
}

// Инициализация событийности Multi
func (e *Multi) Init(args ...int) {
	e.SizeIN = 100 // дефолтный размер буфера канала рассылок
	if len(args) > 0 {
		if args[0] > 0 {
			e.SizeIN = args[0]
		}
	}
	e.SizeOUT = 0 // дефолтный размер буфера канала событий
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

// Инкрементор рассылок события
func (e *Multi) Increment(n ...int) bool {
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

// Декрементор рассылок события
func (e *Multi) Decrement(n ...int) bool {
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

// Структура Bool
type Bool struct {
	SizeOUT   int         // размер буферов канала событий
	SizeIN    int         // размер буферов канала рассылок
	in        chan bool   // ресивер воркера (канал событий)
	out       chan bool   // сендлер воркера (канал рассылок)
	OUT       chan<- bool // воркер пользователя (канал событий)
	IN        <-chan bool // ресивер пользователя (канал рассылок)
	reactions int         // количество рассылок событий
}

// Воркер событий int
func (e *Bool) worker() {
	var in bool
	for {
		in = <-e.in
		for i := 0; i < e.reactions; i++ {
			e.out <- in
		}
	}
}

// Инициализация событийности int
func (e *Bool) Init(args ...int) {
	e.SizeIN = 100 // дефолтный размер буфера канала рассылок
	if len(args) > 0 {
		if args[0] > 0 {
			e.SizeIN = args[0]
		}
	}
	e.SizeOUT = 0 // дефолтный размер буфера канала событий
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

// Инкрементор рассылок события
func (e *Bool) Increment(n ...int) bool {
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

// Декрементор рассылок события
func (e *Bool) Decrement(n ...int) bool {
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

// Структура Int
type Int struct {
	SizeOUT   int        // размер буферов канала событий
	SizeIN    int        // размер буферов канала рассылок
	in        chan int   // ресивер воркера (канал событий)
	out       chan int   // сендлер воркера (канал рассылок)
	OUT       chan<- int // воркер пользователя (канал событий)
	IN        <-chan int // ресивер пользователя (канал рассылок)
	reactions int        // количество рассылок событий
}

// Воркер событий int
func (e *Int) worker() {
	var in int
	for {
		in = <-e.in
		for i := 0; i < e.reactions; i++ {
			e.out <- in
		}
	}
}

// Инициализация событийности int
func (e *Int) Init(args ...int) {
	e.SizeIN = 100 // дефолтный размер буфера канала рассылок
	if len(args) > 0 {
		if args[0] > 0 {
			e.SizeIN = args[0]
		}
	}
	e.SizeOUT = 0 // дефолтный размер буфера канала событий
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

// Инкрементор рассылок события
func (e *Int) Increment(n ...int) bool {
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

// Декрементор рассылок события
func (e *Int) Decrement(n ...int) bool {
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

// Структура String
type String struct {
	SizeOUT   int           // размер буферов канала событий
	SizeIN    int           // размер буферов канала рассылок
	in        chan string   // ресивер воркера (канал событий)
	out       chan string   // сендлер воркера (канал рассылок)
	OUT       chan<- string // воркер пользователя (канал событий)
	IN        <-chan string // ресивер пользователя (канал рассылок)
	reactions int           // количество рассылок событий
}

// Воркер событий String
func (e *String) worker() {
	var in string
	for {
		in = <-e.in
		for i := 0; i < e.reactions; i++ {
			e.out <- in
		}
	}
}

// Инициализация событийности String
func (e *String) Init(args ...int) {
	e.SizeIN = 100 // дефолтный размер буфера канала рассылок
	if len(args) > 0 {
		if args[0] > 0 {
			e.SizeIN = args[0]
		}
	}
	e.SizeOUT = 0 // дефолтный размер буфера канала событий
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

// Инкрементор рассылок события
func (e *String) Increment(n ...int) bool {
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

// Декрементор рассылок события
func (e *String) Decrement(n ...int) bool {
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
