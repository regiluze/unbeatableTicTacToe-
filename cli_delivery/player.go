package main

import (
	"fmt"
	. "github.com/regiluze/unbeatableTicTacToe-"
)

type HumanPlayer struct {
	name string
}

func NewHumanPlayer(name string) Player {
	player := HumanPlayer{name}
	return player
}

func (player HumanPlayer) PutToken(snapshot BoardSnapshot) Position {
	snapshot.Print()
	fmt.Println("your turn ", player.name, ":")
	fmt.Println("select line : ")
	line := readFromUser([]int{0, 1, 2})
	fmt.Println("select column : ")
	column := readFromUser([]int{0, 1, 2})
	return Position{Line: line, Col: column}
}
