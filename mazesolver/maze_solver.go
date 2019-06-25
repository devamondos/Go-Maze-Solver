package mazesolver

import (
	"fmt"
	"time"

	mb "github.com/devamondos/maze_solver/mazebuilder"
	"github.com/golang/glog"
)

// Solve the maze
func Solve(maze *mb.Maze, solution string, debug bool) *MazeMeta {
	glog.Info("Solving maze...")
	start := time.Now()
	var meta *MazeMeta

	switch solution {
	case "alwaysLeft":
		meta = alwaysLeft(maze, debug)
		break
	}

	elapsed := time.Since(start)
	glog.Info("Solution found:")
	glog.Infof(" - Number of moves: %d", len(meta.Moves))
	glog.Infof(" - Time taken: %s", elapsed)
	if debug {
		printSolution(maze, meta)
	}

	return meta
}

func printSolution(maze *mb.Maze, meta *MazeMeta) {
	for row := 0; row < maze.Rows; row++ {
		for rowPos := 0; rowPos < maze.RowLength; rowPos++ {
			if meta.containsMove(row, rowPos) {
				fmt.Print("~")
			} else {
				pixel, err := maze.GetPixel(row, rowPos)
				if err != nil {
					glog.Fatal(err)
				} else {
					if pixel.IsPath {
						fmt.Print(" ")
					} else {
						fmt.Print("@")
					}
				}
			}

		}
		fmt.Print("\n")
	}
}
