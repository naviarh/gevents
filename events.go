package events

// Структура событий
type events struct {
	// Карта функций слушателей событий
	el map[int64]map[int64]func([]interface{}) // map[номер реакции][номер события] := функция реакции
	ll int64 // счётчик номеров реакций
	// Карта карт событий
	en map[int64]map[int64]bool // map[номер события][номер реакции] := признак выполнения
	ln int64 // счётчик номеров событй
}
// Добавить реакцию на событие
func (e *events) AddReaction(ev int64, fn func([]interface{})) (int64, bool) {
	e.ll++ // определение номера реакции
	_, ok := e.en[ev] // проверка существования событя
	if ok {
		e.en[ev][e.ll] = true // добавить номер реакции в карту события
		e.el[e.ll] = make(map[int64]func([]interface{})) // создать карту реакции
		e.el[e.ll][ev] = fn // добавить функцию реакции
	}
	return e.ll, ok
}
// Удалить реакцию на событие
func (e *events) DelReaction(n int64) bool {
	if len(e.el[n]) == 1 { // проверка существования реакции
		for m,_ := range e.el[n] { // вытаскивание номера события
			delete(e.en[m], n) // удалить номер реакции из карты реакций события
			delete(e.el, n) // удалить функцию реакции
		}
	} else {
		return false
	}
	return true
}
// Приостановить реакцию на событие
func (e *events) StopReaction(n int64) bool {
	if len(e.el[n]) == 1 { // проверка существования реакции
		for m,_ := range e.el[n] { // вытаскивание номера события
			e.en[m][n] = false // выключить реакцию
			return true
		}
	}
	return false
}
// Возобновить реакцию на событие
func (e *events) StartReaction(n int64) bool {
	if len(e.el[n]) == 1 { // проверка существования реакции
		for m,_ := range e.el[n] { // вытаскивание номера события
			e.en[m][n] = true // включить реакцию
			return true
		}
	}
	return false
}

// Добавить событие
func (e *events) AddEvent() int64 {
	e.ln++ // определение номера события
	e.en[e.ln] = make(map[int64]bool) // создать карту реакций для события
	return e.ln
}
// Удалить событие
func (e *events) DelEvent(n int64) bool {
	_, ok := e.en[n] // проверка существования события
	if ok {
		for l,_ := range e.en[n] { // Перебрать все реакции на событя
			delete(e.el, l) // удалить функцию реакции
		}
		delete(e.en, n) // удалить событие
	}
	return ok
}
// Метод события
func (e events) Event(n int64, args ...interface{}) {
	for m, b := range e.en[n] {
		if b {
			go e.el[m][n](args)
		}
	}
}
// Создание событийного механизма
var E = events{ el: make(map[int64]map[int64]func([]interface{})), ll: 0, en: make(map[int64]map[int64]bool), ln: 0 }
