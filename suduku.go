package main

import (
	"fmt"
	"math/big"
	"math/rand"
	// "strings"

	"golang.org/x/exp/slices"
)
type Cell struct {
	val int
	allowed []int
	grid int
}

type Board struct {
	game_board [][]Cell
}

func main() {
	const N = 9
	var Sudoku = Board{make([][]Cell, N)}
	gridNum := 1

	// init board to all 0's & a complete allow list, with proper grid number
	for i := 0; i < N; i++ {
		if (i+1) % 3 == 1 {
			gridNum=i+1
		}
		for j := 1; j <= N; j++ {
			Sudoku.game_board[i] = append(Sudoku.game_board[i], Cell{0, []int{1,2,3,4,5,6,7,8,9}, gridNum})
			if j % 3 == 0 {
				gridNum++
			}
		}
		gridNum = gridNum - 3
	}

	drand := "76a48fc2dce4bada35518640a0e414d2cf4566b3e7a1b8a5071b70eb5be34444"
	bi := new(big.Int)
	bi.SetString(drand, 16)
	
	bigUint := bi.Uint64()
	
	seeded := rand.NewSource(int64(bigUint))
	r1 := rand.New(seeded)
	randNum := r1.Intn(9)+1

	fmt.Println("psuedo-random: ", randNum)
	
	// Pass in the r1 seed.
	// Sudoku.fillBoard(0, randNum, []int{1,2,3,4,5,6,7,8,9})
	Sudoku.printBoard()
}

// This function will pretty print the board,
// Written very specifically for a 1D board layout
func (b *Board) printBoard() {
	fmt.Println("| --------------------- |")
	for idx, row := range b.game_board {
		fmt.Print("| ")
		for jdx, cell := range row {
			fmt.Print(cell.val, " ")
			
			if (jdx+1) % 3 == 0 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if (idx+1) % 3 == 0 {
			fmt.Println("| --------------------- |")
		}
	}
}

// This recursive function will fill the whole grid
// Parameters: starting index, rolling index, array

// pass in an array of allowed values respective to every square
// && looping to check for every row & col. For this to work I need a way to reset values after every row, column & grid.

// Pass in the seed value so I can keep getting random values from it
// For this to work, i need remove the currently tried value from the allow list
func (b *Board) fillBoard(ri int, val int, allowed []int) {

// 	// 1. Place a number then check.
// 	// 2. If number has already been placed in the last 9*floor(ri/9). Then try a different number.
// 	// 3. Backtrack if number placed has been seen in Col before.
// 	// if number has been placed in last i - 9 places; for i > 0. Then try different number.
// 	// 4. Backtrack if number placed has been seen in 3x3 Grid
// 	// Use array method. i will give me their Grid location
}

// This function removes an element at specific index from array
func remove(arr []int, key int) []int {
	if i, exists := find(arr, key); exists {
		return slices.Delete(arr, i, i+1)
	}
	return arr
}

// This function finds and returns the index of an element else returns -1 & false
func find(arr[]int, key int) (int, bool) {
	for i := range arr {
		if arr[i] == key {
			return i, true
		}
	}
	return -1, false
}