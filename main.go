package main

import (
	"fmt"
	"math"
	"math/rand"
)

type boardGrid []int

func main() {
	fmt.Println("hello world")

	//printBoard(boardGrid{
	//	11, 12, 21, 22,
	//	32, 33, 42, 43,
	//	51, 52, 61, 62,
	//	71, 72, 81, 82,
	//})
	printBoard(boardGrid{
		1, 1, 2, 2,
		3, 3, 4, 4,
		5, 5, 6, 6,
		7, 7, 8, 8,
	})
	printBoard(createBoard(3))
	printBoard(createBoard(4))
	fmt.Println("_______________________")
	//printBoard(boardGrid{
	//	{11, 12, 13}, {21, 22, 23}, {31, 32, 33},
	//	{41, 42, 43}, {51, 52, 53}, {61, 62, 63},
	//	{71, 72, 73}, {81, 82, 83}, {91, 92, 93},
	//})
}

func createBoard(sizeIndex int) boardGrid {
	var sideLength = sizeIndex * sizeIndex
	var board = boardGrid{}

	for row := 0; row < sideLength; row++ {
		for column := 0; column < sideLength; column++ {
			board = append(board, rand.Intn(9))
		}
	}

	return board
}

func printBoard(board boardGrid) {
	var sideLength = int(math.Sqrt(float64(len(board))))
	var sideIndex = int(math.Sqrt(float64(sideLength)))

	printTopRow(sideLength, sideIndex)

	for row := 0; row < sideLength; row++ {
		fmt.Print(`│`)
		for column := 0; column < sideLength; column++ {
			fmt.Print(board[row*sideLength+column], `│`)
		}
		fmt.Println()
	}

	printBottomRow(sideLength, sideIndex)
}

func printTopRow(sideLength int, sideIndex int) {
	fmt.Print(`╔═`)
	for column := 1; column < sideLength; column++ {
		if column == (sideLength - 1) {
			fmt.Print(`══╗`)
			continue
		}
		if column%sideIndex == 0 {
			fmt.Print(`╦═`)
			continue
		}

		fmt.Print(`╤═`)
	}
	fmt.Println()
}
func printBottomRow(sideLength int, sideIndex int) {
	fmt.Print(`╚═`)
	for column := 1; column < sideLength; column++ {
		if column == (sideLength - 1) {
			fmt.Print(`══╝`)
			continue
		}
		if column%sideIndex == 0 {
			fmt.Print(`╦═`)
			continue
		}

		fmt.Print(`╤═`)
	}
	fmt.Println()
}
