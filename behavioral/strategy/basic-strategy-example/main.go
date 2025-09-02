package main

import "fmt"

type RouteContext struct {
	strategy RouteStrategy
}

func (c *RouteContext) SetStrategy(s RouteStrategy) {
	c.strategy = s
}

func (c *RouteContext) GetDurationInMinutes() int {
	return c.strategy.GetDurationInMinutes()
}

// RouteStrategy
type RouteStrategy interface {
	GetDurationInMinutes() int
}

type CarRouteStrategy struct{}

func (s *CarRouteStrategy) GetDurationInMinutes() int {
	return 10
}

type WalkRouteStrategy struct{}

func (s *WalkRouteStrategy) GetDurationInMinutes() int {
	return 25
}

func main() {
	carStrategy := &CarRouteStrategy{}
	context := &RouteContext{}
	context.SetStrategy(carStrategy)
	fmt.Println("Card route time (in minutes): ", context.strategy.GetDurationInMinutes())

	walkStrategy := &WalkRouteStrategy{}
	context.SetStrategy(walkStrategy)
	fmt.Println("Walking route time (in minutes): ", context.strategy.GetDurationInMinutes())
}
