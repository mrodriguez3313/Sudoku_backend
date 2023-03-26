package main

import (
	"reflect"
	"testing"
)

var test_board = Board{
	[]int{1, 7, 9, 9, 5, 4, 2, 3, 6,
		3, 6, 1, 2, 7, 9, 8, 5, 4,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 3, 0, 0, 7, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 5, 0, 0, 0, 7, 0, 0,
		9, 0, 1, 2, 3, 4, 5, 0, 0,
		0, 2, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 9},
}

var complete_board = Board{
	[]int{
		1, 2, 3, 6, 7, 8, 9, 4, 5,
		5, 8, 4, 2, 3, 9, 7, 6, 1,
		9, 6, 7, 1, 4, 5, 3, 2, 8,
		3, 7, 2, 4, 6, 1, 5, 8, 9,
		6, 9, 1, 5, 8, 3, 2, 7, 4,
		4, 5, 8, 7, 9, 2, 6, 1, 3,
		8, 3, 6, 9, 2, 4, 1, 5, 7,
		2, 1, 9, 8, 5, 7, 4, 3, 6,
		7, 4, 5, 3, 1, 6, 8, 9, 2},
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
		val  int
		want []int
	}{
		{`First_Row`, "r", 1, test_board.game_board[0:9]},
		{`Second_Row`, "r", 2, test_board.game_board[9:18]},
		{`Third_Row`, "r", 3, test_board.game_board[18:27]},
		{`Fourth_Row`, "r", 4, test_board.game_board[27:36]},
		{`Fifth_Row`, "r", 5, test_board.game_board[36:45]},
		{`Sixth_Row`, "r", 6, test_board.game_board[45:54]},
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

func TestCreateSubCol(t *testing.T) {
	var tests = []struct {
		name string
		char string
		val  int
		want []int
	}{
		{`First_Col`, "c", 1, []int{1, 3, 0, 0, 0, 0, 9, 0, 0}},
		{`Second_Col`, "c", 2, []int{7, 6, 0, 3, 0, 0, 0, 2, 0}},
		{`Third_Col`, "c", 3, []int{9, 1, 0, 0, 0, 5, 1, 0, 0}},
		{`Fourth_Col`, "c", 4, []int{9, 2, 0, 0, 0, 0, 2, 0, 0}},
		{`Fifth_Col`, "c", 5, []int{5, 7, 0, 7, 0, 0, 3, 0, 0}},
		{`Sixth_Col`, "c", 6, []int{4, 9, 0, 0, 0, 0, 4, 0, 0}},
		{`Seventh_Col`, "c", 7, []int{2, 8, 0, 0, 0, 7, 5, 0, 0}},
		{`Eighth_Col`, "c", 8, []int{3, 5, 0, 0, 0, 0, 0, 0, 0}},
		{`Ninth_Col`, "c", 9, []int{6, 4, 0, 0, 0, 0, 0, 0, 9}},
	}
	for _, tt := range tests {
		t.Run("Create Sub Cols", func(t *testing.T) {
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

func TestCreateSubGrid(t *testing.T) {
	var tests = []struct {
		name string
		char string
		val  int
		want []int
	}{
		{`First_Grid`, "g", 1, []int{1, 7, 9, 3, 6, 1, 0, 0, 0}},
		{`Second_Grid`, "g", 2, []int{9, 5, 4, 2, 7, 9, 0, 0, 0}},
		{`Third_Grid`, "g", 3, []int{2, 3, 6, 8, 5, 4, 0, 0, 0}},
		{`Fourth_Grid`, "g", 4, []int{0, 3, 0, 0, 0, 0, 0, 0, 5}},
		{`Fifth_Grid`, "g", 5, []int{0, 7, 0, 0, 0, 0, 0, 0, 0}},
		{`Sixth_Grid`, "g", 6, []int{0, 0, 0, 0, 0, 0, 7, 0, 0}},
		{`Seventh_Grid`, "g", 7, []int{9, 0, 1, 0, 2, 0, 0, 0, 0}},
		{`Eighth_Grid`, "g", 8, []int{2, 3, 4, 0, 0, 0, 0, 0, 0}},
		{`Ninth_Grid`, "g", 9, []int{5, 0, 0, 0, 0, 0, 0, 0, 9}},
	}
	for _, tt := range tests {
		t.Run("Create Sub Grids", func(t *testing.T) {
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
		row  []int
		want bool
	}{
		{"Check correct", []int{1, 7, 9, 8, 5, 4, 2, 3, 6}, true},
		{"Check duplicate number", []int{1, 7, 7, 8, 5, 4, 2, 6, 6}, false},
		{"Check one zero", []int{1, 7, 0, 8, 5, 4, 2, 9, 6}, false},
		{"Check multiple zeros", []int{1, 7, 0, 8, 5, 4, 2, 0, 6}, false},
		{"Check all zeros", []int{0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
	}
	for _, tt := range tests {
		t.Run("Check Sub Row", func(t *testing.T) {
			_, result := test_board.check([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, tt.row)
			if result != tt.want {
				t.Error("Repeat number detected in row", tt.row)
			}
		})
	}
}

// not tested thoroughly
func TestColCheck(t *testing.T) {
	var tests = []struct {
		name string
		col  []int
		want bool
	}{
		{"Check correct", []int{1, 7, 9, 8, 5, 4, 2, 3, 6}, true},
		// {"Check duplicate number", []int{1, 7, 7, 8, 5, 4, 2, 6, 6}, false},
		// {"Check one zero", []int{1, 7, 0, 8, 5, 4, 2, 9, 6}, false},
		// {"Check multiple zeros", []int{1, 7, 0, 8, 5, 4, 2, 0, 6}, false},
		// {"Check all zeros", []int{0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
	}
	for _, tt := range tests {
		t.Run("Check Sub Col", func(t *testing.T) {
			_, result := test_board.check([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, tt.col)
			if result != tt.want {
				t.Error("Repeat number detected in Col", tt.col)
			}
		})
	}
}

func TestCheckFin(t *testing.T) {
	t.Run("Check Finale", func(t *testing.T) {
		if !complete_board.checkFin() {
			t.Error("Board is incomplete")
		}
	})
}

func TestRemoveFun(t *testing.T){
	var tests = []struct{
		Name string
		val int
		test_arr []int
		want []int
	}{
		{"RF-Middle #3", 3, []int{6,2,9,3,7,8}, []int{6,2,9,7,8} },
		{"RF-End #9", 9, []int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8}},
		{" #1, DNE", 1, []int{2,3,4,5,6}, []int{2,3,4,5,6}},
		{"RF-Beg", 4, []int{4,8,9}, []int{8,9}},
		{"RF-Last", 7, []int{7},[]int{}},
		{"RF-Empty", 2, []int{}, []int{}},
	}
	for _, test := range tests {
		t.Run("Test remove func", func(t *testing.T) {
			if !reflect.DeepEqual(test.want, remove(test.test_arr, test.val)) {
				t.Error("Slices not the same. res: ", test.test_arr, " want: ", test.want)
			}
				
		})
	}
}