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
	p.ScaleBy(10)  // implicit (&p)
	fmt.Println(p) // {10 20}

	q := &Point{1, 2}
	q.ScaleBy(10)
	fmt.Println(q) // &{10 20}

	// value can not be the receiver of pointer method.
	//Point{3, 4}.ScaleBy(20) // cannot call pointer method on Point literal.
}
