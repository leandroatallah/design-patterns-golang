# Composite

> The composite pattern describes a group of objects that are treated the same way as a single instance of the same type of object. The intent is to "compose" objects into tree structures to represent part-whole hierarchies. Implementing the composite pattern lets clients treat individual objects and compositions uniformly. (Wikipedia)

> **Composite** is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects. (Refactoring Guru)

## Structure

1. The **Component** interface describes operations that are common to both simple and complex elements of the tree.
2. The Leaf is a basic element of a tree that doesn’t have sub-elements.
3. The Container (aka composite) is an element that has sub-elements: leaves or other containers. A container doesn’t know the concrete classes of its children. It works with all sub-elements only via the component interface.
4. The Client works with all elements through the component interface. As a result, the client can work in the same way with both simple or complex elements of the tree.

## Examples

### Simple example

```go
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
```

### Complex example

````go
package main

type NeuronInterface interface {
 Iter() []*Neuron
}

type Neuron struct {
 In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
 return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
 n.Out = append(n.Out, other)
 other.In = append(other.In, n)
}

type NeuronLayer struct {
 Neurons []Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
 result := make([]*Neuron, 0)
 for i := range n.Neurons {
  // result = append(result, &r)
  result = append(result, &n.Neurons[i])
 }
 return result
}

func NewNeuronLayer(count int) *NeuronLayer {
 return &NeuronLayer{make([]Neuron, count)}
}

func Connect(left, right NeuronInterface) {
 for _, l := range left.Iter() {
  for _, r := range right.Iter() {
   l.ConnectTo(r)
  }
 }
}

func main() {
 n1, n2 := &Neuron{}, &Neuron{}
 layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

 Connect(n1, n2)
 Connect(n1, layer2)
 Connect(layer2, n1)
 Connect(layer1, layer2)
}
```# Composite
````
