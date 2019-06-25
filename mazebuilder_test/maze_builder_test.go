package mazebuilder_test

import (
	"testing"

	mb "github.com/devamondos/maze_solver/mazebuilder"
)

const testfile string = "../test.png"

// array position 0: IsPath
// array position 1: IsNode
var expectedPaths [5][5][2]bool = [5][5][2]bool{
	{{false, false}, {true, true}, {false, false}, {false, false}, {false, false}},
	{{false, false}, {true, true}, {true, false}, {true, true}, {false, false}},
	{{false, false}, {true, false}, {false, false}, {false, false}, {false, false}},
	{{false, false}, {true, true}, {true, false}, {true, true}, {false, false}},
	{{false, false}, {false, false}, {false, false}, {true, true}, {false, false}},
}

func TestRead(t *testing.T) {
	maze := mb.Build(testfile, false)
	for row := 0; row < maze.Rows; row++ {
		for rowIndex := 0; rowIndex < maze.RowLength; rowIndex++ {
			pixel, _ := maze.GetPixel(row, rowIndex)
			if (row == 0 && rowIndex == 1) != pixel.IsStart {
				t.Error("Maze builder did not correctly identify start position")
			}
			if (row == 4 && rowIndex == 3) != pixel.IsEnd {
				t.Error("Maze builder did not correctly identify end position")
			}
			if expectedPaths[row][rowIndex][0] != pixel.IsPath {
				t.Errorf("Pixel at array position (%d,%d) is incorrect", row, rowIndex)
			}
			if expectedPaths[row][rowIndex][1] != pixel.IsNode {
				t.Errorf("Pixel at array position (%d,%d) is incorrect", row, rowIndex)
			}
		}
	}
}
