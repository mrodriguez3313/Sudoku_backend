package main

import (
	"reflect"
	"testing"
)

var test_board = Board{
	[]uint8{1, 7, 9, 9, 5, 4, 2, 3, 6,
		3, 6, 1, 2, 7, 9, 8, 5, 4,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 3, 0, 0, 7, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 5, 0, 0, 0, 7, 0, 0,
		9, 0, 1, 2, 3, 4, 5, 0, 0,
		0, 2, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 9},
}

func TestPrint(t *testing.T) {
	t.Run("PrintBoard", func(t *testing.T) {
		test_board.printBoard()
	})
}

func TestCreateSubRow(t *testing.T) {
	var tests = []struct {
		name string
		char string
		val  uint8
		want []uint8
	}{
		{`First_Row`, "r", 1, test_board.game_board[0:9]},
		{`Second_Row`, "r", 2, test_board.game_board[9:18]},
		{`Third_Row`, "r", 3, test_board.game_board[18:27]},
		{`Fourth_Row`, "r", 4, test_board.game_board[27:36]},
		{`Fifth_Row`, "r", 5, test_board.game_board[36:45]},
		{`Six_Row`, "r", 6, test_board.game_board[45:54]},
		{`Seventh_Row`, "r", 7, test_board.game_board[54:63]},
		{`Eighth_Row`, "r", 8, test_board.game_board[63:72]},
		{`Ninth_Row`, "r", 9, test_board.game_board[72:81]},
	}
	for _, tt := range tests {
		t.Run("Create Sub Rows", func(t *testing.T) {
			result := test_board.createSubArray(tt.char, tt.val)
			if len(result) != len(tt.want) {
				t.Errorf("Slices are not equal size.. want:%d res:%d", len(tt.want), len(result))
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Fatal("Slices not the same.. want: ", tt.want, "res: ", result)
			}
		})
	}
}

func TestRowCheck(t *testing.T) {
	var tests = []struct {
		name string
		row  []uint8
		want bool
	}{
		{"Check correct", []uint8{1, 7, 9, 8, 5, 4, 2, 3, 6}, true},
		{"Check duplicate number", []uint8{1, 7, 7, 8, 5, 4, 2, 6, 6}, false},
		{"Check one zero", []uint8{1, 7, 0, 8, 5, 4, 2, 9, 6}, false},
		{"Check multiple zeros", []uint8{1, 7, 0, 8, 5, 4, 2, 0, 6}, false},
		{"Check all zeros", []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
	}
	for _, tt := range tests {
		t.Run("Check Sub Row", func(t *testing.T) {
			result := test_board.checkRow(tt.row)
			if result != tt.want {
				t.Error("Repeat number detected in row", tt.row)
			}
		})
	}
}
