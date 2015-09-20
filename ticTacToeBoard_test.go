package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/regiluze/unbeatableTicTacToe-"
)

const ()

var _ = Describe("Tic Tac Toe game board specs", func() {
	var (
		board              *TicTacToeBoard
		irrelevantPosition Position
	)

	BeforeEach(func() {
		board = NewTicTacToeBoard()
		irrelevantPosition = Position{X: 0, Y: 0}
	})

	Describe("Filling the board with tokens", func() {
		Context("when adding a nought", func() {
			It("fills board position with a nought token", func() {
				boardSnapshot, _ := board.PutNought(irrelevantPosition)

				Expect(boardSnapshot[irrelevantPosition.X][irrelevantPosition.Y]).To(Equal(NOUGHT))
			})
			Context("when the board place is already filled", func() {
				It("returns an error", func() {
					board.PutNought(irrelevantPosition)
					_, err := board.PutNought(irrelevantPosition)

					Expect(err).NotTo(Equal(nil))
				})
			})
		})
		Context("when adding a cross", func() {
			It("fills board position with a cross token", func() {
				boardSnapshot, _ := board.PutCross(irrelevantPosition)

				Expect(boardSnapshot[irrelevantPosition.X][irrelevantPosition.Y]).To(Equal(CROSS))
			})
			Context("when the board place is already filled", func() {
				It("returns an error", func() {
					board.PutCross(irrelevantPosition)
					_, err := board.PutCross(irrelevantPosition)

					Expect(err).NotTo(Equal(nil))
				})
			})
		})
	})
})
