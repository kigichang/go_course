package main

import "fmt"

// Point ...
type Point struct{ X, Y float64 }

// ScaleBy ...
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	p.ScaleBy(10)
	fmt.Println(p) // {1 2}

	q := &Point{1, 2}
	q.ScaleBy(10)
	fmt.Println(q) // &{1 2}
}
