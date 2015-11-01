package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	HUMAN    = 1
	COMPUTER = 2
)

func main() {
	welcome()

	player1Selection := selectPlayer("Player 1")
	player2Selection := selectPlayer("Player 2")
	fmt.Println("your choice is 1", player1Selection)
	fmt.Println("your choice is 2", player2Selection)

}

func welcome() {
	fmt.Println("********* welcome to tic tac toe game ***************")
	for i := 0; i < 4; i++ {
		time.Sleep(time.Millisecond * 300)
		fmt.Print("*")
	}
	fmt.Println("***********************************************")
}

func selectPlayer(player string) int {
	fmt.Println(fmt.Sprintf("Select %s: ", player))
	fmt.Println("1 -> human | 2 -> computer [1, 2]")
	return readFromUser()
}

func readFromUser() int {
	var userSelection string
	_, _ = fmt.Scanf("%s", &userSelection)
	if isWrongInput(userSelection) {
		fmt.Println("Please, select these options [1, 2]")
		fmt.Println("Try again!")
		readFromUser()
	}
	number, _ := strconv.Atoi(userSelection)
	return number
}

func isWrongInput(input string) bool {
	normalizedInput, err := strconv.Atoi(input)
	if err != nil {
		return true
	}
	if normalizedInput < 1 || normalizedInput > 2 {
		return true
	}
	return false
}
