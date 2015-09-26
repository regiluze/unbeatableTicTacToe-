package main

import (
	"errors"
	_ "fmt"
)

const (
	NOUGHT          = "O"
	CROSS           = "X"
	EMPTY_SPACE     = "-"
	NOT_INLINE      = ""
	TOKENS_ON_BOARD = 9
)

type BoardSnapshot [3][3]string

type Position struct {
	Col  int
	Line int
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
			board.snapshot[row][col] = EMPTY_SPACE
		}
	}
}

func (board *TicTacToeBoard) IsOver() (bool, string) {
	if board.numberOfTokens == TOKENS_ON_BOARD {
		return true, NOT_INLINE
	}
	if tokenInLine := board.checkThreeInLine(); tokenInLine != NOT_INLINE {
		return true, tokenInLine
	}
	return false, NOT_INLINE
}

func (board *TicTacToeBoard) putToken(token string, position Position) (BoardSnapshot, error) {
	if board.placeIsFilled(position) {
		return BoardSnapshot{}, errors.New("Place already filled")
	}
	board.snapshot[position.Col][position.Line] = token
	board.numberOfTokens += 1
	return board.snapshot, nil
}

func (board *TicTacToeBoard) checkThreeInLine() string {

	if token, exists := board.inLineHorizontally(); exists {
		return token
	}
	if token, exists := board.inLineVertically(); exists {
		return token
	}
	if token, exists := board.inLineOnCrosses(); exists {
		return token
	}

	return NOT_INLINE
}

func (board *TicTacToeBoard) inLineOnCrosses() (string, bool) {
	cross := [3]string{board.snapshot[0][0], board.snapshot[1][1], board.snapshot[2][2]}
	if token := board.threeInLine(cross); token != NOT_INLINE {
		return token, true
	}
	cross = [3]string{board.snapshot[0][2], board.snapshot[1][1], board.snapshot[2][0]}
	if token := board.threeInLine(cross); token != NOT_INLINE {
		return token, true
	}
	return NOT_INLINE, false
}

func (board *TicTacToeBoard) inLineHorizontally() (string, bool) {
	for line := 0; line < 3; line++ {
		if token := board.threeInLine(board.snapshot[line]); token != NOT_INLINE {
			return token, true
		}
	}
	return NOT_INLINE, false
}

func (board *TicTacToeBoard) inLineVertically() (string, bool) {
	for column := 0; column < 3; column++ {
		tokensLine := [3]string{board.snapshot[0][column], board.snapshot[1][column], board.snapshot[2][column]}
		if token := board.threeInLine(tokensLine); token != NOT_INLINE {
			return token, true
		}
	}
	return NOT_INLINE, false
}

func (board *TicTacToeBoard) threeInLine(line [3]string) string {
	token := line[0]
	if token != EMPTY_SPACE && line[1] == token && line[2] == token {
		return token
	}
	return NOT_INLINE
}

func (board *TicTacToeBoard) placeIsFilled(position Position) bool {
	return board.snapshot[position.Col][position.Line] != EMPTY_SPACE
}
