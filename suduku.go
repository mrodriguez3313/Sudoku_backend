package main

import (
	"fmt"
	"strings"
)

type Board struct {
	game_board []uint8
}

var suduku = Board{
	[]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func main() {
	suduku.printBoard()
}

// This function will pretty print the board,
// Written very specifically for a 1D board layout
func (b *Board) printBoard() {
	fmt.Println("| --------------------- |")
	fmt.Print("| ")
	for idx, val := range b.game_board {
		fmt.Print(val, " ")
		if (idx+1)%3 == 0 {
			fmt.Print("| ")
		}
		if (idx+1)%9 == 0 {
			fmt.Println("")
			fmt.Print("| ")
		}
		if (idx+1)%27 == 0 {
			fmt.Println("--------------------- |")
			if !((idx+1)%81 == 0) {
				fmt.Print("| ")
			}
		}
	}
}

// This recursive function will only fill one box
// Parameters: starting index, rolling index, array
func (b *Board) fillBoard(si uint8, ri uint8) {
	// Closing conditions
	if ri >= uint8(len(b.game_board)) {
		return
	}
	if (ri - 33) == si {
		return
	}

}

// This function checks to see if a row contains the same number more than once
// Parameters: array of numbers to check
func (b *Board) checkRow(row []uint8) bool {
	history := [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range row {
		if val == 0 {
			return false
		}
		if history[val-1] == 0 {
			return false
		}
		if val == history[val-1] {
			history[val-1] = 0
		}
	}
	return true
}

// This function checks to see if a col contains the same number more than once
// Parameters: idx of which col to check, array of numbers to check
func checkCol(idx int, col []uint8) bool {
	history := [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range col {
		if history[val-1] == 0 {
			return false
		}
		if val == history[val-1] {
			history[val-1] = 0
		}
	}
	return true
}

// This function checks to see if a grid contains the same number more than once
// Parameters: idx of which grid to check, array of numbers to check
func checkGrid(idx int, grid []uint8) bool {
	history := [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range grid {
		if history[val-1] == 0 {
			return false
		}
		if val == history[val-1] {
			history[val-1] = 0
		}
	}
	return true
}

// This function will create sub-array out of desired selection
// options: G=Grid, R=Row, C=Col. + #1-9 to choose either Grid, Row, or Col
func (b *Board) createSubArray(char string, num uint8) []uint8 {
	char = strings.ToLower(char)
	if char == "r" {
		return (b.game_board[(num-1)*9 : ((num-1)*9)+9])
	}
	return []uint8{0}
}
