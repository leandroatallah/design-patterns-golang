package main

type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func NewTorroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}
