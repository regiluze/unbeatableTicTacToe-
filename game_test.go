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
		return Position{2, 2}
	}
	if looser.turn == 4 {
		return Position{1, 1}
	}
	if looser.turn == 5 {
		return Position{1, 0}
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

type AnotherLooserPlayer struct {
	turn int
}

func (looser *AnotherLooserPlayer) PutToken(board BoardSnapshot) Position {
	looser.turn += 1
	if looser.turn == 2 {
		return Position{0, 0}
	}
	if looser.turn == 3 {
		return Position{2, 0}
	}
	if looser.turn == 4 {
		return Position{2, 1}
	}
	return Position{1, 2}
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
		Context("the player1 puts 3 croosses on the first line", func() {
			It("returns player1 as winner", func() {
				winnerPlayer = &FirstLineFillerPlayer{}
				game = NewTicTacToeGame(winnerPlayer, looserPlayer)
				winner := game.Start()

				Expect(winner).Should(Equal("player 1"))

			})
		})
		Context("the players fill all board", func() {
			It("returns player1 as winner", func() {
				anotherLooserPlayer := &AnotherLooserPlayer{}
				game = NewTicTacToeGame(looserPlayer, anotherLooserPlayer)
				winner := game.Start()

				Expect(winner).Should(Equal("Draw"))

			})
		})
	})
})
