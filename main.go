package main

import "fmt"

func putSymbol(arr [3][3]string, symbol string, x *int, y *int) [3][3]string {
	arr[*x][*y] = symbol
	return arr
}

func askCoordinates(x *int, y *int) {
	var a, b int
	flag := 0
	for flag < 2 {
		flag = 0
		println("Enter the number of the row for your next move (0-2)")
		fmt.Scanln(&a)
		if a >= 0 && a <= 2 {
			flag++
		}
		println("Enter the number of the column for your next move (0-2)")
		fmt.Scanln(&b)
		if b >= 0 && b <= 2 {
			flag++
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

func gameLoop(arr [3][3]string) {
	userSymbol := "X"
	//computerSymbol := 'O'
	var x, y *int
	var a int
	x = &a
	y = &a

	askCoordinates(x, y)
	arr = putSymbol(arr, userSymbol, x, y)
	displayBoard(arr)
}

func main() {
	println("Welcome to TicTacToe Game\n")

	board := [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}
	//boardTest := [3][3]string{{"X", "X", "O"}, {"", "", ""}, {"", "", ""}}

	//res := checkHorizontals(boardTest)
	gameLoop(board)

}
