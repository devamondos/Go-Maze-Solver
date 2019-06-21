package mazesolvers

import (
	"fmt"
	"os"

	m "github.com/devamondos/maze_solver/mazeutils"
)

type mazeMeta struct {
	row       int
	rowPos    int
	direction string
}

func alwaysLeft(maze *m.Maze) {
	meta := mazeMeta{maze.StartRow, maze.StartRowPos, "down"}
	makeMove(maze, &meta)
}

func makeMove(maze *m.Maze, meta *mazeMeta) {
	nextPixel := checkNextPixel(maze, meta)
	var pixel *m.Pixel
	var err error
	if nextPixel == nil {
		fmt.Println("Error: could not find next pixel")
		os.Exit(1)
	}
	if nextPixel.IsWall {
		changeDirection(maze, meta, true)
	} else if nextPixel.IsNode {
		setMazeLocation(meta, nextPixel)
		changeDirection(maze, meta, false)
	} else {
		switch meta.direction {
		case "up":
			fmt.Println("up")
			pixel, err = maze.GetNextNodeUp(meta.row, meta.rowPos)
			break
		case "right":
			fmt.Println("right")
			pixel, err = maze.GetNextNodeRight(meta.row, meta.rowPos)
			break
		case "down":
			fmt.Println("down")
			pixel, err = maze.GetNextNodeDown(meta.row, meta.rowPos)
			break
		case "left":
			fmt.Println("left")
			pixel, err = maze.GetNextNodeLeft(meta.row, meta.rowPos)
			break
		}

	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else if pixel != nil {
		setMazeLocation(meta, pixel)
		if pixel.IsDeadEnd {
			fmt.Printf("Encountered dead end (%d,%d)\n", meta.row, meta.rowPos)
			changeDirection(maze, meta, false)
			// makeMove(maze, meta)
		} else if pixel.IsNode {
			fmt.Printf("Node array pos (%d,%d)\n", pixel.Row, pixel.RowPos)
			changeDirection(maze, meta, false)
		}
	} else {
		fmt.Println("Pixel and error is nil")
	}
}

func setMazeLocation(meta *mazeMeta, pixel *m.Pixel) {
	meta.row = pixel.Row
	meta.rowPos = pixel.RowPos
}

func checkNextPixel(maze *m.Maze, meta *mazeMeta) *m.Pixel {
	var nextRow, nextRowPos int
	switch meta.direction {
	case "up":
		fmt.Println("up")
		nextRow = meta.row - 1
		nextRowPos = meta.rowPos
		break
	case "right":
		fmt.Println("right")
		nextRow = meta.row
		nextRowPos = meta.rowPos + 1
		break
	case "down":
		fmt.Println("down")
		nextRow = meta.row + 1
		nextRowPos = meta.rowPos
		break
	case "left":
		fmt.Println("left")
		nextRow = meta.row
		nextRowPos = meta.rowPos - 1
		break
	}
	nextPixel, _ := maze.GetPixel(nextRow, nextRowPos)
	return nextPixel
}

func changeDirection(maze *m.Maze, meta *mazeMeta, wallEncountered bool) {
	getNextDirection(meta, wallEncountered)
	makeMove(maze, meta)
}

func getNextDirection(meta *mazeMeta, wallEncountered bool) {
	switch meta.direction {
	case "up":
		if wallEncountered {
			meta.direction = "right"
		} else {
			meta.direction = "left"
		}
		break
	case "right":
		if wallEncountered {
			meta.direction = "down"
		} else {
			meta.direction = "up"
		}
		break
	case "down":
		if wallEncountered {
			meta.direction = "left"
		} else {
			meta.direction = "right"
		}
		break
	case "left":
		if wallEncountered {
			meta.direction = "up"
		} else {
			meta.direction = "down"
		}
		break
	}
}
