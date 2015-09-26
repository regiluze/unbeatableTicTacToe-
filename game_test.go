package main_test

import (
	_ "fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/regiluze/unbeatableTicTacToe-"
)

const ()

type LooserPlayer struct {
	turn int
}

func (looser *LooserPlayer) PutToken(board BoardSnapshot) Position {
	looser.turn += 1
	if looser.turn == 2 {
		return Position{0, 1}
	}
	if looser.turn == 3 {
		return Position{2, 1}
	}
	return Position{0, 2}

}

type FirstLineFillerPlayer struct {
	turn int
}

func (looser *FirstLineFillerPlayer) PutToken(board BoardSnapshot) Position {
	looser.turn += 1
	if looser.turn == 2 {
		return Position{1, 0}
	}
	if looser.turn == 3 {
		return Position{2, 0}
	}
	return Position{0, 0}
}

var _ = Describe("Tic Tac Toe game specs", func() {
	var (
		game         TicTacToeGame
		winnerPlayer Player
		looserPlayer Player
	)

	BeforeEach(func() {
		looserPlayer = &LooserPlayer{}
	})

	Describe("When the game starts", func() {
		Context("the player1 puts 3 croosses on first line", func() {
			It("returns player1 as winner", func() {
				winnerPlayer = &FirstLineFillerPlayer{}
				game = NewTicTacToeGame(winnerPlayer, looserPlayer)
				winner := game.Start()

				Expect(winner).Should(Equal("player 1"))

			})
		})
	})
})
