package main

import ("fmt")

type point struct {
	x,y int
}

func setPoint(ptr *point) {
	ptr.x = 33
	ptr.y = 34
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
