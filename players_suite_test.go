package unbeatable_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPlayers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "unbeatable TicTacToe players Suite")
}
