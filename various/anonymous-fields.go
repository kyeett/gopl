package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

func main() {
	fmt.Println("hej")
	p := Point{1, 2}
	c := Circle{p, 10}
	fmt.Println(p)
	fmt.Println(c)
	fmt.Printf("%v, %v, %v", c.X, c.Y, c.Radius)
}
