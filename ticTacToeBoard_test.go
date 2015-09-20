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
				boardSnapshot := board.PutNought(irrelevantPosition)

				Expect(boardSnapshot[irrelevantPosition.X][irrelevantPosition.Y]).To(Equal(NOUGHT))
			})
		})
		Context("when adding a cross", func() {
			It("fills board position with a cross token", func() {
				boardSnapshot := board.PutCross(irrelevantPosition)

				Expect(boardSnapshot[irrelevantPosition.X][irrelevantPosition.Y]).To(Equal(CROSS))
			})
		})
	})
})
