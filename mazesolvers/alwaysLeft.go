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
	makeMove(maze, meta)
}

func makeMove(maze *m.Maze, meta mazeMeta) {
	var pixel *m.Pixel
	var err error
	switch meta.direction {
	case "up":
		pixel, err = maze.GetNextNodeUp(meta.row, meta.rowPos)
		break
	case "right":
		pixel, err = maze.GetNextNodeRight(meta.row, meta.rowPos)
		break
	case "down":
		pixel, err = maze.GetNextNodeDown(meta.row, meta.rowPos)
		break
	case "left":
		pixel, err = maze.GetNextNodeLeft(meta.row, meta.rowPos)
		break
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else if pixel.IsDeadEnd {
		fmt.Printf("Encountered dead end (%d,%d)\n", meta.row, meta.rowPos)
		getNextDirection(meta)
		// makeMove(maze, meta)
	} else if pixel.IsNode {
		fmt.Printf("Pixel array coords (%d,%d)\n", pixel.Row, pixel.RowPos)
		meta.row = pixel.Row
		meta.rowPos = pixel.RowPos
		makeMove(maze, meta)
	}
	fmt.Print(pixel)
}

func getNextDirection(meta mazeMeta) {
	switch meta.direction {
	case "up":
		meta.direction = "right"
		break
	case "right":
		meta.direction = "down"
		break
	case "down":
		meta.direction = "left"
		break
	case "left":
		meta.direction = "up"
		break
	}
}
