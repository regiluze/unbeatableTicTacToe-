package main

import (
	"fmt"
	"strconv"
)

func readFromUser(options []int) int {
	var userSelection string
	_, _ = fmt.Scanf("%s", &userSelection)
	if isCorrectInput(userSelection, options) {
		number, _ := strconv.Atoi(userSelection)
		return number
	}
	fmt.Println("Please, select these options ", options)
	fmt.Println("Try again!")
	return readFromUser(options)
}

func isCorrectInput(input string, options []int) bool {
	normalizedInput, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return inputInOptions(normalizedInput, options)
}

func inputInOptions(input int, options []int) bool {
	for _, option := range options {
		if input == option {
			return true
		}
	}
	return false
}
