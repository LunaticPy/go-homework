package pattern

import "fmt"

/*

Цепочка обязанностей — это поведенческий паттерн проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков. Каждый последующий обработчик решает,
может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

 + Уменьшает зависимость между клиентом и обработчиками.
 + Реализует принцип единственной обязанности.
 + Реализует принцип открытости/закрытости.

- Запрос может остаться никем не обработанным.


Конвеер - RestApi - микросервисы


*/

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type department interface {
	execute(*patient)
	setNext(department)
}

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}
