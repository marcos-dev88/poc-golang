package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	ray float64
}

func (s Square) calcArea() {
	fmt.Println(s.side * s.side)
}

func (c Circle) calcArea() {
	fmt.Println(2 * math.Pi * c.ray)
}

type Figure interface {
	calcArea()
}

func info(f Figure) {
	switch f.(type) {
	case Square:
		fmt.Println("The square's area is:")
	case Circle:
		fmt.Println("The circle's area is:")
	}
	f.calcArea()
}

func exec5() {
	square := Square{
		side: 42.6,
	}

	circle := Circle{
		ray: 56.9,
	}

	info(square)
	fmt.Println()
	info(circle)

	fmt.Println()
}
