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
	q := Point{3, 4}

	p.ScaleBy(10)
	fmt.Println(p) // {10 20}

	x := p.ScaleBy
	fmt.Printf("%T\n", x) // func(float64)

	x(100)
	fmt.Println(p) // {1000 2000}

	x = q.ScaleBy
	x(20)
	fmt.Println(p) // {1000 2000}
	fmt.Println(q) // {60 80}
}
