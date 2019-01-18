package main

import "fmt"

// Scale ...
type Scale interface {
	ScaleBy(float64)
}

// Point ...
type Point struct {
	X float64
	Y float64
}

// ScaleBy ...
func (p *Point) ScaleBy(a float64) {
	p.X *= a
	p.Y *= a
}

// CallScale ...
func CallScale(s Scale, a float64) {
	s.ScaleBy(a)
}

func main() {
	p := Point{100.0, 200.0}

	fmt.Println(p) // {100 200}
	p.ScaleBy(10)
	fmt.Println(p) // {1000 2000}
	CallScale(&p, 10)
	fmt.Println(p) // {10000 20000}
	//CallScale(p, 10) // cannot use p (type Point) as type Scale in argument to CallScale: Point does not implement Scale (ScaleBy method has pointer receiver)
}
