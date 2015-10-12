package main_test

import (
	_ "fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/regiluze/unbeatableTicTacToe-"
)

const ()

var _ = Describe("Tic Tac Toe unbeatable player specs", func() {
	var (
		player UnbeatablePlayer
	)

	BeforeEach(func() {
		player = NewUnbeateablePlayer(NOUGHT)
	})

	Describe("when the player is trying to win the game", func() {
		Context("when two noughts are on the first line", func() {
			It("puts a nought on the first line free space", func() {
				snapshot := [3][3]string{}
				snapshot[0][0] = NOUGHT
				snapshot[0][1] = NOUGHT

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{0, 2}))

			})
		})
		Context("when two noughts are on the second line", func() {
			It("puts a nought on the second line free space", func() {
				snapshot := [3][3]string{}
				snapshot[1][0] = NOUGHT
				snapshot[1][2] = NOUGHT

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{1, 1}))
			})
		})
		Context("when two noughts are on the first column", func() {
			It("puts a nought on the second column free space", func() {
				snapshot := [3][3]string{}
				snapshot[0][0] = NOUGHT
				snapshot[1][0] = NOUGHT

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{2, 0}))
			})
		})
		Context("when two noughts are on one of the cross line", func() {
			It("puts a nought on the cross line free space", func() {
				snapshot := [3][3]string{}
				snapshot[0][0] = NOUGHT
				snapshot[2][2] = NOUGHT

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{1, 1}))
			})
		})
		Context("when two noughts are on the othe cross line", func() {
			It("puts a nought on the cross line free space", func() {
				snapshot := [3][3]string{}
				snapshot[0][2] = NOUGHT
				snapshot[1][1] = NOUGHT

				position := player.PutToken(snapshot)

				Expect(position).Should(Equal(Position{2, 0}))
			})
		})
	})
})
