package main_test

import (
	_ "fmt"
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
		Context("when the board is completly full", func() {
			It("returns true and empty string", func() {

				fillAllBoard(board)

				result, winner := board.IsOver()

				Expect(result).Should(BeTrue())
				Expect(winner).Should(BeEmpty())

			})
		})

		Context("when the same type three tokens are in line", func() {
			Context("when three crosses are on the first line", func() {
				It("returns true and a cross", func() {

					board.PutCross(Position{0, 0})
					board.PutCross(Position{0, 1})
					board.PutCross(Position{0, 2})

					result, winner := board.IsOver()

					Expect(result).Should(BeTrue())
					Expect(winner).Should(Equal(CROSS))

				})
			})
			Context("when there three crosses are on the second line", func() {
				It("returns true and a cross", func() {

					board.PutCross(Position{1, 0})
					board.PutCross(Position{1, 1})
					board.PutCross(Position{1, 2})

					result, winner := board.IsOver()

					Expect(result).Should(BeTrue())
					Expect(winner).Should(Equal(CROSS))

				})
			})
			Context("when three noughts are on the third line", func() {
				It("returns true and a hought", func() {

					board.PutNought(Position{2, 0})
					board.PutNought(Position{2, 1})
					board.PutNought(Position{2, 2})

					result, winner := board.IsOver()

					Expect(result).Should(BeTrue())
					Expect(winner).Should(Equal(NOUGHT))

				})
			})
		})
		Context("when the same type three tokens are on column line", func() {
			Context("when three crosses are on the first column", func() {
				It("returns true and a cross", func() {

					board.PutCross(Position{0, 0})
					board.PutCross(Position{1, 0})
					board.PutCross(Position{2, 0})

					result, winner := board.IsOver()

					Expect(result).Should(BeTrue())
					Expect(winner).Should(Equal(CROSS))

				})
			})
			Context("when three croosses are on the second column", func() {
				It("returns true and a cross", func() {

					board.PutCross(Position{0, 1})
					board.PutCross(Position{1, 1})
					board.PutCross(Position{2, 1})

					result, winner := board.IsOver()

					Expect(result).Should(BeTrue())
					Expect(winner).Should(Equal(CROSS))

				})
			})
		})
		Context("when the same type three tokens  are on cross lines", func() {
			Context("when three noughts are on first cross", func() {
				It("returns true and a nought", func() {

					board.PutNought(Position{0, 0})
					board.PutNought(Position{1, 1})
					board.PutNought(Position{2, 2})

					result, winner := board.IsOver()

					Expect(result).Should(BeTrue())
					Expect(winner).Should(Equal(NOUGHT))

				})
			})
			Context("when three crosses are on second cross", func() {
				It("returns true and a nought", func() {

					board.PutCross(Position{0, 2})
					board.PutCross(Position{1, 1})
					board.PutCross(Position{2, 0})

					result, winner := board.IsOver()

					Expect(result).Should(BeTrue())
					Expect(winner).Should(Equal(CROSS))

				})
			})
		})
	})
})

func fillAllBoard(board *TicTacToeBoard) {
	token := CROSS
	for row := 0; row < 3; row++ {
		if row == 1 {
			token = CROSS
		}
		for col := 0; col < 3; col++ {
			if token == CROSS {
				board.PutNought(Position{row, col})
				token = NOUGHT
			} else {
				board.PutCross(Position{row, col})
				token = CROSS
			}
		}
	}
}
