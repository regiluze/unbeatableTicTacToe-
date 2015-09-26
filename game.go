package main

import (
	"fmt"
)

type Player interface {
	PutToken(BoardSnapshot) Position
}

const ()

type TicTacToeGame struct {
	player1   Player
	player2   Player
	board     *TicTacToeBoard
	winnerMap map[string]string
}

func NewTicTacToeGame(player1 Player, player2 Player) TicTacToeGame {
	board := NewTicTacToeBoard()
	winnerMap := map[string]string{
		CROSS:  "player 1",
		NOUGHT: "player 2",
		"":     "Draw",
	}
	game := TicTacToeGame{player1, player2, board, winnerMap}
	return game

}

func (game TicTacToeGame) Start() string {
	var board BoardSnapshot
	for {
		player1Turn := game.player1.PutToken(board)
		board, _ = game.board.PutCross(player1Turn)
		player2Turn := game.player2.PutToken(board)
		board, _ = game.board.PutNought(player2Turn)
		if isOver, winner := game.board.IsOver(); isOver {
			fmt.Println("EGI >>> winner", winner)
			return game.winnerMap[winner]
		}
	}

	return ""

}
