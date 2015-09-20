package main

import (
	"errors"
	_ "fmt"
)

const (
	NOUGHT = "O"
	CROSS  = "X"
)

type BoardSnapshot [3][3]string

type Position struct {
	X int
	Y int
}

type TicTacToeBoard struct {
	snapshot BoardSnapshot
}

func NewTicTacToeBoard() *TicTacToeBoard {

	board := &TicTacToeBoard{}
	return board

}

func (board *TicTacToeBoard) PutNought(position Position) (BoardSnapshot, error) {
	board.snapshot[position.X][position.Y] = NOUGHT
	return board.snapshot, nil
}

func (board *TicTacToeBoard) PutCross(position Position) (BoardSnapshot, error) {
	if len(board.snapshot[position.X][position.Y]) > 0 {
		return BoardSnapshot{}, errors.New("Place already filled")
	}
	board.snapshot[position.X][position.Y] = CROSS
	return board.snapshot, nil
}
