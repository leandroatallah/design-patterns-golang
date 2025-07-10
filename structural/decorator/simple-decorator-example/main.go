package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %.2f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %.2f", s.Side)
}

type ColoredShape struct {
	shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.shape.Render(), c.Color)
}

type TransparentShape struct {
	shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %.2f%% transparency", t.shape.Render(), t.Transparency)
}

func main() {
	circle := Circle{2}
	circle.Resize(2)
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())

	rhsShape := TransparentShape{&redCircle, 50.0}
	fmt.Println(rhsShape.Render())
}
