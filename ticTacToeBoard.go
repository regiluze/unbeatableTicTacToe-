package main

import (
	"errors"
	_ "fmt"
)

const (
	NOUGHT          = "O"
	CROSS           = "X"
	TOKENS_ON_BOARD = 9
)

type BoardSnapshot [3][3]string

type Position struct {
	X int
	Y int
}

type TicTacToeBoard struct {
	snapshot       BoardSnapshot
	numberOfTokens int
}

func NewTicTacToeBoard() *TicTacToeBoard {

	board := &TicTacToeBoard{}
	board.Reset()
	return board

}

func (board *TicTacToeBoard) PutNought(position Position) (BoardSnapshot, error) {
	return board.putToken(NOUGHT, position)
}

func (board *TicTacToeBoard) PutCross(position Position) (BoardSnapshot, error) {
	return board.putToken(CROSS, position)
}

func (board *TicTacToeBoard) Reset() {
	board.numberOfTokens = 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			board.snapshot[row][col] = ""
		}
	}
}

func (board *TicTacToeBoard) IsOver() (bool, string) {
	if board.numberOfTokens == TOKENS_ON_BOARD {
		return true, ""
	}
	if tokenInLine := board.checkThreeInLine(); tokenInLine != "" {
		return true, tokenInLine
	}
	return false, ""
}

func (board *TicTacToeBoard) putToken(token string, position Position) (BoardSnapshot, error) {
	if board.placeIsFilled(position) {
		return BoardSnapshot{}, errors.New("Place already filled")
	}
	board.snapshot[position.X][position.Y] = token
	board.numberOfTokens += 1
	return board.snapshot, nil
}

func (board *TicTacToeBoard) checkThreeInLine() string {

	if winner, match := board.winnerHorizontally(); match {
		return winner
	}
	if winnerToken, match := board.winnerVertically(); match {
		return winnerToken
	}
	if winnerToken, match := board.winnerOnCrosses(); match {
		return winnerToken
	}

	return ""
}

func (board *TicTacToeBoard) winnerOnCrosses() (string, bool) {
	cross := [3]string{board.snapshot[0][0], board.snapshot[1][1], board.snapshot[2][2]}
	if token := board.threeInLine(cross); token != "" {
		return token, true
	}
	cross = [3]string{board.snapshot[0][2], board.snapshot[1][1], board.snapshot[2][0]}
	if token := board.threeInLine(cross); token != "" {
		return token, true
	}
	return "", false
}

func (board *TicTacToeBoard) winnerHorizontally() (string, bool) {
	for line := 0; line < 3; line++ {
		if token := board.threeInLine(board.snapshot[line]); token != "" {
			return token, true
		}
	}
	return "", false
}

func (board *TicTacToeBoard) winnerVertically() (string, bool) {
	for column := 0; column < 3; column++ {
		tokensLine := [3]string{board.snapshot[0][column], board.snapshot[1][column], board.snapshot[2][column]}
		if token := board.threeInLine(tokensLine); token != "" {
			return token, true
		}
	}
	return "", false
}

func (board *TicTacToeBoard) threeInLine(line [3]string) string {
	token := line[0]
	if token != "" && line[1] == token && line[2] == token {
		return token
	}
	return ""
}

func (board *TicTacToeBoard) placeIsFilled(position Position) bool {
	return len(board.snapshot[position.X][position.Y]) > 0
}
