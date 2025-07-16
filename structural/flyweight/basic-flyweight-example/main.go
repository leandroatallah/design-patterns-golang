// flyweight.go
package main

import "fmt"

// Dados compartilhados (flyweight)
type EnemyType struct {
	Name   string
	Sprite string // caminho do sprite
	Speed  float64
}

// Dados únicos de cada inimigo
type Enemy struct {
	Type *EnemyType // referência ao flyweight
	X, Y float64
	HP   int
}

func main() {
	// Flyweights
	zombieType := &EnemyType{Name: "Zombie", Sprite: "zombie.png", Speed: 1.0}
	batType := &EnemyType{Name: "Bat", Sprite: "bat.png", Speed: 2.5}

	enemies := []*Enemy{
		{Type: zombieType, X: 10, Y: 20, HP: 100},
		{Type: zombieType, X: 15, Y: 25, HP: 100},
		{Type: batType, X: 5, Y: 8, HP: 50},
	}

	for _, e := range enemies {
		fmt.Printf("%s at (%.0f,%.0f) HP:%d\n", e.Type.Name, e.X, e.Y, e.HP)
	}
}
