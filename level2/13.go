package main

import "fmt"

/*
Посетитель —
это поведенческий паттерн, который позволяет добавить новую операцию для целой иерархии классов, не изменяя код этих классов.

 + Упрощает добавление операций, работающих со сложными структурами объектов.
 + Объединяет родственные операции в одном классе.
 + Посетитель может накапливать состояние при обходе структуры элементов.

- Паттерн не оправдан, если иерархия элементов часто меняется.
- Может привести к нарушению инкапсуляции элементов.


Мы хотим сделать специальный запрос к каждому типу БД, не изменяя базовый интерфейс


*/

type shape interface {
	getType() string
	accept(visitor)
}

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
}

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}

func main() {
	square := &square{side: 2}
	circle := &circle{radius: 3}

	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)

	fmt.Println()

}
