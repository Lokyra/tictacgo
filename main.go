package main

import "fmt"

func putSymbol(arr [3][3]string, symbol string, x int, y int) [3][3]string {
	arr[x][y] = symbol
	return arr
}

func askCoordinates(x *int, y *int) {
	var a, b int
	for true {
		println("Enter the number of the row for your next move (0-2)")
		fmt.Scanln(&a)
		if a < 0 || a > 2 {
			break
		}
		println("Enter the number of the column for your next move (0-2)")
		fmt.Scanln(&b)
		if b < 0 || b > 2 {
			break
		}
	}

	*x = a
	*y = b
}

func checkDiagonals(arr [3][3]string) bool {
	if arr[0][0] == arr[1][1] && arr[1][1] == arr[2][2] && arr[0][0] != "" {
		return true
	}
	if arr[0][2] == arr[1][1] && arr[1][1] == arr[2][0] && arr[0][2] != "" {
		return true
	}
	return false
}

func checkHorizontals(arr [3][3]string) bool {
	if arr[0][0] == arr[0][1] && arr[0][1] == arr[0][2] && arr[0][0] != "" {
		return true
	}
	if arr[1][0] == arr[1][1] && arr[1][1] == arr[1][2] && arr[1][0] != "" {
		return true
	}
	if arr[2][0] == arr[2][1] && arr[2][1] == arr[2][2] && arr[2][0] != "" {
		return true
	}
	return false
}

func checkVerticals(arr [3][3]string) bool {
	if arr[0][0] == arr[1][0] && arr[1][0] == arr[2][0] && arr[0][0] != "" {
		return true
	}
	if arr[0][1] == arr[1][1] && arr[1][1] == arr[2][1] && arr[0][1] != "" {
		return true
	}
	if arr[0][1] == arr[1][2] && arr[1][2] == arr[2][2] && arr[0][1] != "" {
		return true
	}
	return false
}

func displayBoard(arr [3][3]string) {
	println(arr[0][0], "|", arr[0][1], "|", arr[0][2])
	println(arr[1][0], "|", arr[1][1], "|", arr[1][2])
	println(arr[2][0], "|", arr[2][1], "|", arr[2][2])
}

func main() {
	println("Welcome to TicTacToe Game")

	board := [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}
	//boardTest := [3][3]string{{"X", "X", "O"}, {"", "", ""}, {"", "", ""}}

	//res := checkHorizontals(boardTest)
	displayBoard(board)

}
