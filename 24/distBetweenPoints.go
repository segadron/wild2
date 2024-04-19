package main

import (
	"fmt"
	"math"
)

// Point - структура, представляющая точку в 2D пространстве с координатами x и y
type Point struct {
	x, y float64
}

// NewPoint возвращает новую точку с заданными координатами x и y
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Distance вычисляет евклидово расстояние между вызывающей точкой Point и другой точкой Point
func (p *Point) Distance(other *Point) float64 {
	// Вычисляет квадрат расстояния между двумя точками
	squaredDistance := (p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y)
	// Вычисляет действительное расстояние, возведя квадратный расстояние в квадратные корень
	return math.Sqrt(squaredDistance)
}

func main() {
	// Создаем две точки
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)

	// Выводим расстояние между двумя точками
	fmt.Println(p1.Distance(p2)) // Вывод: 4.242640687119285
}
