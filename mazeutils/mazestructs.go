package mazeutils

import (
	"fmt"
	"os"
)

// Maze struct
type Maze struct {
	pixels      [][]Pixel
	Rows        int
	RowLength   int
	nodes       int
	StartRow    int
	StartRowPos int
	EndRow      int
	EndRowPos   int
}

func (m *Maze) setStart(row int, rowPos int) {
	m.StartRow = row
	m.StartRowPos = rowPos
}

func (m *Maze) setEnd(row int, rowPos int) {
	m.EndRow = row
	m.EndRowPos = rowPos
}

func (m *Maze) setRowLength(rowLength int) {
	if m.RowLength != 0 && (m.RowLength != rowLength) {
		fmt.Println("Error: Image is not a square. Uneven column counts.")
		os.Exit(1)
	}
	m.RowLength = rowLength
}

// GetPixel will return a pixel at a given location
func (m *Maze) GetPixel(row int, rowPos int) (*Pixel, error) {
	if row < len(m.pixels) && rowPos < len(m.pixels[row]) {
		return &m.pixels[row][rowPos], nil
	}
	return nil, fmt.Errorf("Pixel does not exist at array position (%d,%d)", row, rowPos)
}

// GetSurroundingPaths will return the options for a given location in the maze
func (m *Maze) GetSurroundingPaths(row int, rowPos int) (*Pixel, *Pixel, *Pixel, *Pixel) {
	var upperPath, rightPath, lowerPath, leftPath *Pixel
	// Upper
	if (row - 1) > 0 { // ignore upper wall
		pixel, _ := m.GetPixel(row-1, rowPos)
		if pixel.IsPath {
			upperPath = pixel
		}
	}
	// Right
	if (rowPos + 1) < (m.RowLength - 1) { // ignore right wall
		pixel, _ := m.GetPixel(row, rowPos+1)
		if pixel.IsPath {
			rightPath = pixel
		}
	}
	// Lower
	if (row + 1) < (m.Rows - 1) { // ignore bottom wall
		pixel, _ := m.GetPixel(row+1, rowPos)
		if pixel.IsPath {
			lowerPath = pixel
		}
	}
	// Left
	if (rowPos - 1) > 0 { // ignore left wall
		pixel, _ := m.GetPixel(row, rowPos-1)
		if pixel.IsPath {
			leftPath = pixel
		}
	}
	return upperPath, rightPath, lowerPath, leftPath
}

// GetNextNodeUp - gets next node up (stupid linter)
func (m *Maze) GetNextNodeUp(row int, rowPos int) (*Pixel, [][2]int, error) {
	moves := [][2]int{}
	for rowIndex := (row - 1); rowIndex > 0; rowIndex-- { // ignoring top wall
		pixel, err := m.GetPixel(rowIndex, rowPos)
		if err != nil {
			return nil, moves, fmt.Errorf("Pixel does not exist at (%d,%d)", rowIndex, rowPos)
		}
		if !pixel.IsPath {
			return pixel, moves, fmt.Errorf("Error: Next node up is a wall (%d,%d)", rowIndex, rowPos)
		} else if pixel.IsNode {
			return pixel, moves, nil
		} else {
			moves = append(moves, [2]int{rowIndex, rowPos})
			continue
		}
	}
	return nil, moves, fmt.Errorf("Error: Could not find next node (%d,%d)", row, rowPos)
}

// GetNextNodeRight - gets next node right (stupid linter)
func (m *Maze) GetNextNodeRight(row int, rowPos int) (*Pixel, [][2]int, error) {
	moves := [][2]int{}
	for rowPosIndex := (rowPos + 1); rowPosIndex < (m.RowLength - 1); rowPosIndex++ { // ignoring right wall
		pixel, err := m.GetPixel(row, rowPosIndex)
		if err != nil {
			return nil, moves, fmt.Errorf("Pixel does not exist at (%d,%d)", row, rowPosIndex)
		}
		if !pixel.IsPath {
			return pixel, moves, fmt.Errorf("Error: Next node up is a wall (%d,%d)", row, rowPosIndex)
		} else if pixel.IsNode {
			return pixel, moves, nil
		} else {
			moves = append(moves, [2]int{row, rowPosIndex})
			continue
		}
	}
	return nil, moves, fmt.Errorf("Error: Could not find next node (%d,%d)", row, rowPos)
}

// GetNextNodeDown - gets next node down (stupid linter)
func (m *Maze) GetNextNodeDown(row int, rowPos int) (*Pixel, [][2]int, error) {
	moves := [][2]int{}
	for rowIndex := (row + 1); rowIndex < m.Rows; rowIndex++ { // not ignoring bottom wall, that's where the exit is
		pixel, err := m.GetPixel(rowIndex, rowPos)
		if err != nil {
			return nil, moves, fmt.Errorf("Pixel does not exist at (%d,%d)", rowIndex, rowPos)
		}
		if !pixel.IsPath {
			return pixel, moves, fmt.Errorf("Error: Next node up is a wall (%d,%d)", rowIndex, rowPos)
		} else if pixel.IsNode {
			return pixel, moves, nil
		} else {
			moves = append(moves, [2]int{rowIndex, rowPos})
			continue
		}
	}
	return nil, moves, fmt.Errorf("Error: Could not find next node (%d,%d)", row, rowPos)
}

// GetNextNodeLeft - gets next node left (stupid linter)
func (m *Maze) GetNextNodeLeft(row int, rowPos int) (*Pixel, [][2]int, error) {
	moves := [][2]int{}
	for rowPosIndex := (rowPos - 1); rowPosIndex > 0; rowPosIndex-- { // ignoring left wall
		pixel, err := m.GetPixel(row, rowPosIndex)
		if err != nil {
			return nil, moves, fmt.Errorf("Pixel does not exist at (%d,%d)", row, rowPosIndex)
		}
		if !pixel.IsPath {
			return pixel, moves, fmt.Errorf("Error: Next node up is a wall (%d,%d)", row, rowPosIndex)
		} else if pixel.IsNode {
			return pixel, moves, nil
		} else {
			moves = append(moves, [2]int{row, rowPosIndex})
			continue
		}
	}
	return nil, moves, fmt.Errorf("Error: Could not find next node (%d,%d)", row, rowPos)
}

// Rgba struct
type Rgba struct {
	R, G, B, A int
}

// Pixel struct example
type Pixel struct {
	rgba      Rgba
	Row       int
	RowPos    int
	visited   bool
	IsStart   bool
	IsEnd     bool
	IsPath    bool
	IsNode    bool
	IsDeadEnd bool
}

func (p *Pixel) setPath(isPath bool) {
	// Probably don't need both of these values since one would tell us what we need to know
	// It just reads better
	p.IsPath = isPath
}
