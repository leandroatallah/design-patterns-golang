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
