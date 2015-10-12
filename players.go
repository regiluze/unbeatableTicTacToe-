package main

import (
	_ "fmt"
)

type WinnerCheckFunc func(snapshot BoardSnapshot) (bool, Position)

type UnbeatablePlayer struct {
	tokenType string
}

func NewUnbeateablePlayer(tokenType string) UnbeatablePlayer {
	player := UnbeatablePlayer{tokenType}
	return player

}

func (player UnbeatablePlayer) PutToken(snapshot BoardSnapshot) Position {
	if match, position := player.matchWinnerPut(snapshot); match {
		return position
	}

	return Position{0, 0}

}

func (player UnbeatablePlayer) matchWinnerPut(snapshot BoardSnapshot) (bool, Position) {
	winnerCheckFuncs := []WinnerCheckFunc{player.checkLines, player.checkColumns, player.checkFirstCrossLine, player.checkSecondCrossLine}
	for _, checkFunc := range winnerCheckFuncs {
		if match, position := checkFunc(snapshot); match {
			return true, position
		}
	}
	return false, Position{}

}

func (player UnbeatablePlayer) checkLines(snapshot BoardSnapshot) (bool, Position) {
	for column, line := range snapshot {
		if match, position := player.checkLineToWin(line); match {
			return true, Position{column, position}
		}
	}
	return false, Position{}
}

func (player UnbeatablePlayer) checkColumns(snapshot BoardSnapshot) (bool, Position) {
	for column, _ := range snapshot {
		columnLine := [3]string{snapshot[0][column], snapshot[1][column], snapshot[2][column]}
		if match, position := player.checkLineToWin(columnLine); match {
			return true, Position{position, column}
		}
	}
	return false, Position{}
}

func (player UnbeatablePlayer) checkFirstCrossLine(snapshot BoardSnapshot) (bool, Position) {
	crossLine := [3]string{snapshot[0][0], snapshot[1][1], snapshot[2][2]}
	if match, position := player.checkLineToWin(crossLine); match {
		return true, Position{position, position}
	}
	return false, Position{}
}

func (player UnbeatablePlayer) checkSecondCrossLine(snapshot BoardSnapshot) (bool, Position) {
	crossLine := [3]string{snapshot[0][2], snapshot[1][1], snapshot[2][0]}
	if match, position := player.checkLineToWin(crossLine); match {
		return true, Position{position, 2 - position}
	}
	return false, Position{}
}

func (u UnbeatablePlayer) checkLineToWin(line [3]string) (bool, int) {
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
