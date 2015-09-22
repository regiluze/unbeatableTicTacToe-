package main_test

import (
	"fmt"
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

				Expect(boardSnapshot[irrelevantPosition.X][irrelevantPosition.Y]).Should(Equal(NOUGHT))
			})
			Context("when the board place is already filled", func() {
				It("returns an error", func() {
					board.PutNought(irrelevantPosition)
					_, err := board.PutNought(irrelevantPosition)

					Expect(err).ShouldNot(BeNil())
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

					Expect(err).ShouldNot(Equal(nil))
				})
			})
		})
	})
	Context("when reset the board", func() {
		It("cleans all tokens from the board", func() {
			board.PutNought(irrelevantPosition)

			board.Reset()

			_, err := board.PutNought(irrelevantPosition)
			Expect(err).Should(BeNil())
		})
	})
	Describe("the game is over", func() {
		Context("when the board is completly full and call isOver method", func() {
			It("returns true and empty string", func() {

				fillTheBoard(board)

				result, winner := board.IsOver()

				Expect(result).Should(BeTrue())
				Expect(winner).Should(BeEmpty())

			})
		})
	})
})

func fillTheBoard(board *TicTacToeBoard) {
	token := CROSS
	var snapshot BoardSnapshot
	for row := 0; row < 3; row++ {
		if row == 1 {
			token = CROSS
		}
		for col := 0; col < 3; col++ {
			if token == CROSS {
				snapshot, _ = board.PutNought(Position{row, col})
				token = NOUGHT
			} else {
				snapshot, _ = board.PutCross(Position{row, col})
				token = CROSS
			}
		}
	}
	fmt.Printf("%q", snapshot)

}
