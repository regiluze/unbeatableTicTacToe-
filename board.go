package unbeatable

import (
	"errors"
	"fmt"
)

const (
	NOUGHT          = "O"
	CROSS           = "X"
	UNTYPED         = ""
	EMPTY_SPACE     = "-"
	NOT_INLINE      = ""
	TOKENS_ON_BOARD = 9
)

type Token struct {
	Type string
}

type BoardLine [3]string
type BoardSnapshot [3][3]string

func (board BoardSnapshot) Print() {
	fmt.Println("   | 0 | 1 | 2 |")
	fmt.Println("----------------")
	for column := 0; column < 3; column++ {
		fmt.Print(" ", column, " ")
		for row := 0; row < 3; row++ {
			fmt.Print("|")
			fmt.Print(" ", board[row][column], " ")
			if row == 2 {
				fmt.Print("|")
			}
		}
		fmt.Println("")
		fmt.Println("----------------")
	}
}

func (board *BoardSnapshot) Reset() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			board[row][col] = EMPTY_SPACE
		}
	}
}

type Position struct {
	Col  int
	Line int
}

type TicTacToeBoard struct {
	Snapshot       BoardSnapshot
	numberOfTokens int
}

func NewTicTacToeBoard() *TicTacToeBoard {
	board := &TicTacToeBoard{}
	board.Reset()
	return board
}

func (board *TicTacToeBoard) PutNought(position Position) (BoardSnapshot, error) {
	return board.putToken(NOUGHT, position)
}

func (board *TicTacToeBoard) PutCross(position Position) (BoardSnapshot, error) {
	return board.putToken(CROSS, position)
}

func (board *TicTacToeBoard) Reset() {
	board.numberOfTokens = 0
	board.Snapshot.Reset()
}

func (board *TicTacToeBoard) IsFull() bool {
	if board.numberOfTokens == TOKENS_ON_BOARD {
		return true
	}
	return false
}

func (board *TicTacToeBoard) IsAnyInLine() (bool, Token) {
	token := board.checkThreeInLine()
	match := token.Type != UNTYPED
	return match, token
}

func (board *TicTacToeBoard) putToken(token string, position Position) (BoardSnapshot, error) {
	if board.placeIsFilled(position) {
		return BoardSnapshot{}, errors.New("Place already filled")
	}
	board.Snapshot[position.Col][position.Line] = token
	board.numberOfTokens += 1
	return board.Snapshot, nil
}

func (board *TicTacToeBoard) checkThreeInLine() Token {
	if token, exists := board.inLineOnRows(); exists {
		return token
	}
	if token, exists := board.inLineOnColumns(); exists {
		return token
	}
	if token, exists := board.inLineOnCrosses(); exists {
		return token
	}
	return Token{Type: UNTYPED}
}

func (board *TicTacToeBoard) inLineOnRows() (Token, bool) {
	for row := 0; row < 3; row++ {
		if tokenType := board.threeInLine(board.Snapshot[row]); tokenType != UNTYPED {
			return Token{Type: tokenType}, true
		}
	}
	return Token{Type: UNTYPED}, false
}

func (board *TicTacToeBoard) inLineOnColumns() (Token, bool) {
	for column := 0; column < 3; column++ {
		tokensLine := [3]string{board.Snapshot[0][column], board.Snapshot[1][column], board.Snapshot[2][column]}
		if tokenType := board.threeInLine(tokensLine); tokenType != UNTYPED {
			return Token{Type: tokenType}, true
		}
	}
	return Token{Type: UNTYPED}, false
}

func (board *TicTacToeBoard) inLineOnCrosses() (Token, bool) {
	cross := [3]string{board.Snapshot[0][0], board.Snapshot[1][1], board.Snapshot[2][2]}
	if tokenType := board.threeInLine(cross); tokenType != UNTYPED {
		return Token{Type: tokenType}, true
	}
	cross = [3]string{board.Snapshot[0][2], board.Snapshot[1][1], board.Snapshot[2][0]}
	if tokenType := board.threeInLine(cross); tokenType != UNTYPED {
		return Token{Type: tokenType}, true
	}
	return Token{Type: UNTYPED}, false
}

func (board *TicTacToeBoard) threeInLine(row BoardLine) string {
	token := row[0]
	if token != EMPTY_SPACE && row[1] == token && row[2] == token {
		return token
	}
	return UNTYPED
}

func (board *TicTacToeBoard) placeIsFilled(position Position) bool {
	return board.Snapshot[position.Col][position.Line] != EMPTY_SPACE
}
