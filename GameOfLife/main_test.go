package main

import (
	"reflect"
	"testing"
)

var testBoard0 = [][]string{
	{"*", ".", ".", "."},
	{".", ".", ".", "."},
	{"*", "*", ".", "."},
	{"*", ".", "*", "."},
}

var testBoard1 = [][]string{
	{"*", "*", "*"},
	{"*", "*", "*"},
	{"*", "*", "*"},
}

var testBoardImbalanced0 = [][]string{
	{"."},
	{".", "*"},
	{".", "*"},
}

// the following expected result is derived from testBoard0 -> getNextIteration()
var testGetNextIterationExpectedBoard0 = [][]string{
	{".", "*", ".", "*"},
	{"*", "*", ".", "."},
	{"*", "*", ".", "*"},
	{"*", ".", ".", "."},
}

func checkExpectedEqualsActualString(es, as string, t *testing.T) {
	if es != as {
		t.Errorf("Expected: %v != Actual: %v", es, as)
	}
}

func checkExpectedEqualsActualSlice(funcCallString string, es, as []string, t *testing.T) {
	if reflect.TypeOf(es) != reflect.TypeOf(as) {
		t.Errorf("%v expected type %T but type %T", funcCallString, es, as)
	}

	if len(es) != len(as) {
		t.Errorf("%v expected length %v but was actually %v", funcCallString, len(es), len(as))
	}

	for i := range as {
		if as[i] != es[i] {
			t.Errorf("value at index %v in %v expected to equal \"%v\" but was actually \"%v\"", i, funcCallString, es, as)
		}
	}
}

func checkExpectedEqualsActualSliceOfSlice(funcCallString string, ess, ass [][]string, t *testing.T) {
	for i := range ess {
		checkExpectedEqualsActualSlice(funcCallString, ess[i], ass[i], t)
	}
}

func TestIsAlive(t *testing.T) {
	if !isAlive("*") {
		t.Errorf("isAlive(\"*\"), expected true but got false.")
	}

	if isAlive(".") {
		t.Errorf("isAlive(\".\"), expected false but got true.")
	}

	if isAlive("!$*@!(") {
		t.Errorf("isAlive(\"!$*@!(\"), expected false but got true.")
	}
}

func TestSetBoard(t *testing.T) {
	setBoard(testBoard0)
	testBoard := board(testBoard0)

	if reflect.TypeOf(gameBoard) != reflect.TypeOf(testBoard) {
		t.Errorf("gameBoard (type %T) doesn't have same type as testBoard (type %T)", gameBoard, testBoard)
	}

	if len(gameBoard) != len(testBoard) {
		t.Errorf("len(gameBoard) expected %v but got %v", len(gameBoard), len(testBoard))
	}
}

func TestGetAboveNeighbors(t *testing.T) {
	setBoard(testBoard0)

	getAboveRowNeighborsExpectedX0Y0 := []string{".", "*", "."}
	getAboveRowNeighborsActualX0Y0 := getAboveRowNeighbors(0, 0)

	getAboveRowNeighborsExpectedX1Y1 := []string{"*", ".", "."}
	getAboveRowNeighborsActualX1Y1 := getAboveRowNeighbors(1, 1)

	getAboveRowNeighborsExpectedX3Y3 := []string{".", ".", "*"}
	getAboveRowNeighborsActualX3Y3 := getAboveRowNeighbors(3, 3)

	checkExpectedEqualsActualSlice("getAboveRowNeighbors(0, 0)",
		getAboveRowNeighborsActualX0Y0, getAboveRowNeighborsExpectedX0Y0, t)

	checkExpectedEqualsActualSlice("getAboveRowNeighbors(1, 1)",
		getAboveRowNeighborsActualX1Y1, getAboveRowNeighborsExpectedX1Y1, t)

	checkExpectedEqualsActualSlice("getAboveRowNeighbors(3, 3)",
		getAboveRowNeighborsActualX3Y3, getAboveRowNeighborsExpectedX3Y3, t)
}

func TestGetSameRowNeighbors(t *testing.T) {
	setBoard(testBoard0)

	getSameRowNeighborsExpectedX0Y0 := []string{".", "."}
	getSameRowNeighborsActualX0Y0 := getSameRowNeighbors(0, 0)

	getSameRowNeighborsExpectedX1Y1 := []string{".", "."}
	getSameRowNeighborsActualX1Y1 := getSameRowNeighbors(1, 1)

	getSameRowNeighborsExpectedX3Y3 := []string{"*", "*"}
	getSameRowNeighborsActualX3Y3 := getSameRowNeighbors(3, 3)

	checkExpectedEqualsActualSlice("getSameRowNeighbors(0, 0)",
		getSameRowNeighborsActualX0Y0, getSameRowNeighborsExpectedX0Y0, t)

	checkExpectedEqualsActualSlice("getSameRowNeighbors(1, 1)",
		getSameRowNeighborsActualX1Y1, getSameRowNeighborsExpectedX1Y1, t)

	checkExpectedEqualsActualSlice("getSameRowNeighbors(3, 3)",
		getSameRowNeighborsActualX3Y3, getSameRowNeighborsExpectedX3Y3, t)
}

