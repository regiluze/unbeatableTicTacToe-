package main

import (
	"fmt"
)

type UnbeatablePlayer struct {
	tokenType string
}

func NewUnbeateablePlayer(tokenType string) UnbeatablePlayer {
	player := UnbeatablePlayer{tokenType}
	return player

}

func (u UnbeatablePlayer) PutToken(snapshot BoardSnapshot) Position {
	for column, line := range snapshot {
		if match, position := u.checkLineToWin(line); match {
			return Position{column, position}
		}
		cc := [3]string{snapshot[0][column], snapshot[1][column], snapshot[2][column]}
		if match, position := u.checkLineToWin(cc); match {
			return Position{position, column}
		}
	}

	return Position{0, 0}

}

func (u UnbeatablePlayer) checkLineToWin(line [3]string) (bool, int) {
	fmt.Println(" line", line)
	position := -1
	match := 0
	for i, token := range line {
		if token != EMPTY_SPACE && token == u.tokenType {
			match += 1
		} else if token == "" {
			position = i
		}
	}
	if match == 2 && position != -1 {
		return true, position
	}
	return false, 0
}
