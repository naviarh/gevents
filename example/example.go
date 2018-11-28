package main

import (
	"time"
	. "../../gevents"
)

// Перечень событий
var (
	event1  = E.AddEvent()
)

// Функция реагтрования на событие
func listen(args []interface{}) {
	print(len(args), " vars ->")
	for i := range args {
		print(" ")
		switch args[i].(type) {
			case string: print(args[i].(string))
			case int: print(args[i].(int))
		}
	}
	println()
}

func main() {

	// Пример работы событийного механизма

println("add")
	numb1,_ := E.AddReaction(event1, listen) // подписание реагирования-1 на событие
	numb2,_ := E.AddReaction(event1, listen) // подписание реагирования-2 на событие
	E.DelReaction(numb2) // отписание реагирования-2 от события
	numb2,_ = E.AddReaction(event1, listen) // подписание реагирования-2 на событие
println(numb1,numb2)
	go E.Event(event1) // событие без аргументов
	time.Sleep(time.Second)
println("stop")

	E.StopReaction(numb1)

	go E.Event(event1, "qwerty", 3.1416, 1234567) // второй аргумент (float) реагирующей функцией не используется
	time.Sleep(time.Second)
println("start")

	E.StartReaction(numb1)

	go E.Event(event1, "qwerty", 1234567) // передача двух аргументов
	time.Sleep(time.Second)

	//e.DelEvent(num) // удалить событие

	E.DelReaction(numb1) // отписание реагирования-1 от события

println("del")
	go E.Event(event1, 1234567) // генерация события
	time.Sleep(time.Second)

	E.DelReaction(numb2) // отписание реагирования-2 от события
	E.Event(event1) // генерация события

	E.DelEvent(event1) // удалить событие
	E.Event(event1) // генерация события

	println(E.DelEvent(event1))

	time.Sleep(time.Second)
}
