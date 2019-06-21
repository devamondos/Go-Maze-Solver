package mazeutils

import (
	"fmt"
	"os"
)

// Maze struct
type Maze struct {
	pixels      [][]Pixel
	rows        int
	rowLength   int
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
	if m.rowLength != 0 && (m.rowLength != rowLength) {
		fmt.Println("Error: Image is not a square. Uneven column counts.")
		os.Exit(1)
	}
	m.rowLength = rowLength
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
		if pixel.isPath {
			upperPath = pixel
		}
	}
	// Right
	if (rowPos + 1) < (m.rowLength - 1) { // ignore right wall
		pixel, _ := m.GetPixel(row, rowPos+1)
		if pixel.isPath {
			rightPath = pixel
		}
	}
	// Lower
	if (row + 1) < (m.rows - 1) { // ignore bottom wall
		pixel, _ := m.GetPixel(row+1, rowPos)
		if pixel.isPath {
			lowerPath = pixel
		}
	}
	// Left
	if (rowPos - 1) > 0 { // ignore left wall
		pixel, _ := m.GetPixel(row, rowPos-1)
		if pixel.isPath {
			leftPath = pixel
		}
	}
	return upperPath, rightPath, lowerPath, leftPath
}

// GetNextNodeUp - gets next node up (stupid linter)
func (m *Maze) GetNextNodeUp(row int, rowPos int) (*Pixel, error) {
	for rowIndex := (row - 1); rowIndex > 1; rowIndex-- { // ignoring top wall
		pixel, err := m.GetPixel(rowIndex, rowPos)
		if err != nil {
			return nil, fmt.Errorf("Pixel does not exist at (%d,%d)", rowIndex, row)
		}
		if pixel.isWall {
			return pixel, fmt.Errorf("Error: Next node up is a wall (%d,%d)", rowIndex, row)
		} else if pixel.IsNode {
			return pixel, nil
		} else if pixel.isPath {
			pixel, err := m.GetNextNodeUp(rowIndex, row)
			if err != nil {
				return nil, err
			}
			return pixel, nil
		} else {
			return nil, fmt.Errorf("Error: Pixel is unknown (%d,%d)", rowIndex, row)
		}
	}
	return nil, fmt.Errorf("Error: Could not find next node (%d,%d)", row, rowPos)
}

// GetNextNodeRight - gets next node right (stupid linter)
func (m *Maze) GetNextNodeRight(row int, rowPos int) (*Pixel, error) {
	for rowPosIndex := (rowPos + 1); rowPosIndex < (m.rowLength - 2); rowPosIndex++ { // ignoring right wall
		pixel, err := m.GetPixel(row, rowPosIndex)
		if err != nil {
			return nil, fmt.Errorf("Pixel does not exist at (%d,%d)", row, rowPosIndex)
		}
		if pixel.isWall {
			return pixel, fmt.Errorf("Error: Next node up is a wall (%d,%d)", row, rowPosIndex)
		} else if pixel.IsNode {
			return pixel, nil
		} else if pixel.isPath {
			pixel, err := m.GetNextNodeRight(row, rowPosIndex)
			if err != nil {
				return nil, err
			}
			return pixel, nil
		} else {
			return nil, fmt.Errorf("Error: Pixel is unknown (%d,%d)", row, rowPosIndex)
		}
	}
	return nil, fmt.Errorf("Error: Could not find next node (%d,%d)", row, rowPos)
}

// GetNextNodeDown - gets next node down (stupid linter)
func (m *Maze) GetNextNodeDown(row int, rowPos int) (*Pixel, error) {
	for rowIndex := (row + 1); rowIndex < (m.rows - 1); rowIndex++ { // not ignoring bottom wall, that's where the exit is
		pixel, err := m.GetPixel(rowIndex, rowPos)
		if err != nil {
			return nil, fmt.Errorf("Pixel does not exist at (%d,%d)", rowIndex, rowPos)
		}
		if pixel.isWall {
			return pixel, fmt.Errorf("Error: Next node up is a wall (%d,%d)", rowIndex, rowPos)
		} else if pixel.IsNode {
			return pixel, nil
		} else if pixel.isPath {
			pixel, err := m.GetNextNodeDown(rowIndex, rowPos)
			if err != nil {
				return nil, err
			}
			return pixel, nil
		} else {
			return nil, fmt.Errorf("Error: Pixel is unknown (%d,%d)", rowIndex, rowPos)
		}
	}
	return nil, fmt.Errorf("Error: Could not find next node (%d,%d)", row, rowPos)
}

// GetNextNodeLeft - gets next node left (stupid linter)
func (m *Maze) GetNextNodeLeft(row int, rowPos int) (*Pixel, error) {
	for rowPosIndex := (rowPos - 1); rowPosIndex > 1; rowPosIndex-- { // ignoring left wall
		pixel, err := m.GetPixel(row, rowPosIndex)
		if err != nil {
			return nil, fmt.Errorf("Pixel does not exist at (%d,%d)", row, rowPosIndex)
		}
		if pixel.isWall {
			return pixel, fmt.Errorf("Error: Next node up is a wall (%d,%d)", row, rowPosIndex)
		} else if pixel.IsNode {
			return pixel, nil
		} else if pixel.isPath {
			pixel, err := m.GetNextNodeLeft(row, rowPosIndex)
			if err != nil {
				return nil, err
			}
			return pixel, nil
		} else {
			return nil, fmt.Errorf("Error: Pixel is unknown (%d,%d)", row, rowPosIndex)
		}
	}
	return nil, fmt.Errorf("Error: Could not find next node (%d,%d)", row, rowPos)
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
	isPath    bool
	isWall    bool
	isStart   bool
	isEnd     bool
	IsNode    bool
	IsDeadEnd bool
}

func (p *Pixel) setPath(isPath bool) {
	// Probably don't need both of these values since one would tell us what we need to know
	// It just reads better
	p.isPath = isPath
	p.isWall = !isPath
}
