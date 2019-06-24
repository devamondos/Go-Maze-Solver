package mazesolvers

import (
	"log"
	"os"
	"strconv"
	"time"

	m "github.com/devamondos/maze_solver/mazeutils"
)

// MazeMeta holds meta data for the solution taken
type MazeMeta struct {
	row        int
	rowPos     int
	debug      bool
	direction  int
	directions [4]string
	moves      []string
}

func (m *MazeMeta) addMove(row int, rowPos int) {
	m.moves = append(m.moves, strconv.Itoa(row)+","+strconv.Itoa(rowPos))
}

func (m *MazeMeta) addMoves(moves [][2]int) {
	for _, move := range moves {
		m.addMove(move[0], move[1])
	}
}

func (m *MazeMeta) containsMove(row int, rowPos int) bool {
	strPos := strconv.Itoa(row) + "," + strconv.Itoa(rowPos)
	for _, a := range m.moves {
		if a == strPos {
			return true
		}
	}
	return false
}

func alwaysLeft(maze *m.Maze, debug bool) *MazeMeta {
	meta := &MazeMeta{
		row:        maze.StartRow,
		rowPos:     maze.StartRowPos,
		debug:      debug,
		direction:  2,
		directions: [4]string{"up", "right", "down", "left"},
		moves:      []string{},
	}
	makeMove(maze, meta)
	return meta
}

func makeMove(maze *m.Maze, meta *MazeMeta) {
	log.Printf("At (%d,%d) ... going %s\n", meta.row, meta.rowPos, meta.directions[meta.direction])
	if meta.debug {
		time.Sleep(1 * time.Second)
	}
	nextPixel := checkNextPixel(maze, meta)
	var pixel *m.Pixel
	var moves [][2]int
	var err error
	if nextPixel == nil {
		log.Printf("Error: could not find next pixel")
		os.Exit(1)
	}
	if nextPixel.IsEnd {
		log.Printf("EXIT FOUND!! FUCK YEAH!! (%d, %d)", nextPixel.Row, nextPixel.RowPos)
		meta.addMove(meta.row, meta.rowPos)
		return
	} else if !nextPixel.IsPath || nextPixel.IsDeadEnd {
		log.Println("Next pixel is wall or dead end")
		changeDirection(maze, meta, nextPixel)
	} else if nextPixel.IsNode {
		log.Println("Next pixel is node")
		meta.addMove(meta.row, meta.rowPos)
		setMazeLocation(meta, nextPixel)
		changeDirection(maze, meta, nextPixel)
	} else {
		// log.Printf("%d,%d - ", meta.row, meta.rowPos)
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
	}

	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else if pixel != nil {
		moves = append(moves, [2]int{meta.row, meta.rowPos})
		meta.addMoves(moves)
		if pixel.IsDeadEnd {
			log.Printf("Encountered dead end (%d,%d)\n", pixel.Row, pixel.RowPos)
			changeDirection(maze, meta, pixel)
		} else if pixel.IsNode {
			log.Printf("Node array pos (%d,%d)\n", pixel.Row, pixel.RowPos)
			setMazeLocation(meta, pixel)
			changeDirection(maze, meta, nextPixel)
		}
	}
}

func setMazeLocation(meta *MazeMeta, pixel *m.Pixel) {
	meta.row = pixel.Row
	meta.rowPos = pixel.RowPos
}

func checkNextPixel(maze *m.Maze, meta *MazeMeta) *m.Pixel {
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

func changeDirection(maze *m.Maze, meta *MazeMeta, pixel *m.Pixel) {
	getNextDirection(meta, pixel)
	makeMove(maze, meta)
}

func getNextDirection(meta *MazeMeta, pixel *m.Pixel) {
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
