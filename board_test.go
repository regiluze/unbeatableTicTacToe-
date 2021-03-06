package unbeatable_test

import (
	_ "fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/regiluze/unbeatableTicTacToe-"
)

const ()

var _ = Describe("Unbeatable TicTacToe game board specs", func() {
	var (
		board              *TokensBoard
		irrelevantPosition Position
	)

	BeforeEach(func() {
		board = NewTokensBoard()
		irrelevantPosition = Position{Col: 0, Line: 0}
	})

	Describe("Filling the board with tokens", func() {
		Context("when adding a nought", func() {
			It("fills board position with a nought token", func() {
				boardSnapshot, _ := board.PutNought(irrelevantPosition)

				Expect(boardSnapshot[irrelevantPosition.Col][irrelevantPosition.Line]).Should(Equal(NOUGHT))
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

				Expect(boardSnapshot[irrelevantPosition.Col][irrelevantPosition.Line]).To(Equal(CROSS))
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
	Context("when reseting the board", func() {
		It("cleans all tokens from the board", func() {
			board.PutNought(irrelevantPosition)

			board.Reset()

			_, err := board.PutNought(irrelevantPosition)
			Expect(err).Should(BeNil())
		})
	})

	Describe("checking the board status", func() {
		Context("when the board is completly full", func() {
			It("returns true", func() {
				fillAllBoard(board)

				result := board.IsFull()

				Expect(result).Should(BeTrue())
			})
		})

		Context("when there are three tokens in line", func() {
			Context("when three crosses are on the first line", func() {
				It("returns true and a cross", func() {

					board.PutCross(Position{0, 0})
					board.PutCross(Position{0, 1})
					board.PutCross(Position{0, 2})

					result, winner := board.IsAnyInLine()

					Expect(result).Should(BeTrue())
					Expect(winner.Type).Should(Equal(CROSS))

				})
			})
			Context("when three crosses are on the second line", func() {
				It("returns true and a cross", func() {

					board.PutCross(Position{1, 0})
					board.PutCross(Position{1, 1})
					board.PutCross(Position{1, 2})

					result, winner := board.IsAnyInLine()

					Expect(result).Should(BeTrue())
					Expect(winner.Type).Should(Equal(CROSS))

				})
			})
			Context("when three noughts are on the third line", func() {
				It("returns true and a hought", func() {
					board.PutNought(Position{2, 0})
					board.PutNought(Position{2, 1})
					board.PutNought(Position{2, 2})

					result, winner := board.IsAnyInLine()

					Expect(result).Should(BeTrue())
					Expect(winner.Type).Should(Equal(NOUGHT))
				})
			})
		})
		Context("when the same type three tokens are on column line", func() {
			Context("when three crosses are on the first column", func() {
				It("returns true and a cross", func() {
					board.PutCross(Position{0, 0})
					board.PutCross(Position{1, 0})
					board.PutCross(Position{2, 0})

					result, winner := board.IsAnyInLine()

					Expect(result).Should(BeTrue())
					Expect(winner.Type).Should(Equal(CROSS))
				})
			})
			Context("when three croosses are on the second column", func() {
				It("returns true and a cross", func() {
					board.PutCross(Position{0, 1})
					board.PutCross(Position{1, 1})
					board.PutCross(Position{2, 1})

					result, winner := board.IsAnyInLine()

					Expect(result).Should(BeTrue())
					Expect(winner.Type).Should(Equal(CROSS))
				})
			})
		})
		Context("when the same type three tokens  are on cross lines", func() {
			Context("when three noughts are on first cross", func() {
				It("returns true and a nought", func() {
					board.PutNought(Position{0, 0})
					board.PutNought(Position{1, 1})
					board.PutNought(Position{2, 2})

					result, winner := board.IsAnyInLine()

					Expect(result).Should(BeTrue())
					Expect(winner.Type).Should(Equal(NOUGHT))
				})
			})
			Context("when three crosses are on second cross", func() {
				It("returns true and a nought", func() {

					board.PutCross(Position{0, 2})
					board.PutCross(Position{1, 1})
					board.PutCross(Position{2, 0})

					result, winner := board.IsAnyInLine()

					Expect(result).Should(BeTrue())
					Expect(winner.Type).Should(Equal(CROSS))

				})
			})
		})
	})
})

func fillAllBoard(board *TokensBoard) {
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
