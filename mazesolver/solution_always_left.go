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
	// lets first check if the next pixel is a wall
	nextPixelIsPath := isNextPixelPath(maze, meta)
	if !nextPixelIsPath {
		// we need to turn around
		glog.V(2).Info("Next pixel is wall")
		changeDirection(maze, meta, true)
	} else {
		// next pixel is not a wall let's carry on
		var node *mb.Pixel
		var moves [][2]int
		var err error
		switch meta.directions[meta.direction] {
		case "up":
			node, moves, err = maze.GetNextNodeUp(meta.row, meta.rowPos)
			break
		case "right":
			node, moves, err = maze.GetNextNodeRight(meta.row, meta.rowPos)
			break
		case "down":
			node, moves, err = maze.GetNextNodeDown(meta.row, meta.rowPos)
			break
		case "left":
			node, moves, err = maze.GetNextNodeLeft(meta.row, meta.rowPos)
			break
		}
		if err != nil {
			// there was an error getting the next node, we need to stop execution
			glog.Fatal(err)
			os.Exit(1)
		} else if node != nil {
			// we have found the next node
			meta.addMoves(moves)
			if node.IsEnd {
				// We have found the exit
				glog.V(2).Infof("EXIT FOUND!! (%d, %d)", node.Row, node.RowPos)
				meta.IsSuccessful = true
				return
			} else if node.IsDeadEnd {
				// it is a dead end so we don't want to set the location
				// this means it will resume execution from the last node
				glog.V(2).Infof("Encountered dead end (%d,%d)\n", node.Row, node.RowPos)
				changeDirection(maze, meta, true)
			} else {
				// this is a node that is not a dead end
				// we should set the location to here
				glog.V(2).Infof("Node array pos (%d,%d)\n", node.Row, node.RowPos)
				setMazeLocation(meta, node)
				changeDirection(maze, meta, false)
			}
		}
	}
}

func setMazeLocation(meta *MazeMeta, pixel *mb.Pixel) {
	meta.row = pixel.Row
	meta.rowPos = pixel.RowPos
}

func isNextPixelPath(maze *mb.Maze, meta *MazeMeta) bool {
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
	return nextPixel.IsPath
}

func changeDirection(maze *mb.Maze, meta *MazeMeta, isWall bool) {
	getNextDirection(meta, isWall)
	makeMove(maze, meta)
}

// will turn left unless we have encountered a wall
func getNextDirection(meta *MazeMeta, isWall bool) {
	if isWall {
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
