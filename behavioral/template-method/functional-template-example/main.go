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
