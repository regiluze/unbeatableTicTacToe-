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

	player1Selection := selectPlayer("Player 1")
	player2Selection := selectPlayer("Player 2")
	player1 := playerFactory(player1Selection, CROSS, "Player 1")
	player2 := playerFactory(player2Selection, NOUGHT, "Player 2")
	game := NewTicTacToeGame(player1, player2)
	initGame(player1Selection, player2Selection)
	winner := game.Start()
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

func playerFactory(player int, tokenType string, name string) Player {
	if player == HUMAN {
		return NewHumanPlayer(name)
	}
	rules := NewRules(tokenType)
	return NewUnbeateablePlayer(rules)
}

func initGame(player1, player2 int) {
	player := map[int]string{HUMAN: "Human", COMPUTER: "Computer"}
	fmt.Println("********* Let's start the game  ***************")
	fmt.Sprintln("*********", player[player1], "vs", player[player2], " ***************************")
	time.Sleep(time.Millisecond * 300)
	fmt.Println("***********************************************")
}
