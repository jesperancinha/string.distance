package algorithms

import (
	"bytes"
	"log"
)

type StringDistance interface {
	CalculateDistance(fromString string, toString string) int
}

type StringDistanceImpl struct {
	StringDistance
}

func (dist StringDistanceImpl) traceback(matrixroute [][]int, fromString string, toString string, startIndexFrom int, startIndexTo int) (bytes.Buffer, bytes.Buffer, int) {
	i := startIndexFrom
	j := startIndexTo
	var bufferFrom bytes.Buffer
	var bufferTo bytes.Buffer
	walk := 0

	for i := len(fromString) - 1; i >= startIndexFrom; i-- {
		bufferFrom.WriteRune([]rune(fromString)[i])
	}
	for j := len(toString) - 1; j >= startIndexTo; j-- {
		bufferTo.WriteRune([]rune(toString)[j])
	}

	for !(i == 0 && j == 0) {
		currentarrow := matrixroute[i][j]
		switch currentarrow {
		case int(Diag):
			bufferFrom.WriteRune([]rune(fromString)[i-1])
			bufferTo.WriteRune([]rune(toString)[j-1])
			i--
			j--
		case int(Left):
			bufferFrom.WriteByte('-')
			bufferTo.WriteRune([]rune(toString)[j-1])
			j--
		case int(Up):
			bufferFrom.WriteRune([]rune(fromString)[i-1])
			bufferTo.WriteByte('-')
			i--
		}
		walk++
	}

	return bufferFrom, bufferTo, walk
}

func (dist StringDistanceImpl) revertedStringsAndWalk(bufferFrom bytes.Buffer, bufferTo bytes.Buffer, walk int) (string, string, int) {
	modifiedFrom := revertString(bufferFrom.String())
	modifiedTo := revertString(bufferTo.String())
	log.Println("Number of steps: ", walk)
	log.Println(modifiedFrom)
	log.Println(modifiedTo)
	return modifiedFrom, modifiedTo, walk
}
