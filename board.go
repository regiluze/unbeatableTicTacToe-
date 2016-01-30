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

func (board *BoardSnapshot) ExtractColumn(columnIndex int) BoardLine {
	column := BoardLine{}
	for rowIndex := 0; rowIndex < SIZE; rowIndex++ {
		column[rowIndex] = board[rowIndex][columnIndex]
	}
	return column
}

func (board *BoardSnapshot) ExtractFirstCrossline() BoardLine {
	column := BoardLine{}
	for rowIndex := 0; rowIndex < SIZE; rowIndex++ {
		column[rowIndex] = board[rowIndex][rowIndex]
	}
	return column
}

func (board *BoardSnapshot) ExtractSecondCrossline() BoardLine {
	crossLine := BoardLine{}
	for rowIndex := 0; rowIndex < SIZE; rowIndex++ {
		crossLine[rowIndex] = board[rowIndex][SIZE-rowIndex-1]
	}
	return crossLine
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

type Position struct {
	Col  int
	Line int
}

type TokensBoard struct {
	Snapshot       BoardSnapshot
	numberOfTokens int
}

func NewTokensBoard() *TokensBoard {
	board := &TokensBoard{}
	board.Reset()
	return board
}

func (t *TokensBoard) PutNought(position Position) (BoardSnapshot, error) {
	return t.putToken(NOUGHT, position)
}

func (t *TokensBoard) PutCross(position Position) (BoardSnapshot, error) {
	return t.putToken(CROSS, position)
}

func (t *TokensBoard) Reset() {
	t.numberOfTokens = 0
	t.Snapshot.Reset()
}

func (t *TokensBoard) IsFull() bool {
	if t.numberOfTokens == SIZE*SIZE {
		return true
	}
	return false
}

func (t *TokensBoard) IsAnyInLine() (bool, Token) {
	token := t.checkInline()
	match := token.Type != UNTYPED
	return match, token
}

func (t *TokensBoard) putToken(token string, position Position) (BoardSnapshot, error) {
	if t.placeIsFilled(position) {
		return BoardSnapshot{}, errors.New("Place already filled")
	}
	t.Snapshot[position.Col][position.Line] = token
	t.numberOfTokens += 1
	return t.Snapshot, nil
}

func (t *TokensBoard) checkInline() Token {
	if token, exists := t.inlineOnRows(); exists {
		return token
	}
	if token, exists := t.inlineOnColumns(); exists {
		return token
	}
	if token, exists := t.inlineOnCrosses(); exists {
		return token
	}
	return Token{Type: UNTYPED}
}

func (t *TokensBoard) inlineOnRows() (Token, bool) {
	for row := 0; row < SIZE; row++ {
		if tokenType := t.isInline(t.Snapshot[row]); tokenType != UNTYPED {
			return Token{Type: tokenType}, true
		}
	}
	return Token{Type: UNTYPED}, false
}

func (t *TokensBoard) inlineOnColumns() (Token, bool) {
	for column := 0; column < SIZE; column++ {
		tokensLine := t.Snapshot.ExtractColumn(column)
		if tokenType := t.isInline(tokensLine); tokenType != UNTYPED {
			return Token{Type: tokenType}, true
		}
	}
	return Token{Type: UNTYPED}, false
}

func (t *TokensBoard) inlineOnCrosses() (Token, bool) {
	if tokenType := t.isInline(t.Snapshot.ExtractFirstCrossline()); tokenType != UNTYPED {
		return Token{Type: tokenType}, true
	}
	if tokenType := t.isInline(t.Snapshot.ExtractSecondCrossline()); tokenType != UNTYPED {
		return Token{Type: tokenType}, true
	}
	return Token{Type: UNTYPED}, false
}

func (t *TokensBoard) isInline(row BoardLine) string {
	baseTokenType := row[0]
	for _, t := range row {
		if t == EMPTY_SPACE || t != baseTokenType {
			return UNTYPED
		}
	}
	return baseTokenType
}

func (t *TokensBoard) placeIsFilled(position Position) bool {
	return t.Snapshot[position.Col][position.Line] != EMPTY_SPACE
}
