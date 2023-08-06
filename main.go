package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var boardMap = map[string][2]int{
	"0": {0, 0},
	"1": {0, 1},
	"2": {0, 2},

	"3": {1, 0},
	"4": {1, 1},
	"5": {1, 2},

	"6": {2, 0},
	"7": {2, 1},
	"8": {2, 2},
}
var refBoard = Board{
	{"0", "1", "2"},
	{"3", "4", "5"},
	{"6", "7", "8"},
}

var gameBoard = Board{
	{"-", "-", "-"},
	{"-", "-", "-"},
	{"-", "-", "-"},
}

var currentPlayer string = "X"

type Board [][]string

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		isBoardFilled := isBoardFilled()
		isWin := isWin()
		isGameOver := isBoardFilled || isWin
		if isGameOver {
			fmt.Println("x-x-x-x-x-x-x-x-x-x-x")
			fmt.Println(" ")
			refBoard.print()
			gameBoard.print()

			fmt.Println("Game Finished!")
			fmt.Println(currentPlayer, "Is the winner!")
			break
		}
		switchCurrentPlayer()

		fmt.Println("x-x-x-x-x-x-x-x-x-x-x")
		fmt.Println(" ")
		fmt.Println("->", currentPlayer, "'s TURN")
		refBoard.print()
		gameBoard.print()

		cellIndex, err := reader.ReadString('\n')

		if err != nil {
			handleErr(err)
		}

		writeOnCellCords(cellIndex)
	}
}

func writeOnCellCords(value string) {
	value = strings.TrimSpace(value)
	cords, ok := boardMap[value]
	if ok != true {
		handleErr("Number is not on range (0 - 8)")
	}
	row := cords[0]
	index := cords[1]
	if gameBoard[row][index] == "-" {
		gameBoard[row][index] = currentPlayer

	} else {
		fmt.Println("You cant play there!")
	}

}

func handleErr(err interface{}) {
	log.Fatal(err)
}

func switchCurrentPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

func (board Board) print() {
	for row := 0; row < 3; row++ {
		for cell := 0; cell < 3; cell++ {
			fmt.Print(" ", board[row][cell])
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")

}

func isBoardFilled() bool {
	for row := 0; row < 3; row++ {
		for cell := 0; cell < 3; cell++ {
			if gameBoard[row][cell] == "-" {
				return false
			}
		}
	}
	return true
}

func isWin() bool {

	// rows
	if gameBoard[0][0] == gameBoard[0][1] && gameBoard[0][1] == gameBoard[0][2] && gameBoard[0][0] != "-" {
		return true
	} else if gameBoard[1][0] == gameBoard[1][1] && gameBoard[1][1] == gameBoard[1][2] && gameBoard[1][0] != "-" {
		return true
	} else if gameBoard[2][0] == gameBoard[2][1] && gameBoard[2][1] == gameBoard[2][2] && gameBoard[2][0] != "-" {
		return true
	} else if gameBoard[0][0] == gameBoard[1][0] && gameBoard[1][0] == gameBoard[2][0] && gameBoard[0][0] != "-" {
		//columns
		return true
	} else if gameBoard[0][1] == gameBoard[1][1] && gameBoard[1][1] == gameBoard[2][1] && gameBoard[0][1] != "-" {
		return true
	} else if gameBoard[0][2] == gameBoard[1][2] && gameBoard[1][2] == gameBoard[2][2] && gameBoard[0][2] != "-" {
		return true
	} else if gameBoard[0][0] == gameBoard[1][1] && gameBoard[1][1] == gameBoard[2][2] && gameBoard[0][0] != "-" {
		//obliques
		return true
	} else if gameBoard[0][2] == gameBoard[1][1] && gameBoard[1][1] == gameBoard[2][0] && gameBoard[0][2] != "-" {
		//obliques
		return true
	}

	return false
}
