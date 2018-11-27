package main

import (
	"time"
	. "https://github.com/naviarh/gevents"
)

// Перечень событий
var (
	event1 = E.AddEvent()
)

// Функция тестовой реакции на событие
func listen() {
	println("text")
}

func main() {

	// Проверка работы событийного механизма

println("add")
	numb1,_ := E.AddReaction(event1, listen) // подписание слушателя 1 на событие
	numb2,_ := E.AddReaction(event1, listen) // подписание слушателя 2 на событие
	E.DelReaction(numb2) // отписание слушателя 2 от события
	numb2,_ = E.AddReaction(event1, listen) // подписание слушателя 2 на событие
println(numb1,numb2)
	go E.Event(event1) // генерация события
	time.Sleep(time.Second)
println("stop")

	E.StopReaction(numb1)

	go E.Event(event1) // генерация события
	time.Sleep(time.Second)
println("start")

	E.StartReaction(numb1)

	go E.Event(event1) // генерация события
	time.Sleep(time.Second)

	//e.DelEvent(num) // удалить событие

	E.DelReaction(numb1) // отписание слушателя 1 от события

println("del")
	go E.Event(event1) // генерация события
	time.Sleep(time.Second)

	E.DelReaction(numb2) // отписание слушателя 2 от события
	E.Event(event1) // генерация события

	E.DelEvent(event1) // удалить событие
	E.Event(event1) // генерация события

	println(E.DelEvent(event1))

	time.Sleep(time.Second)
}
