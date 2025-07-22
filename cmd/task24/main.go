package main

import (
	"fmt"
	"math"
)

// Структура Point с приватными полями
type Point struct {
	x, y float64
}

// Конструктор - возвращает Point с нужными координатами
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Метод Distance - возвращает расстояние до другой точки
func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Hypot(dx, dy) // sqrt(dx^2 + dy^2)
}

func main() {
	a := NewPoint(1, 2)
	b := NewPoint(4, 6)
	fmt.Printf("Расстояние: %.3f\n", a.Distance(b)) // 5.000
}
