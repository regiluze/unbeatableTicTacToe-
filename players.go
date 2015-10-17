package main

import (
	_ "fmt"
)

type RuleCheckFunc func(snapshot BoardSnapshot) (bool, Position)

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
	winnerCheckFuncs := []RuleCheckFunc{player.checkLinesToWin, player.checkColumns, player.checkFirstCrossLine, player.checkSecondCrossLine}
	safeCheckFuncs := []RuleCheckFunc{player.checkLinesToSave, player.checkSaveColumns}
	rules := append(winnerCheckFuncs, safeCheckFuncs...)
	for _, checkFunc := range rules {
		if match, position := checkFunc(snapshot); match {
			return true, position
		}
	}
	return false, Position{}

}

func (player UnbeatablePlayer) checkLinesToWin(snapshot BoardSnapshot) (bool, Position) {
	return player.checkLines(snapshot, player.sameTokenType)
}

func (player UnbeatablePlayer) checkLinesToSave(snapshot BoardSnapshot) (bool, Position) {
	return player.checkLines(snapshot, player.differentTokenType)
}

func (player UnbeatablePlayer) checkLines(snapshot BoardSnapshot, matchFunc func(string) bool) (bool, Position) {
	for column, line := range snapshot {
		if match, position := player.checkLineToMatch(line, matchFunc); match {
			return true, Position{column, position}
		}
	}
	return false, Position{}
}

func (player UnbeatablePlayer) checkColumns(snapshot BoardSnapshot) (bool, Position) {
	for column, _ := range snapshot {
		columnLine := [3]string{snapshot[0][column], snapshot[1][column], snapshot[2][column]}
		if match, position := player.checkLineToMatch(columnLine, player.sameTokenType); match {
			return true, Position{position, column}
		}
	}
	return false, Position{}
}

func (player UnbeatablePlayer) checkSaveColumns(snapshot BoardSnapshot) (bool, Position) {
	for column, _ := range snapshot {
		columnLine := [3]string{snapshot[0][column], snapshot[1][column], snapshot[2][column]}
		if match, position := player.checkLineToMatch(columnLine, player.differentTokenType); match {
			return true, Position{position, column}
		}
	}
	return false, Position{}
}

func (player UnbeatablePlayer) checkFirstCrossLine(snapshot BoardSnapshot) (bool, Position) {
	crossLine := [3]string{snapshot[0][0], snapshot[1][1], snapshot[2][2]}
	if match, position := player.checkLineToMatch(crossLine, player.sameTokenType); match {
		return true, Position{position, position}
	}
	return false, Position{}
}

func (player UnbeatablePlayer) checkSecondCrossLine(snapshot BoardSnapshot) (bool, Position) {
	crossLine := [3]string{snapshot[0][2], snapshot[1][1], snapshot[2][0]}
	if match, position := player.checkLineToMatch(crossLine, player.sameTokenType); match {
		return true, Position{position, 2 - position}
	}
	return false, Position{}
}

func (player UnbeatablePlayer) checkLineToMatch(line [3]string, matchFunc func(string) bool) (bool, int) {
	if player.filter(line, matchFunc) == 2 {
		position := player.getEmtySpacePosition(line)
		return position != -1, position
	}
	return false, 0
}

func (player UnbeatablePlayer) filter(line [3]string, matchFunc func(string) bool) int {
	matchNumber := 0
	for _, token := range line {
		if matchFunc(token) {
			matchNumber += 1
		}
	}
	return matchNumber
}

func (player UnbeatablePlayer) getEmtySpacePosition(line [3]string) int {
	position := -1
	for i, token := range line {
		if token == "" {
			position = i
		}
	}
	return position
}

func (u UnbeatablePlayer) sameTokenType(tokenType string) bool {
	return tokenType == u.tokenType
}

func (player UnbeatablePlayer) differentTokenType(tokenType string) bool {
	return tokenType != player.tokenType && tokenType != ""
}
