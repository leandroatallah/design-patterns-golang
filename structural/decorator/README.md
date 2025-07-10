# Decorator

> Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors. (Refactoring Guru)

> The decorator pattern allows behavior to be added to an individual object, dynamically, without affecting the behavior of other instances of the same class. (Wikipedia)

## Structure

1. The Component declares the common interface for both wrappers and wrapped objects.
2. Concrete Component is a class of objects being wrapped. It defines the basic behavior, which can be altered by decorators.
3. The Base Decorator class has a field for referencing a wrapped object. The field’s type should be declared as the component interface so it can contain both concrete components and decorators. The base decorator delegates all operations to the wrapped object.
4. Concrete Decorators define extra behaviors that can be added to components dynamically. Concrete decorators override methods of the base decorator and execute their behavior either before or after calling the parent method.
5. The Client can wrap components in multiple layers of decorators, as long as it works with all objects via the component interface.

## Examples

### Simple decorator example

```go
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
```

### Pizza decorator example

```go
package main

import "fmt"

type IPizza interface {
 getPrice() int
}

// Mozzarella
type Mozzarella struct{}

func (p *Mozzarella) getPrice() int {
 return 10
}

// Tomato topping
type TomatoTopping struct {
 pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
 price := c.pizza.getPrice()
 return price + 4
}

// Stuffed edge
type StuffedEdge struct {
 pizza IPizza
}

func (c *StuffedEdge) getPrice() int {
 price := c.pizza.getPrice()
 return price + 4
}

func main() {
 pizza := &Mozzarella{}
 fmt.Println("Mozzarella price:", pizza.getPrice())

 pizzaWithTomato := &TomatoTopping{pizza}
 fmt.Println("Mozzarella with tomato topping price:", pizzaWithTomato.getPrice())

 pizzaWithTomatoAndStuffedEdge := &StuffedEdge{pizzaWithTomato}
 fmt.Println("Mozzarella with tomato topping and stuffed edge price:", pizzaWithTomatoAndStuffedEdge.getPrice())
}
```

### Multiple aggregation example (without decorator)

```go
package main

import "fmt"

type Aged interface {
 Age() int
 SetAge(age int)
}

type Bird struct {
 age int
}

func (b *Bird) Age() int       { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }

func (b *Bird) Fly() {
 if b.Age() >= 10 {
  fmt.Println("Flying!")
 }
}

type Lizard struct {
 age int
}

func (l *Lizard) Age() int       { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }

func (l *Lizard) Crawl() {
 if l.Age() < 10 {
  fmt.Println("Crawling!")
 }
}

type Dragon struct {
 bird   Bird
 lizard Lizard
}

func (d *Dragon) Age() int {
 return d.bird.Age()
}
func (d *Dragon) SetAge(age int) {
 d.bird.SetAge(age)
 d.lizard.SetAge(age)
}
func (d *Dragon) Fly() {
 d.bird.Fly()
}
func (d *Dragon) Crawl() {
 d.lizard.Crawl()
}

func NewDragon() *Dragon {
 return &Dragon{Bird{}, Lizard{}}
}

func main() {
 d := NewDragon()
 d.SetAge(10)
 d.Fly()
 d.Crawl()
}
```
