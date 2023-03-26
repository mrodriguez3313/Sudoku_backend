package main

import (
	// "crypto/rand"
	"math/rand"
	"math/big"
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
	drand := "76a48fc2dce4bada35518640a0e414d2cf4566b3e7a1b8a5071b70eb5be34444"
	bi := new(big.Int)
	bi.SetString(drand, 16)
	
	bigUint := bi.Uint64()
	
	seeded := rand.NewSource(int64(bigUint))
	r1 := rand.New(seeded)
	randNum := uint8(r1.Intn(9)+1)

	fmt.Println("psuedo-random: ", randNum)
	
	// suduku.fillBoard(0, randNum)
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

// This recursive function will fill the whole grid
// Parameters: starting index, rolling index, array
func (b *Board) fillBoard(ri int, val uint8, tries uint8) {
	// Closing condition should be:
	// if every Col,Row,&Grid are solved
	// Or I reached the end of the board??
	if ri >= len(b.game_board) {
		return
	}
	if tries >= 9 {
		b.fillBoard(ri-1, val+1, 0)
	}
	if val >= 9 {
		val = 1
	}
	// for idx, val := range b.game_board {
	// 	b.game_board[ri] = val
	// // 	// if bad row, backtrack
	// 	for i, v := range b.game_board[9*(ri/9):ri] {
	// 		if ri == v {
	// 			b.fillBoard(ri, val+1)
	// 		}
	// 	}
	// 	// if good, continue
	// 	b.fillBoard(ri+1, val+1)
	// }
	// 1. Place a number then check.
	// 2. Backtrack if the number placed has already been placed in the row before.
	// if number has already been placed in the last i - remainder 9. Then backtrack.
	// 3. Backtrack if number placed has been seen in Col before.
	// if number has been placed in last i - 9 places. Then backtrack.
	// 4. Backtrack if number placed has been seen in 3x3 Grid
	// Use array method. i  will give me their Grid location

}

// This function will do a final check of the whole grid
// Return True if complete and False if there is a mistake
func (b *Board) checkFin() bool {
	char := []string{"R", "C", "G"}
	for i := 0; i < 3; i++ {
		for j := 1; j <= 9; j++ {
			if !b.check(b.createSubArray(char[i], uint8(j))) {
				return false
			}
		}
	}
	return true
}

// This function checks to see if a row contains the same number more than once
// Parameters: array of numbers to check
func (b *Board) check(row []uint8) bool {
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

// This function will create sub-array out of desired selection
// options: G=Grid, R=Row, C=Col. + #1-9 to choose either Grid, Row, or Col
func (b *Board) createSubArray(char string, num uint8) []uint8 {
	char = strings.ToLower(char)
	if num <= 0 || num >= 10 {
		fmt.Printf("num out of bounds")
	}
	switch char {
	case "r":
		return b.game_board[(num-1)*9 : ((num-1)*9)+9]
	case "c":
		tempArr := []uint8{}
		for idx := num - 1; int(idx) < len(b.game_board); idx += 9 {
			tempArr = append(tempArr, b.game_board[idx])
		}
		return tempArr
	case "g":
		tempArr := []uint8{}
		m := map[int]int{1: 0, 2: 3, 3: 6, 4: 27, 5: 30, 6: 33, 7: 54, 8: 57, 9: 60}
		for idx := m[int(num)]; len(tempArr) < 9; idx += 9 {
			tempArr = append(tempArr, b.game_board[idx], b.game_board[idx+1], b.game_board[idx+2])
		}
		return tempArr
	}
	return []uint8{0}
}
