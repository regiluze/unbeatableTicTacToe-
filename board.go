package unbeatable

import (
	"errors"
	"fmt"
)

const (
	NOUGHT      = "O"
	CROSS       = "X"
	UNTYPED     = ""
	EMPTY_SPACE = "-"
	SIZE        = 3
)

type Token struct {
	Type string
}

type BoardLine [SIZE]string
type BoardSnapshot [SIZE]BoardLine

func (board *BoardSnapshot) Reset() {
	for row := 0; row < SIZE; row++ {
		for col := 0; col < SIZE; col++ {
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
	if board.numberOfTokens == SIZE*SIZE {
		return true
	}
	return false
}

func (board *TicTacToeBoard) IsAnyInLine() (bool, Token) {
	token := board.checkInline()
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

func (board *TicTacToeBoard) checkInline() Token {
	if token, exists := board.inlineOnRows(); exists {
		return token
	}
	if token, exists := board.inlineOnColumns(); exists {
		return token
	}
	if token, exists := board.inlineOnCrosses(); exists {
		return token
	}
	return Token{Type: UNTYPED}
}

func (board *TicTacToeBoard) inlineOnRows() (Token, bool) {
	for row := 0; row < SIZE; row++ {
		if tokenType := board.isInline(board.Snapshot[row]); tokenType != UNTYPED {
			return Token{Type: tokenType}, true
		}
	}
	return Token{Type: UNTYPED}, false
}

func (board *TicTacToeBoard) inlineOnColumns() (Token, bool) {
	for column := 0; column < SIZE; column++ {
		tokensLine := board.extractColumn(column)
		if tokenType := board.isInline(tokensLine); tokenType != UNTYPED {
			return Token{Type: tokenType}, true
		}
	}
	return Token{Type: UNTYPED}, false
}

func (board *TicTacToeBoard) extractColumn(columnIndex int) BoardLine {
	column := BoardLine{}
	for rowIndex := 0; rowIndex < SIZE; rowIndex++ {
		column[rowIndex] = board.Snapshot[rowIndex][columnIndex]
	}
	return column
}

func (board *TicTacToeBoard) inlineOnCrosses() (Token, bool) {
	if tokenType := board.isInline(board.extractFirstCrossline()); tokenType != UNTYPED {
		return Token{Type: tokenType}, true
	}
	if tokenType := board.isInline(board.extractSecondCrossline()); tokenType != UNTYPED {
		return Token{Type: tokenType}, true
	}
	return Token{Type: UNTYPED}, false
}

func (board *TicTacToeBoard) extractFirstCrossline() BoardLine {
	column := BoardLine{}
	for rowIndex := 0; rowIndex < SIZE; rowIndex++ {
		column[rowIndex] = board.Snapshot[rowIndex][rowIndex]
	}
	return column
}

func (board *TicTacToeBoard) extractSecondCrossline() BoardLine {
	crossLine := BoardLine{}
	for rowIndex := 0; rowIndex < SIZE; rowIndex++ {
		crossLine[rowIndex] = board.Snapshot[rowIndex][SIZE-rowIndex-1]
	}
	return crossLine
}

func (board *TicTacToeBoard) isInline(row BoardLine) string {
	baseTokenType := row[0]
	for _, t := range row {
		if t == EMPTY_SPACE || t != baseTokenType {
			return UNTYPED
		}
	}
	return baseTokenType
}

func (board *TicTacToeBoard) placeIsFilled(position Position) bool {
	return board.Snapshot[position.Col][position.Line] != EMPTY_SPACE
}

func (board BoardSnapshot) Print() {
	fmt.Print("   | ")
	for row := 0; row < SIZE; row++ {
		fmt.Print(row, " | ")
	}
	fmt.Println("")
	fmt.Println("----------------")
	for column := 0; column < SIZE; column++ {
		fmt.Print(" ", column, " ")
		for row := 0; row < SIZE; row++ {
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
