package main

import (
	_ "fmt"
)

type UnbeatablePlayer struct {
	rules Rules
}

func NewUnbeateablePlayer(rules Rules) UnbeatablePlayer {
	player := UnbeatablePlayer{rules}
	return player
}

func (player UnbeatablePlayer) PutToken(snapshot BoardSnapshot) Position {
	if match, position := player.matchRuleToPut(snapshot); match {
		return position
	}
	return Position{0, 0}
}

func (player UnbeatablePlayer) matchRuleToPut(snapshot BoardSnapshot) (bool, Position) {
	for _, checkFunc := range player.rules.All {
		if match, position := checkFunc(snapshot); match {
			return true, position
		}
	}
	return false, Position{}

}

type RuleCheckFunc func(snapshot BoardSnapshot) (bool, Position)

type Rules struct {
	tokenType string
	All       []RuleCheckFunc
}

func NewRules(tokenType string) Rules {
	r := Rules{tokenType: tokenType}
	winnerCheckFuncs := []RuleCheckFunc{r.checkLinesToWin, r.checkColumnsToWin, r.checkFirstCrossLineToWin, r.checkSecondCrossLine}
	safeCheckFuncs := []RuleCheckFunc{r.checkLinesToSave, r.checkColumnsToSave, r.checkFirstCrossLineToSave}
	allRules := append(winnerCheckFuncs, safeCheckFuncs...)
	r.All = allRules
	return r
}

func (r Rules) checkLinesToWin(snapshot BoardSnapshot) (bool, Position) {
	return r.checkLines(snapshot, r.sameTokenType)
}

func (r Rules) checkLinesToSave(snapshot BoardSnapshot) (bool, Position) {
	return r.checkLines(snapshot, r.differentTokenType)
}

func (r Rules) checkLines(snapshot BoardSnapshot, matchFunc func(string) bool) (bool, Position) {
	for column, line := range snapshot {
		if match, position := r.checkLineToMatch(line, matchFunc); match {
			return true, Position{column, position}
		}
	}
	return false, Position{}
}

func (r Rules) checkColumnsToWin(snapshot BoardSnapshot) (bool, Position) {
	return r.checkColumns(snapshot, r.sameTokenType)
}

func (r Rules) checkColumnsToSave(snapshot BoardSnapshot) (bool, Position) {
	return r.checkColumns(snapshot, r.differentTokenType)
}

func (r Rules) checkColumns(snapshot BoardSnapshot, matchFunc func(string) bool) (bool, Position) {
	for column, _ := range snapshot {
		columnLine := [3]string{snapshot[0][column], snapshot[1][column], snapshot[2][column]}
		if match, position := r.checkLineToMatch(columnLine, matchFunc); match {
			return true, Position{position, column}
		}
	}
	return false, Position{}
}

func (r Rules) checkFirstCrossLineToWin(snapshot BoardSnapshot) (bool, Position) {
	return r.checkFirstCrossLine(snapshot, r.sameTokenType)
}

func (r Rules) checkFirstCrossLineToSave(snapshot BoardSnapshot) (bool, Position) {
	return r.checkFirstCrossLine(snapshot, r.differentTokenType)
}

func (r Rules) checkFirstCrossLine(snapshot BoardSnapshot, matchFunc func(string) bool) (bool, Position) {
	crossLine := [3]string{snapshot[0][0], snapshot[1][1], snapshot[2][2]}
	if match, position := r.checkLineToMatch(crossLine, matchFunc); match {
		return true, Position{position, position}
	}
	return false, Position{}
}

func (r Rules) checkSecondCrossLine(snapshot BoardSnapshot) (bool, Position) {
	crossLine := [3]string{snapshot[0][2], snapshot[1][1], snapshot[2][0]}
	if match, position := r.checkLineToMatch(crossLine, r.sameTokenType); match {
		return true, Position{position, 2 - position}
	}
	return false, Position{}
}

func (r Rules) checkLineToMatch(line [3]string, matchFunc func(string) bool) (bool, int) {
	if r.filter(line, matchFunc) == 2 {
		position := r.getEmtySpacePosition(line)
		return position != -1, position
	}
	return false, 0
}

func (r Rules) filter(line [3]string, matchFunc func(string) bool) int {
	matchNumber := 0
	for _, token := range line {
		if matchFunc(token) {
			matchNumber += 1
		}
	}
	return matchNumber
}

func (r Rules) getEmtySpacePosition(line [3]string) int {
	position := -1
	for i, token := range line {
		if token == "" {
			position = i
		}
	}
	return position
}

func (u Rules) sameTokenType(tokenType string) bool {
	return tokenType == u.tokenType
}

func (r Rules) differentTokenType(tokenType string) bool {
	return tokenType != r.tokenType && tokenType != ""
}
