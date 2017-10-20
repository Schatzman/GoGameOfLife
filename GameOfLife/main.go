package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	clearScreen()
	fmt.Printf("\n\n\n\tThis is Conway's Game of Life!")
	time.Sleep(3 * time.Second)

	startingBoard := [][]string{
		{"*", ".", ".", ".", "*", "*", ".", ".", "*", ".", "*", ".", "*", ".", "*", ".", "*", ".", "*", "*", ".", ".", ".", "."},
		{".", ".", ".", ".", "*", ".", "*", ".", ".", ".", ".", "*", ".", ".", ".", ".", ".", ".", ".", "*", "*", ".", "*", "*"},
		{"*", "*", ".", ".", "*", "*", ".", ".", "*", ".", ".", ".", "*", "*", ".", ".", "*", ".", "*", ".", "*", "*", ".", "."},
		{"*", ".", "*", ".", "*", "*", ".", ".", "*", ".", ".", ".", ".", "*", ".", ".", "*", ".", ".", ".", ".", "*", ".", "."},
		{"*", "*", "*", ".", ".", ".", "*", ".", ".", ".", ".", ".", "*", ".", "*", ".", "*", ".", ".", ".", "*", "*", ".", "."},
		{".", ".", ".", "*", "*", ".", ".", ".", "*", "*", ".", "*", "*", "*", ".", ".", ".", "*", "*", ".", ".", "*", ".", "."},
		{"*", "*", ".", "*", "*", "*", ".", ".", "*", ".", ".", "*", ".", ".", ".", "*", ".", ".", ".", "*", "*", ".", ".", "*"},
		{"*", ".", "*", ".", "*", ".", "*", ".", "*", ".", "*", ".", "*", "*", ".", ".", ".", ".", ".", "*", "*", ".", ".", "."},
		{"*", ".", ".", ".", "*", "*", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "*", ".", ".", "*", ".", "."},
		{".", ".", ".", ".", "*", ".", "*", ".", ".", ".", "*", ".", ".", "*", ".", ".", ".", ".", ".", ".", "*", "*", ".", "."},
		{"*", "*", ".", ".", "*", "*", ".", ".", "*", "*", "*", ".", ".", ".", "*", ".", "*", ".", "*", ".", "*", "*", ".", "."},
		{"*", ".", "*", ".", "*", "*", ".", ".", "*", ".", ".", ".", "*", "*", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"*", "*", "*", ".", ".", ".", "*", ".", "*", "*", "*", ".", ".", ".", "*", ".", "*", ".", ".", ".", "*", "*", ".", "."},
		{".", ".", ".", "*", "*", ".", ".", ".", ".", "*", "*", ".", ".", "*", "*", ".", ".", ".", ".", ".", "*", "*", ".", "."},
		{"*", "*", ".", "*", "*", "*", ".", ".", ".", "*", "*", ".", ".", "*", ".", ".", ".", "*", ".", "*", "*", ".", "*", "*"},
		{"*", ".", "*", ".", "*", ".", "*", ".", "*", "*", ".", "*", "*", "*", ".", ".", "*", "*", "*", ".", ".", ".", "*", "."},
		{"*", ".", ".", ".", "*", "*", ".", ".", "*", "*", ".", ".", "*", "*", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "*", ".", "*", ".", ".", "*", ".", "*", ".", "*", ".", ".", ".", ".", ".", "*", "*", ".", ".", "."},
		{"*", "*", ".", ".", "*", "*", ".", ".", "*", "*", ".", ".", "*", ".", ".", ".", ".", ".", ".", "*", "*", ".", ".", "*"},
		{"*", ".", "*", ".", "*", "*", ".", ".", "*", ".", "*", ".", "*", ".", "*", ".", ".", ".", "*", ".", "*", ".", ".", "*"},
		{"*", "*", "*", ".", ".", ".", "*", ".", "*", "*", ".", ".", "*", "*", "*", ".", ".", ".", ".", ".", "*", "*", ".", "."},
		{".", ".", ".", "*", "*", ".", ".", ".", "*", "*", ".", ".", "*", "*", ".", ".", "*", ".", ".", ".", "*", "*", ".", "."},
		{"*", "*", ".", "*", "*", "*", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "*", ".", ".", "*", ".", ".", ".", "*"},
		{"*", ".", "*", ".", "*", ".", "*", ".", "*", ".", "*", ".", "*", ".", "*", ".", ".", ".", "*", ".", "*", "*", ".", "."},
	}

	setBoard(startingBoard)

	for i := 0; i < 1000; i++ {
		clearScreen()
		fmt.Printf("\n\n\n\n")
		getNextIteration()
		for _, row := range gameBoard {
			fmt.Printf("\n\t")
			for _, cell := range row {
				fmt.Printf("%v", cell)
			}
		}
		time.Sleep(300 * time.Millisecond)
	}
}

