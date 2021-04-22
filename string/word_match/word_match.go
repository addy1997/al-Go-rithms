package main

import "fmt"

// Given an m x n grid of characters board and a string word,
// return true if word exists in the grid.
var x = []int{1, -1, 0, 0}
var y = []int{0, 0, 1, -1}

func search(board [][]byte, word string, i int, j int, idx int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '*' {
		return false
	}

	if word[idx] != board[i][j] {
		return false
	}

	if len(word)-1 == idx {
		return true
	}
	board[i][j] = '*'
	for k := 0; k < len(x); k++ {
		if search(board, word, i+x[k], j+y[k], idx+1) {
			return true
		}
	}
	board[i][j] = word[idx]
	return false
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && search(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "SEEN"
	result := exist(board, word)
	fmt.Println(result)
}
