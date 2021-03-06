/*
 * Gevents
 * Version 1.1.2
 * https://github.com/naviarh/gevents
 *
 * Event Model for Go
 *
 * Copyright (c) 2018 Navi-Arh
 * Licensed under the MIT license
 * https://raw.githubusercontent.com/naviarh/gevents/master/LICENSE
 *
 */

package events

import "sync"

// Структура событий
type events struct {
	m sync.RWMutex
	// Карта функций реакций на события
	el map[int64]map[int64]func([]interface{}) // map[номер реакции][номер события] := функция реакции
	ll int64                                   // счётчик номеров реакций
	// Карта карт событий
	en map[int64]map[int64]bool // map[номер события][номер реакции] := признак выполнения
	ln int64                    // счётчик номеров событй
}

// Добавить реакцию на событие
func (e *events) AddReaction(ev int64, fn func([]interface{})) (int64, bool) {
	e.ll++ // определение номера реакции
	e.m.Lock()
	_, ok := e.en[ev] // проверка существования событя
	if ok {
		e.en[ev][e.ll] = true                            // добавить номер реакции в карту события
		e.el[e.ll] = make(map[int64]func([]interface{})) // создать карту реакции
		e.el[e.ll][ev] = fn                              // добавить функцию реакции
	}
	e.m.Unlock()
	return e.ll, ok
}

// Удалить реакцию на событие
func (e *events) DelReaction(n int64) bool {
	e.m.Lock()
	if len(e.el[n]) == 1 { // проверка существования реакции
		for m, _ := range e.el[n] { // вытаскивание номера события
			delete(e.en[m], n) // удалить номер реакции из карты реакций события
			delete(e.el, n)    // удалить функцию реакции
		}
	} else {
		e.m.Unlock()
		return false
	}
	e.m.Unlock()
	return true
}

// Приостановить реакцию на событие
func (e *events) StopReaction(n int64) bool {
	e.m.Lock()
	if len(e.el[n]) == 1 { // проверка существования реакции
		for m, _ := range e.el[n] { // вытаскивание номера события
			e.en[m][n] = false // выключить реакцию
			e.m.Unlock()
			return true
		}
	}
	e.m.Unlock()
	return false
}

// Возобновить реакцию на событие
func (e *events) StartReaction(n int64) bool {
	e.m.Lock()
	if len(e.el[n]) == 1 { // проверка существования реакции
		for m, _ := range e.el[n] { // вытаскивание номера события
			e.en[m][n] = true // включить реакцию
			e.m.Unlock()
			return true
		}
	}
	e.m.Unlock()
	return false
}

// Добавить событие
func (e *events) AddEvent() int64 {
	e.ln++ // определение номера события
	e.m.Lock()
	e.en[e.ln] = make(map[int64]bool) // создать карту реакций для события
	e.m.Unlock()
	return e.ln
}

// Удалить событие
func (e *events) DelEvent(n int64) bool {
	e.m.Lock()
	_, ok := e.en[n] // проверка существования события
	if ok {
		for l, _ := range e.en[n] { // Перебрать все реакции на событя
			delete(e.el, l) // удалить функцию реакции
		}
		delete(e.en, n) // удалить событие
	}
	e.m.Unlock()
	return ok
}

// Метод события
func (e events) Event(n int64, args ...interface{}) {
	e.m.RLock()
	for m, b := range e.en[n] {
		if b {
			// (e.el[m][n])(args)
			go (e.el[m][n])(args)
			// go func() { (e.el[m][n])(args) }()
		}
	}
	e.m.RUnlock()
}

// Создание событийного механизма
var E = events{el: make(map[int64]map[int64]func([]interface{})), ll: 0, en: make(map[int64]map[int64]bool), ln: 0}
