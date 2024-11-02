package main

import (
	"fmt"
	"os"
	"strings"
)

const size = 9

func main() {
	if len(os.Args) != size+1 {
		fmt.Println("Error")
		return
	}

	board := make([][]byte, size)
	for i := 0; i < size; i++ {
		if len(os.Args[i+1]) != size {
			fmt.Printf("Invalid row length for row %d\n", i+1)
			return
		}
		board[i] = []byte(os.Args[i+1])
	}

	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}

func solveSudoku(board [][]byte) bool {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if board[row][col] == '.' {
				for num := byte('1'); num <= '9'; num++ {
					if isValid(board, row, col, num) {
						board[row][col] = num
						if solveSudoku(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, num byte) bool {
	boxRow, boxCol := row/3*3, col/3*3
	for i := 0; i < size; i++ {
		if board[row][i] == num || board[i][col] == num || board[boxRow+i/3][boxCol+i%3] == num {
			return false
		}
	}
	return true
}

func printBoard(board [][]byte) {
	for i := 0; i < size; i++ {
		fmt.Println(strings.Join(strings.Split(string(board[i]), ""), " "))
	}
}
