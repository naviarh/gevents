/*
 * Gevents
 * Version 1.2.alpha
 * https://github.com/naviarh/gevents
 *
 * Event Model for Go
 *
 * Copyright (c) 2018 Navi-Arh
 * Licensed under the MIT license
 * https://raw.githubusercontent.com/naviarh/gevents/master/LICENSE
 *
 */

package alternative

// Структура
type events struct {
	e   int                     // счётчик событй
	se  [][]int                 // слайсы номеров реакций для событи
	seb []bool                  // флаги включенности для событий
	r   int                     // счётчик реакций
	sr  []int                   // номера событий для реакций
	srf []*func(...interface{}) // функции для реакций
	srb []bool                  // флаги включенности для реакций
}

// Создание событийного механизма
var E = events{}

// Добавить событие
func (e *events) AddEvent() int {
	e.se = append(e.se, make([]int, 0)) // создать слайс реакций для события
	e.seb = append(e.seb, true)         // включить событие
	e.e += 1                            // новый номер события
	return e.e - 1
}

// Отключить событие
func (e *events) OffEvent(ine int) bool {
	if ine > e.e { // проверка на корректность номера события
		return false
	}
	e.seb[ine] = false // отключение события
	return true
}

// Включить событие
func (e *events) OnEvent(ine int) bool {
	if ine > e.e { // проверка на корректность номера события
		return false
	}
	e.seb[ine] = true // включение события
	return true
}

// Добавить реакцию на событие
func (e *events) AddReaction(ine int, f func(...interface{})) (int, bool) {
	if ine > e.e { // проверка на корректность номера события
		return 0, false
	}
	e.se[ine] = append(e.se[ine], e.r) // добавление номера реакции в список реакций события
	e.srf = append(e.srf, &f)          // добавление функции реакции
	e.sr = append(e.sr, ine)           // добавение реакции
	e.srb = append(e.srb, true)        // включить реакцию
	e.r += 1                           // новый номер реакции
	return e.r - 1, true
}

// Отключить реакцию
func (e *events) OffReaction(inr int) bool {
	if inr > e.r { // проверка на корректность номера реакции
		return false
	}
	e.srb[inr] = false // отключение реакции
	return true
}

// Включить реакцию
func (e *events) OnReaction(inr int) bool {
	if inr > e.r { // проверка на корректность номера реакции
		return false
	}
	e.srb[inr] = true // включение реакции
	return true
}

// Метод события
func (e events) Event(ine int, args ...interface{}) bool {
	if ine > e.e { // проверка на корректность номера события
		return false
	}
	if !e.seb[ine] { // проверка на включенность события
		return false
	}
	for _, m := range e.se[ine] { // перебор списка событий реакции
		if !e.srb[m] { // если реакция выключена, то пропустить
			continue
		}
		go func() { (*e.srf[m])(args) }() // запуск функции реакции из списка
	}
	return true
}
