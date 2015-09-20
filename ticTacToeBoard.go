package main

import (
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

func (board *TicTacToeBoard) PutNought(position Position) BoardSnapshot {
	board.snapshot[position.X][position.Y] = NOUGHT
	return board.snapshot
}
