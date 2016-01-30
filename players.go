package unbeatable

import (
	_ "fmt"
)

const (
	INVALID_POSITION = -1
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
	return player.defaultRule(snapshot)
}

func (player UnbeatablePlayer) matchRuleToPut(snapshot BoardSnapshot) (bool, Position) {
	for _, checkFunc := range player.rules.All {
		if match, position := checkFunc(snapshot); match {
			return true, position
		}
	}
	return false, Position{}
}

func (player UnbeatablePlayer) defaultRule(snapshot BoardSnapshot) Position {
	if snapshot[1][1] == EMPTY_SPACE {
		return Position{1, 1}
	}
	return player.firstFreeSpacePosition(snapshot)
}

func (player UnbeatablePlayer) firstFreeSpacePosition(snapshot BoardSnapshot) Position {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if snapshot[row][column] == EMPTY_SPACE {
				return Position{row, column}
			}
		}
	}
	return Position{0, 0}
}

type RuleCheckFunc func(snapshot BoardSnapshot) (bool, Position)

type Rules struct {
	tokenType string
	All       []RuleCheckFunc
}

func NewRules(tokenType string) Rules {
	r := Rules{tokenType: tokenType}
	winnerCheckFuncs := []RuleCheckFunc{r.checkLinesToWin, r.checkColumnsToWin, r.checkFirstCrossLineToWin, r.checkSecondCrossLineToWin}
	safeCheckFuncs := []RuleCheckFunc{r.checkLinesToSave, r.checkColumnsToSave, r.checkFirstCrossLineToSave, r.checkSecondCrossLineToSave}
	allRules := append(winnerCheckFuncs, safeCheckFuncs...)

	r.All = append(allRules, r.initRule)
	return r
}

func (r Rules) checkLinesToWin(snapshot BoardSnapshot) (bool, Position) {
	return r.checkLines(snapshot, r.sameTokenType)
}

func (r Rules) checkLinesToSave(snapshot BoardSnapshot) (bool, Position) {
	return r.checkLines(snapshot, r.differentTokenType)
}

func (r Rules) checkLines(snapshot BoardSnapshot, matchFunc func(string) bool) (bool, Position) {
	for column, row := range snapshot {
		if match, position := r.checkLineToMatch(row, matchFunc); match {
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
		columnLine := BoardLine{snapshot[0][column], snapshot[1][column], snapshot[2][column]}
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
	crossLine := BoardLine{snapshot[0][0], snapshot[1][1], snapshot[2][2]}
	if match, position := r.checkLineToMatch(crossLine, matchFunc); match {
		return true, Position{position, position}
	}
	return false, Position{}
}

func (r Rules) checkSecondCrossLineToWin(snapshot BoardSnapshot) (bool, Position) {
	return r.checkSecondCrossLine(snapshot, r.sameTokenType)
}

func (r Rules) checkSecondCrossLineToSave(snapshot BoardSnapshot) (bool, Position) {
	return r.checkSecondCrossLine(snapshot, r.differentTokenType)
}

func (r Rules) checkSecondCrossLine(snapshot BoardSnapshot, matchFunc func(string) bool) (bool, Position) {
	crossLine := BoardLine{snapshot[0][2], snapshot[1][1], snapshot[2][0]}
	if match, position := r.checkLineToMatch(crossLine, matchFunc); match {
		return true, Position{position, 2 - position}
	}
	return false, Position{}
}

func (r Rules) initRule(snapshot BoardSnapshot) (bool, Position) {
	for _, row := range snapshot {
		if matchNumber := r.filter(row, r.emptySpaces); matchNumber != 3 {
			return false, Position{0, 0}
		}
	}
	return true, Position{1, 1}
}

func (r Rules) checkLineToMatch(row BoardLine, matchFunc func(string) bool) (bool, int) {
	if r.filter(row, matchFunc) == 2 {
		position := r.getEmtySpacePosition(row)
		return position != INVALID_POSITION, position
	}
	return false, 0
}

func (r Rules) filter(row BoardLine, matchFunc func(string) bool) int {
	matchNumber := 0
	for _, token := range row {
		if matchFunc(token) {
			matchNumber += 1
		}
	}
	return matchNumber
}

func (r Rules) getEmtySpacePosition(row BoardLine) int {
	position := INVALID_POSITION
	for i, token := range row {
		if token == EMPTY_SPACE {
			position = i
		}
	}
	return position
}

func (u Rules) emptySpaces(tokenType string) bool {
	return tokenType == EMPTY_SPACE
}

func (u Rules) sameTokenType(tokenType string) bool {
	return tokenType == u.tokenType
}

func (r Rules) differentTokenType(tokenType string) bool {
	return tokenType != r.tokenType && tokenType != EMPTY_SPACE
}
