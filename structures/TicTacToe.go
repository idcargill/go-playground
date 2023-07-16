package structures

import (
	"fmt"
	"strings"
)


func TicTacToe() {
	board := make([][]string, 3)
	
	board[0] = []string{
		"_",
		"_",
		"X",
	}
	board[1] = []string{
		"_",
		"O",
		"_",
	}
	board[2] = []string{
		"_",
		"_",
		"_",
	}
	
	
	fmt.Println("Tic Tac Toe")
	for i := 0; i <len(board); i ++ {
		fmt.Println(strings.Join(board[i], " "))
	}
}

