package main

import "fmt"

func displayBoard(arr [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(arr[i][j])
		}
		println()
	}
}

// [0, 0, 0]
// [0, 0, 0]
// [0, 0, 0]

func main() {
	fmt.Println("Hello, world!")

	var first bool
	first = true

	second := false

	fmt.Println(first, second)
	array := [3]int{1, 2, 3}

	for _, value := range array {
		fmt.Print(value)
	}
	println()

	matrix := [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	displayBoard(matrix)

}
