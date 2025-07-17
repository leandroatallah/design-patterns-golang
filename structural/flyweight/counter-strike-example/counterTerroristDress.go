package main

type CounterTerroristDress struct {
	color string
}

func (t *CounterTerroristDress) getColor() string {
	return t.color
}

func NewCounterTorroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}
