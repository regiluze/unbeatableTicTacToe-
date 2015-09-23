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
	winner         string
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
	board.winner = ""
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
	if winner := board.checkWinner(); winner != "" {
		return true, winner
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

func (board *TicTacToeBoard) checkWinner() string {
	for line := 0; line < 4; line++ {
		if token := board.threeInLine(board.snapshot[line]); token != "" {
			return token
		}
	}

	return ""
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
