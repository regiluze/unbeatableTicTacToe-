package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTicTacToeBoard(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tic tac toe board Suite")
}
