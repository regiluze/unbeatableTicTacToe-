package unbeatable_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTicTacToeBoard(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "unbeatable TicTacToe game Suite")
}