type board [][]string

var gameBoard board

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func setBoard(ss [][]string) {
	gameBoard = board(ss)
}

func isAlive(s string) bool {
	return s == "*"
}

func getAboveRowNeighbors(x, y int) []string {
	var s []string

	if x > 0 {
		s = getRowNeighbors(x-1, y, true)
	} else if x == 0 {
		s = getRowNeighbors(len(gameBoard)-1, y, true)
	}
	return s
}

func getSameRowNeighbors(x, y int) []string {
	s := getRowNeighbors(x, y, false)
	return s
}

func getBelowRowNeighbors(x, y int) []string {
	var s []string

	if x+1 < len(gameBoard) {
		s = getRowNeighbors(x+1, y, true)
	} else if x+1 == len(gameBoard) {
		s = getRowNeighbors(0, y, true)
	}
	return s
}

func getRowNeighbors(x, y int, getOwnCell bool) []string {
	var s []string

	if y > 0 {
		s = append(s, gameBoard[x][y-1])
	} else if y == 0 {
		s = append(s, gameBoard[x][len(gameBoard[x])-1])
	}
	if getOwnCell {
		s = append(s, gameBoard[x][y])
	}
	if y+1 < len(gameBoard[x]) {
		s = append(s, gameBoard[x][y+1])
	} else if y+1 == len(gameBoard[x]) {
		s = append(s, gameBoard[x][0])
	}
	return s
}

func concatSlicesOfString(s1 []string, s2 []string) []string {
	for _, v := range s2 {
		s1 = append(s1, v)
	}
	return s1
}

func getNeighbors(x, y int) []string {
	// adding the boardWrap bool... TODO!!!!
	var neighbors []string

	aboveRowNeighbors := getAboveRowNeighbors(x, y)
	sameRowNeighbors := getSameRowNeighbors(x, y)
	belowRowNeighbors := getBelowRowNeighbors(x, y)

	neighbors = concatSlicesOfString(neighbors, aboveRowNeighbors)
	neighbors = concatSlicesOfString(neighbors, sameRowNeighbors)
	neighbors = concatSlicesOfString(neighbors, belowRowNeighbors)

	return neighbors
}

// We could just count living cells to begin with,
// but this leaves it open for expansion to track
// other characters to expand the game more later.
func countNeighbors(cn []string) int {
	living := 0

	for _, v := range cn {
		if isAlive(v) {
			living = living + 1
		}
	}
	return living
}

// The conwaysRules function enforces the rules of Conway's Game of Life
//
// 1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
// 2. Any live cell with two or three live neighbours lives on to the next generation.
// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
// Clever guy. British chap so he spelled neighbors weirdly.
// https://en.wikipedia.org/wiki/Conway's_Game_of_Life

// Modified slightly. In this version, the board wraps around. So cells on the edge of
// the board react to their neighbors on the opposite edge in the vertical and horizontal
// directions.
func conwaysRules(thisCellAlive bool, livingNeighbors int) string {
	if thisCellAlive {
		if ln := livingNeighbors; ln < 2 || ln > 3 {
			return "."
		}
		return "*"
	}
	if livingNeighbors == 3 {
		return "*"
	}
	return "."
}

func getCellInfo(cellNeighbors []string, x, y int) (int, bool) {
	return countNeighbors(cellNeighbors), isAlive(gameBoard[x][y])
}

func determineCellFate(x, y int) string {
	cellNeighbors := getNeighbors(x, y)
	living, cellAlive := getCellInfo(cellNeighbors, x, y)
	return conwaysRules(cellAlive, living)
}

func checkBoardShape(b board) bool {
	for _, v := range b {
		if len(v) != len(b[0]) {
			fmt.Printf("Imbalanced board: %v, %v", len(v), len(b[0]))
			return false
		}
	}
	return true
}

func getNextIteration() {
	var newBoard board
	var err error

	if checkBoardShape(gameBoard) {
		for i := range gameBoard {
			newRow := []string{}
			for j := range gameBoard[i] {
				newRow = append(newRow, determineCellFate(i, j))
			}
			newBoard = append(newBoard, newRow)
		}
	} else {
		err = fmt.Errorf("board shape is not even - all rows must be the same length")
	}
	if err == nil {
		setBoard(newBoard)
	}
}
