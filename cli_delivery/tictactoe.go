package main

import (
	"fmt"
	. "github.com/regiluze/unbeatableTicTacToe-"
	"time"
)

const (
	HUMAN    = 1
	COMPUTER = 2
)

func main() {
	welcome()
	player1, player2 := gamePlayers()
	game := NewTicTacToeGame(player1, player2)
	winner, finishBoard := game.Start()
	finishBoard.Print()
	fmt.Println("the winner is: ", winner)

}

func welcome() {
	fmt.Println("********* welcome to tic tac toe game ***************")
	for i := 0; i < 4; i++ {
		time.Sleep(time.Millisecond * 300)
		fmt.Print("*")
	}
	fmt.Println("***********************************************")
}

func gamePlayers() (Player, Player) {
	player1Selection := selectPlayer("Player 1")
	player2Selection := selectPlayer("Player 2")
	player1 := playerFactory(player1Selection, CROSS, "Player 1")
	player2 := playerFactory(player2Selection, NOUGHT, "Player 2")
	displayPlayers(player1Selection, player2Selection)
	return player1, player2
}

func selectPlayer(name string) int {
	fmt.Println("Select ", name)
	fmt.Println("1 -> Human | 2 -> Computer [1, 2]")
	return readFromUser([]int{HUMAN, COMPUTER})
}

func playerFactory(player int, tokenType string, name string) Player {
	if player == HUMAN {
		return NewHumanPlayer(name)
	}
	rules := NewRules(tokenType)
	return NewUnbeateablePlayer(rules)
}

func displayPlayers(player1, player2 int) {
	player := map[int]string{HUMAN: "Human", COMPUTER: "Computer"}
	fmt.Println("********* Let's start the game  ***************")
	fmt.Println("*********", player[player1], "vs", player[player2])
	time.Sleep(time.Millisecond * 300)
	fmt.Println("***********************************************")
}
