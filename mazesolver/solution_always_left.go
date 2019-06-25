package mazesolver

import (
	"os"
	"time"

	mb "github.com/devamondos/maze_solver/mazebuilder"
	"github.com/golang/glog"
)

func alwaysLeft(maze *mb.Maze, debug bool) *MazeMeta {
	meta := &MazeMeta{
		row:          maze.StartRow,
		rowPos:       maze.StartRowPos,
		debug:        debug,
		direction:    2,
		directions:   [4]string{"up", "right", "down", "left"},
		Moves:        [][2]int{},
		IsSuccessful: false,
	}
	meta.addMove(meta.row, meta.rowPos)
	makeMove(maze, meta)
	return meta
}

func makeMove(maze *mb.Maze, meta *MazeMeta) {
	glog.V(2).Infof("At (%d,%d) ... going %s\n", meta.row, meta.rowPos, meta.directions[meta.direction])
	if meta.debug {
		time.Sleep(1 * time.Second)
	}
	nextPixel := checkNextPixel(maze, meta)
	var pixel *mb.Pixel
	var moves [][2]int
	var err error
	if nextPixel == nil {
		glog.Fatal("Error: could not find next pixel")
		os.Exit(1)
	}
	if nextPixel.IsEnd {
		glog.V(2).Infof("EXIT FOUND!! (%d, %d)", nextPixel.Row, nextPixel.RowPos)
		meta.addMove(nextPixel.Row, nextPixel.RowPos)
		meta.IsSuccessful = true
		return
	} else if !nextPixel.IsPath || nextPixel.IsDeadEnd {
		glog.V(2).Info("Next pixel is wall or dead end")
		if nextPixel.IsDeadEnd {
			meta.addMove(nextPixel.Row, nextPixel.RowPos)
		}
		changeDirection(maze, meta, nextPixel)
	} else if nextPixel.IsNode {
		glog.V(2).Info("Next pixel is node")
		setMazeLocation(meta, nextPixel)
		meta.addMove(meta.row, meta.rowPos)
		changeDirection(maze, meta, nextPixel)
	} else {
		switch meta.directions[meta.direction] {
		case "up":
			pixel, moves, err = maze.GetNextNodeUp(meta.row, meta.rowPos)
			break
		case "right":
			pixel, moves, err = maze.GetNextNodeRight(meta.row, meta.rowPos)
			break
		case "down":
			pixel, moves, err = maze.GetNextNodeDown(meta.row, meta.rowPos)
			break
		case "left":
			pixel, moves, err = maze.GetNextNodeLeft(meta.row, meta.rowPos)
			break
		}
		if err != nil {
			glog.Fatal(err)
			os.Exit(1)
		} else if pixel != nil {
			meta.addMoves(moves)
			if pixel.IsDeadEnd {
				glog.V(2).Infof("Encountered dead end (%d,%d)\n", pixel.Row, pixel.RowPos)
				changeDirection(maze, meta, pixel)
			} else if pixel.IsNode {
				glog.V(2).Infof("Node array pos (%d,%d)\n", pixel.Row, pixel.RowPos)
				setMazeLocation(meta, pixel)
				changeDirection(maze, meta, nextPixel)
			}
		}
	}
}

func setMazeLocation(meta *MazeMeta, pixel *mb.Pixel) {
	meta.row = pixel.Row
	meta.rowPos = pixel.RowPos
}

func checkNextPixel(maze *mb.Maze, meta *MazeMeta) *mb.Pixel {
	var nextRow, nextRowPos int
	switch meta.directions[meta.direction] {
	case "up":
		nextRow = meta.row - 1
		nextRowPos = meta.rowPos
		break
	case "right":
		nextRow = meta.row
		nextRowPos = meta.rowPos + 1
		break
	case "down":
		nextRow = meta.row + 1
		nextRowPos = meta.rowPos
		break
	case "left":
		nextRow = meta.row
		nextRowPos = meta.rowPos - 1
		break
	}
	nextPixel, _ := maze.GetPixel(nextRow, nextRowPos)
	return nextPixel
}

func changeDirection(maze *mb.Maze, meta *MazeMeta, pixel *mb.Pixel) {
	getNextDirection(meta, pixel)
	makeMove(maze, meta)
}

func getNextDirection(meta *MazeMeta, pixel *mb.Pixel) {
	if !pixel.IsPath {
		meta.direction++
	} else if pixel.IsDeadEnd {
		meta.direction++
	} else {
		meta.direction--
	}
	if meta.direction > 3 {
		meta.direction -= 4
	} else if meta.direction < 0 {
		meta.direction += 4
	}
}
