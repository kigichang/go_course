package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (pt *Point) String() string {
	return fmt.Sprintf(`(%d,%d)`, pt.X, pt.Y)
}

func main() {
	points := make([]Point, 5)
	fmt.Println("before:", points)

	for i, pt := range points {
		fmt.Printf("%d: %p, %p\n", i, &points[i], &pt)
		pt.X += i
		pt.Y += i
	}

	fmt.Println("after:", points)
}
