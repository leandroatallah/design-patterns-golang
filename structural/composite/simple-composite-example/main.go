package main

import "fmt"

type Graphic interface {
	Move(x, y int)
	Draw()
}

type Dot struct {
	x, y int
}

func (d *Dot) Move(x, y int) {
	d.x += x
	d.y += y
}

func (d *Dot) Draw() {
	fmt.Println("Dot draw")
}

type Circle struct {
	Dot
	radius int
}

func (c *Circle) Draw() {
	fmt.Println("Circle draw")
}

type CompoundGraphic struct {
	children []Graphic
}

func (cg *CompoundGraphic) Add(child Graphic) {
	cg.children = append(cg.children, child)
}

func (cg *CompoundGraphic) Remove() {
	fmt.Println("Remove child from compound graphic")
}

func (cg *CompoundGraphic) Move(x, y int) {
	for _, child := range cg.children {
		child.Move(x, y)
	}
}

func (cg *CompoundGraphic) Draw() {
	fmt.Println("CompoundGraphic draw")
	for _, child := range cg.children {
		fmt.Println("draw child,", child)
	}
}

type ImageEditor struct {
	All CompoundGraphic
}

func (ie *ImageEditor) load() {
	ie.All = CompoundGraphic{[]Graphic{}}
	ie.All.Add(&Dot{1, 2})
	ie.All.Add(&Circle{Dot{5, 3}, 10})
}

func (ie *ImageEditor) GroupSelected(components []Graphic) {
	group := CompoundGraphic{}
	for _, child := range components {
		group.Add(child)
	}
	ie.All.Add(&group)
	ie.All.Draw()
}

func main() {
	editor := ImageEditor{}
	editor.load()
	editor.GroupSelected([]Graphic{
		&Dot{2, 3},
		&Dot{3, 3},
		&Dot{4, 3},
		&Dot{2, 6},
		&Circle{Dot{2, 6}, 10},
		&Circle{Dot{1, 1}, 2},
	})
}
