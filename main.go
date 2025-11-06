package main

func printBoard(arr [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(arr[i][j])
		}
		println()
	}
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

func displayBoard() {

}

// 0 | 0 | 0 |
// 0 | 0 | 0 |
// 0 | 0 | 0 |

func main() {
	println("Welcome to TicTacToe Game")

	//board := [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}
	boardTest := [3][3]string{{"X", "X", "O"}, {"", "", ""}, {"", "", ""}}

	res := checkHorizontals(boardTest)
	println(res)

}
