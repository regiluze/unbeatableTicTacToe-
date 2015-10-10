package main

import (
	_ "fmt"
)

type putToken func(Position) (BoardSnapshot, error)

type Player interface {
	PutToken(BoardSnapshot) Position
}

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
		board = game.safeTokePut(game.board.PutCross, game.player1, board)
		if isOver, winner := game.board.IsOver(); isOver {
			return game.winnerMap[winner]
		}
		board = game.safeTokePut(game.board.PutNought, game.player2, board)
		if isOver, winner := game.board.IsOver(); isOver {
			return game.winnerMap[winner]
		}
	}
	return ""
}

func (game TicTacToeGame) safeTokePut(putFunction putToken, player Player, board BoardSnapshot) BoardSnapshot {
	position := player.PutToken(board)
	board, err := putFunction(position)
	if err != nil {
		game.safeTokePut(putFunction, player, board)
	}
	return board
}
