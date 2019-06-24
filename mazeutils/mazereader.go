package mazeutils

import (
	"fmt"
	"os"

	"github.com/golang/glog"
)

// Analyse maze for path, wall, and node pixels
func Analyse(maze *Maze, print bool) *Maze {
	analysePath(maze)
	// I want to do another loop over the maze rather than do it in the above loop
	// Reason being is that I want all pixels analysed (paths or wall etc.)
	// There are ways it can be done so an improvement could be made at some point
	analyseNodes(maze)

	glog.Infof(" - Number of rows: %d\n", maze.Rows)
	glog.Infof(" - Row length: %d\n", maze.RowLength)
	glog.Infof(" - Start array position: (%d,%d)\n", maze.StartRow, maze.StartRowPos)
	glog.Infof(" - End array position: (%d,%d)\n", maze.EndRow, maze.EndRowPos)
	glog.Infof(" - Number of nodes: %d\n\n", maze.nodes) // Expecting 74

	if print {
		printMaze(maze)
	}
	return maze
}

func analysePath(maze *Maze) {
	rows := len(maze.pixels)
	for row := 0; row < rows; row++ {
		// glog.Printf("Length of row %d: %d\n", (row + 1), len(pixels[row]))
		rowLength := len(maze.pixels[row])
		maze.setRowLength(rowLength)
		for rowPos := 0; rowPos < rowLength; rowPos++ {
			pixel := maze.pixels[row][rowPos]
			// glog.Printf("RGBA of pixel %d,%d: %d,%d,%d,%d\n", row, col, pixel.rgba.R, pixel.rgba.G, pixel.rgba.B)
			if isBlack(pixel) {
				maze.pixels[row][rowPos].setPath(false)
			} else if isTransparent(pixel) {
				maze.pixels[row][rowPos].setPath(true)
				if row == 0 {
					maze.pixels[row][rowPos].IsStart = true
					maze.setStart(row, rowPos)
				} else if (row + 1) == rows {
					maze.pixels[row][rowPos].IsEnd = true
					maze.setEnd(row, rowPos)
				}
			} else {
				glog.Errorf("Error: Program only understands black and transparent pixels. RGBA of offending pixel: %d,%d,%d,%d", pixel.rgba.R, pixel.rgba.G, pixel.rgba.B, pixel.rgba.A)
				os.Exit(1)
			}
		}
	}
	maze.Rows = rows
}

func analyseNodes(maze *Maze) {
	for row := 0; row < maze.Rows; row++ {
		for rowPos := 0; rowPos < maze.RowLength; rowPos++ {
			pixel, err := maze.GetPixel(row, rowPos)
			if err != nil {
				glog.Fatal(err)
			} else if pixel.IsPath {
				if pixel.IsStart || pixel.IsEnd {
					pixel.IsNode = true
					maze.nodes++
				} else {
					upper, right, lower, left := maze.GetSurroundingPaths(row, rowPos)
					pathCount := 0
					if upper != nil {
						pathCount++
					}
					if right != nil {
						pathCount++
					}
					if lower != nil {
						pathCount++
					}
					if left != nil {
						pathCount++
					}
					if pathCount > 0 {
						// glog.Printf("Paths next to pixel (%d,%d): %d\n", x, y, pathCount)
						switch pathCount {
						case 1:
							// Is dead end
							pixel.IsNode = true
							pixel.IsDeadEnd = true
							maze.nodes++
						case 3, 4:
							// Is junction
							pixel.IsNode = true
							maze.nodes++
							break
						case 2:
							if !(upper == nil && lower == nil) && !(left == nil && right == nil) {
								// It's not a right to left path OR upper to lower path
								// Therefor must be a corner
								pixel.IsNode = true
								maze.nodes++
							}
							break
						}
					}
				}
			}
		}
	}
}

func printMaze(maze *Maze) {
	for row := 0; row < maze.Rows; row++ {
		for rowPos := 0; rowPos < maze.RowLength; rowPos++ {
			pixel, err := maze.GetPixel(row, rowPos)
			if err != nil {
				glog.Fatal(err)
			} else {
				if pixel.IsStart || pixel.IsEnd {
					fmt.Print("~")
				} else if pixel.IsNode {
					fmt.Print("O")
				} else if pixel.IsDeadEnd {
					fmt.Print("X")
				} else if pixel.IsPath {
					fmt.Print(" ")
				} else {
					fmt.Print("@")
				}
			}
		}
		fmt.Print("\n")
	}
}

func isBlack(pixel Pixel) bool {
	if pixel.rgba.R == 0 &&
		pixel.rgba.G == 0 &&
		pixel.rgba.B == 0 &&
		pixel.rgba.A == 255 {
		return true
	}
	return false
}

func isTransparent(pixel Pixel) bool {
	if pixel.rgba.R == 0 &&
		pixel.rgba.G == 0 &&
		pixel.rgba.B == 0 &&
		pixel.rgba.A == 0 {
		return true
	}
	return false
}
