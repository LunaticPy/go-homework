package main

/*
Фасад

Предоставление упрощенного интерфеса с сложножной системе, библиотеке и тд.

+ Изолирует клиентов от компонентов сложной подсистемы.
-Фасад рискует стать божественным объектом, привязанным ко всем классам программы.


Сотрудник службы поддержки/бот является вашим фасадом ко всем службам и отделам магазина.
Он предоставляет вам упрощённый интерфейс к системе создания заказа, платёжной системе и отделу доставки.


*/
type Phasad interface {
	Call()
	Get()
	Put()
}

type Employee struct {
	A, B, C int
}

func newEmployee() Phasad {
	return Employee{
		A: 1,
		B: 1,
		C: 1,
	}
}

func (s Employee) Call() { /*logic*/ }
func (s Employee) Get()  { /*logic*/ }
func (s Employee) Put()  { /*logic*/ }

type Bot struct {
	A, B, C int
}

func newBot() Phasad {
	return Bot{
		A: 1,
		B: 1,
		C: 1,
	}
}

func (s Bot) Call() { /*logic*/ }
func (s Bot) Get()  { /*logic*/ }
func (s Bot) Put()  { /*logic*/ }
