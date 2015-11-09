package unbeatable

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

func (game TicTacToeGame) Start() (string, BoardSnapshot) {
	gameResult := game.run()
	return game.winnerMap[gameResult], game.board.Snapshot
}

func (game TicTacToeGame) run() string {
	var gameSnapshot BoardSnapshot
	for {
		gameSnapshot = game.safeTokePut(game.board.PutCross, game.player1, gameSnapshot)
		if isOver, winner := game.board.IsOver(); isOver {
			return winner
		}
		gameSnapshot = game.safeTokePut(game.board.PutNought, game.player2, gameSnapshot)
		if isOver, winner := game.board.IsOver(); isOver {
			return winner
		}
	}
}

func (game TicTacToeGame) safeTokePut(putFunction putToken, player Player, gameSnapshot BoardSnapshot) BoardSnapshot {
	position := player.PutToken(gameSnapshot)
	newGameSnapshot, err := putFunction(position)
	if err != nil {
		return game.safeTokePut(putFunction, player, gameSnapshot)
	}
	return newGameSnapshot
}
