package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y}
}

// Вычисление растояния(Евклидова метрика)
func (p *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p.x-p2.x, 2) + math.Pow(p.y-p2.y, 2))
}

func main() {
	p1 := NewPoint(3, 6)
	p2 := NewPoint(4, 8)
	fmt.Println(p1.Distance(p2))
}