func TestGetBelowRowNeighbors(t *testing.T) {
	setBoard(testBoard0)

	getBelowRowNeighborsExpectedX0Y0 := []string{".", ".", "."}
	getBelowRowNeighborsActualX0Y0 := getBelowRowNeighbors(0, 0)

	getBelowRowNeighborsExpectedX1Y1 := []string{"*", "*", "."}
	getBelowRowNeighborsActualX1Y1 := getBelowRowNeighbors(1, 1)

	getBelowRowNeighborsExpectedX3Y3 := []string{".", ".", "*"}
	getBelowRowNeighborsActualX3Y3 := getBelowRowNeighbors(3, 3)

	checkExpectedEqualsActualSlice("getBelowRowNeighbors(0, 0)",
		getBelowRowNeighborsActualX0Y0, getBelowRowNeighborsExpectedX0Y0, t)

	checkExpectedEqualsActualSlice("getBelowRowNeighbors(1, 1)",
		getBelowRowNeighborsActualX1Y1, getBelowRowNeighborsExpectedX1Y1, t)

	checkExpectedEqualsActualSlice("getBelowRowNeighbors(3, 3)",
		getBelowRowNeighborsActualX3Y3, getBelowRowNeighborsExpectedX3Y3, t)
}

func TestGetNeighbors(t *testing.T) {
	setBoard(testBoard0)

	getNeighborsExpectedX0Y0 := []string{".", "*", ".", ".", ".", ".", ".", "."}
	getNeighborsActualX0Y0 := getNeighbors(0, 0)

	getNeighborsExpectedX1Y1 := []string{"*", ".", ".", ".", ".", "*", "*", "."}
	getNeighborsActualX1Y1 := getNeighbors(1, 1)

	getNeighborsExpectedX2Y0 := []string{".", ".", ".", ".", "*", ".", "*", "."}
	getNeighborsActualX2Y0 := getNeighbors(2, 0)

	getNeighborsExpectedX3Y3 := []string{".", ".", "*", "*", "*", ".", ".", "*"}
	getNeighborsActualX3Y3 := getNeighbors(3, 3)

	checkExpectedEqualsActualSlice("getNeighbors(0, 0)",
		getNeighborsExpectedX0Y0, getNeighborsActualX0Y0, t)

	checkExpectedEqualsActualSlice("getNeighbors(1, 1)",
		getNeighborsExpectedX1Y1, getNeighborsActualX1Y1, t)

	checkExpectedEqualsActualSlice("getNeighbors(2, 0)",
		getNeighborsExpectedX2Y0, getNeighborsActualX2Y0, t)

	checkExpectedEqualsActualSlice("getNeighbors(3, 3)",
		getNeighborsExpectedX3Y3, getNeighborsActualX3Y3, t)
}

func TestConcatSlicesOfString(t *testing.T) {
	s1 := []string{"*", ".", "*"}
	s2 := []string{".", "*", "."}
	expectedSliceOfString := []string{"*", ".", "*", ".", "*", "."}
	actualSliceOfString := concatSlicesOfString(s1, s2)

	checkExpectedEqualsActualSlice("concatSlicesOfString(s1, s2)",
		expectedSliceOfString, actualSliceOfString, t)
}

func TestDetermineCellFate(t *testing.T) {
	setBoard(testBoard0)

	expectedFateX0Y0 := "."
	actualFateX0Y0 := determineCellFate(0, 0)

	expectedFateX1Y0 := "*"
	actualFateX1Y0 := determineCellFate(1, 0)

	expectedFateX1Y1 := "*"
	actualFateX1Y1 := determineCellFate(1, 1)

	expectedFateX2Y1 := "*"
	actualFateX2Y1 := determineCellFate(2, 1)

	expectedFateX2Y2 := "."
	actualFateX2Y2 := determineCellFate(2, 2)

	expectedFateX3Y3 := "."
	actualFateX3Y3 := determineCellFate(3, 3)

	checkExpectedEqualsActualString(expectedFateX0Y0, actualFateX0Y0, t)

	checkExpectedEqualsActualString(expectedFateX1Y0, actualFateX1Y0, t)

	checkExpectedEqualsActualString(expectedFateX1Y1, actualFateX1Y1, t)

	checkExpectedEqualsActualString(expectedFateX2Y1, actualFateX2Y1, t)

	checkExpectedEqualsActualString(expectedFateX2Y2, actualFateX2Y2, t)

	checkExpectedEqualsActualString(expectedFateX3Y3, actualFateX3Y3, t)

	setBoard(testBoard1)

	expectedFateX1Y1t2 := "."
	actualFateX1Y1t2 := determineCellFate(1, 1)

	checkExpectedEqualsActualString(expectedFateX1Y1t2, actualFateX1Y1t2, t)

	setBoard(testBoard0)

	expectedFateWrapBoardX1Y1 := "*"
	actualFateWrapBoardX1Y1 := determineCellFate(0, 1)

	checkExpectedEqualsActualString(expectedFateWrapBoardX1Y1, actualFateWrapBoardX1Y1, t)
}

func TestCheckBoardShape(t *testing.T) {

	if !checkBoardShape(gameBoard) {
		t.Errorf("testBoard0 is not balanced?! That makes no sense")
	}

	setBoard(testBoardImbalanced0)

	if checkBoardShape(gameBoard) {
		t.Errorf("testBoardImbalanced0 is balanced but is expected to not be")
	}
}

func TestGetNextIteration(t *testing.T) {
	setBoard(testBoard0)
	getNextIteration()
	checkExpectedEqualsActualSliceOfSlice("getNextIteration() - testBoard0",
		testGetNextIterationExpectedBoard0, gameBoard, t)

}
