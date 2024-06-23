package main

import (
	"fmt"
	"math"
)

// определяем структуру Point с двумя инкапсулированными полями x и y типа float64
type Point struct {
	x, y float64
}

// cоздаем конструктор NewPoint(), который принимает значения для x и y и возвращает указатель на новый экземпляр Point
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// определяем метод Distance() на Point, который принимает другую точку и возвращает расстояние между двумя точками, рассчитанное по формуле Пифагора
func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1.5, 2.3)
	p2 := NewPoint(4.2, 6.1)

	distance := p1.Distance(p2)
	fmt.Printf("The distance between (%.2f, %.2f) and (%.2f, %.2f) is %.2f\n",
		p1.x, p1.y, p2.x, p2.y, distance)
}
