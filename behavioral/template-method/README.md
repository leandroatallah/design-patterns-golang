# Template Method

> Template Method is a behavioral design pattern that defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure. (Refactoring Guru)

> The template method is a method in a superclass, usually an abstract superclass, and defines the skeleton of an operation in terms of a number of high-level steps. These steps are themselves implemented by additional helper methods in the same class as the template method. (Wikipedia)

## Examples

### Basic Template Method example

```go
package main

import "fmt"

type Game interface {
 Start()
 TakeTurn()
 HaveWinner() bool
 WinningPlayer() int
}

// Template method
func PlayGame(g Game) {
 g.Start()
 for !g.HaveWinner() {
  g.TakeTurn()
 }
 fmt.Printf("Player %d wins.\n", g.WinningPlayer())
}

type chess struct {
 turn, maxTurns, currentPlayer int
}

func (c *chess) Start() {
 fmt.Println("Starting a new game of chess.")
}
func (c *chess) TakeTurn() {
 c.turn++
 fmt.Printf("Turn %d take by player %d\n", c.turn, c.currentPlayer)
 c.currentPlayer = 1 - c.currentPlayer
}
func (c *chess) HaveWinner() bool {
 return c.turn == c.maxTurns
}
func (c *chess) WinningPlayer() int {
 return c.currentPlayer
}

func NewGameOfChess() Game {
 return &chess{1, 10, 0}
}

func main() {
 chess := NewGameOfChess()
 PlayGame(chess)
}
```

### Functional Template Method example

```go
package main

import "fmt"

// Template method
func PlayGame(
 start, takeTurn func(),
 haveWinner func() bool,
 winningPlayer func() int,
) {
 start()
 for !haveWinner() {
  takeTurn()
 }

 fmt.Printf("Player %d wins.\n", winningPlayer())
}

func main() {
 turn, maxTurns, currentPlayer := 1, 10, 0

 start := func() {
  fmt.Println("Starting a new game of chess.")
 }
 takeTurn := func() {
  turn++
  fmt.Printf("Turn %d take by player %d\n", turn, currentPlayer)
  currentPlayer = 1 - currentPlayer
 }
 haveWinner := func() bool {
  return turn == maxTurns
 }
 winningPlayer := func() int {
  return currentPlayer
 }

 PlayGame(start, takeTurn, haveWinner, winningPlayer)
}
```
