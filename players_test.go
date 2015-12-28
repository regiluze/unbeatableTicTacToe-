package unbeatable_test

import (
	_ "fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/regiluze/unbeatableTicTacToe-"
)

const ()

var _ = Describe("Tic Tac Toe unbeatable player specs", func() {
	var (
		player   UnbeatablePlayer
		snapshot BoardSnapshot
	)

	BeforeEach(func() {
		rules := NewRules(NOUGHT)
		player = NewUnbeateablePlayer(rules)
		snapshot = BoardSnapshot{}
		snapshot.Reset()
		snapshot.Print()
	})

	Describe("when the player is trying to win the game", func() {
		Context("when two noughts are on the first line", func() {
			It("puts a nought on the first line free space", func() {
				snapshot[0][0] = NOUGHT
				snapshot[0][1] = NOUGHT

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{0, 2}))

			})
		})
		Context("when two noughts are on the second line", func() {
			It("puts a nought on the second line free space", func() {
				snapshot[1][0] = NOUGHT
				snapshot[1][2] = NOUGHT

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{1, 1}))
			})
		})
		Context("when two noughts are on the first column", func() {
			It("puts a nought on the second column free space", func() {
				snapshot[0][0] = NOUGHT
				snapshot[1][0] = NOUGHT

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{2, 0}))
			})
		})
		Context("when two noughts are on one of the cross line", func() {
			It("puts a nought on the cross line free space", func() {
				snapshot[0][0] = NOUGHT
				snapshot[2][2] = NOUGHT

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{1, 1}))
			})
		})
		Context("when two noughts are on the othe cross line", func() {
			It("puts a nought on the cross line free space", func() {
				snapshot[0][2] = NOUGHT
				snapshot[1][1] = NOUGHT

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{2, 0}))
			})
		})
	})
	Describe("when the player is trying to not loose", func() {
		Context("when two crosses are on the first line", func() {
			It("puts a nought on the first line free space", func() {
				snapshot[0][0] = NOUGHT
				snapshot[1][0] = CROSS
				snapshot[1][1] = CROSS

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{1, 2}))

			})
		})
		Context("when two crosses are on the first column", func() {
			It("puts a nought on the first column space", func() {
				snapshot[0][0] = CROSS
				snapshot[2][0] = CROSS

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{1, 0}))

			})
		})
		Context("when two crosses are on the first cross line", func() {
			It("puts a nought on the first cross line space", func() {
				snapshot[0][0] = CROSS
				snapshot[2][2] = CROSS

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{1, 1}))

			})
		})
		Context("when two crosses are on the second cross line", func() {
			It("puts a nought on the second cross line space", func() {
				snapshot[0][2] = CROSS
				snapshot[2][0] = CROSS

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{1, 1}))

			})
		})
	})
	Describe("when any player has not chance to win", func() {
		Context("when the board is empty", func() {
			It("puts the token on the center of the board", func() {
				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{1, 1}))

			})
		})
		Context("when there is one token on the board", func() {
			Context("when central space is not free", func() {
				It("puts a nought on the first free space of the board", func() {
					snapshot[1][1] = CROSS

					position := player.PutToken(snapshot)

					Expect(position).Should(Equal(Position{0, 0}))

				})
			})
			Context("when central space is free", func() {
				It("puts a nought on central space", func() {
					snapshot[0][0] = CROSS

					position := player.PutToken(snapshot)

					Expect(position).Should(Equal(Position{1, 1}))

				})
			})
		})
	})
})
