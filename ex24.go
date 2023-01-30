package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func newPoint(x, y int) *point { //конструктор
	return &point{x, y}
}

func dist(p1, p2 *point) float64 {
	x := p1.x - p2.x
	y := p1.y - p2.y
	sum := float64(x*x + y*y) //сумма квадратов
	return math.Sqrt(sum)     //корень суммы
}

func main() {
	p1 := newPoint(4, 4)
	p2 := newPoint(5, 5)
	fmt.Println(dist(p1, p2))
}
