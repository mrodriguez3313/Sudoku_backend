package main

import (
	"os"
	"reflect"
	"testing"
)

var complete_board = Board{[][]Cell{
		[]Cell{ {1,[]int{},1}, {2,[]int{},1}, {3,[]int{},1}, {6,[]int{},2}, {7,[]int{},2}, {8,[]int{},2}, {9,[]int{},3}, {4,[]int{},3}, {5,[]int{},3}},
		[]Cell{ {5,[]int{},1}, {8,[]int{},1}, {4,[]int{},1}, {2,[]int{},2}, {3,[]int{},2}, {9,[]int{},2}, {7,[]int{},3}, {6,[]int{},3}, {1,[]int{},3}},
		[]Cell{ {9,[]int{},1}, {6,[]int{},1}, {7,[]int{},1}, {1,[]int{},2}, {4,[]int{},2}, {5,[]int{},2}, {3,[]int{},3}, {2,[]int{},3}, {8,[]int{},3}},
		[]Cell{ {3,[]int{},4}, {7,[]int{},4}, {2,[]int{},4}, {4,[]int{},5}, {6,[]int{},5}, {1,[]int{},5}, {5,[]int{},6}, {8,[]int{},6}, {9,[]int{},6}},
		[]Cell{ {6,[]int{},4}, {9,[]int{},4}, {1,[]int{},4}, {5,[]int{},5}, {8,[]int{},5}, {3,[]int{},5}, {2,[]int{},6}, {7,[]int{},6}, {4,[]int{},6}},
		[]Cell{ {4,[]int{},4}, {5,[]int{},4}, {8,[]int{},4}, {7,[]int{},5}, {9,[]int{},5}, {2,[]int{},5}, {6,[]int{},6}, {1,[]int{},6}, {3,[]int{},6}},
		[]Cell{ {8,[]int{},7}, {3,[]int{},7}, {6,[]int{},7}, {9,[]int{},8}, {2,[]int{},8}, {4,[]int{},8}, {1,[]int{},9}, {5,[]int{},9}, {7,[]int{},9}},
		[]Cell{ {2,[]int{},7}, {1,[]int{},7}, {9,[]int{},7}, {8,[]int{},8}, {5,[]int{},8}, {7,[]int{},8}, {4,[]int{},9}, {3,[]int{},9}, {6,[]int{},9}},
		[]Cell{ {7,[]int{},7}, {4,[]int{},7}, {5,[]int{},7}, {3,[]int{},8}, {1,[]int{},8}, {6,[]int{},8}, {8,[]int{},9}, {9,[]int{},9}, {2,[]int{},9}},
	},
}

const N = 9
var test_board = Board{make([][]Cell, N)}

func TestMain(m *testing.M){
	gridNum := 1
		// init board to all 0's & a complete allow list, with proper grid number
		for i := 0; i < N; i++ {
			if (i+1) % 3 == 1 {
				gridNum=i+1
			}
			for j := 1; j <= N; j++ {
				test_board.game_board[i] = append(test_board.game_board[i], Cell{0, []int{1,2,3,4,5,6,7,8,9}, gridNum})
				if j % 3 == 0 {
					gridNum++
				}
			}
			gridNum = gridNum - 3
		}
		os.Exit(m.Run())
}

func TestPrint(t *testing.T) {
	t.Run("PrintBoard", func(t *testing.T) {
		test_board.printBoard()
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