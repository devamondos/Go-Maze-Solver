package mazesolvers

import (
	"fmt"

	m "github.com/devamondos/maze_solver/mazeutils"
	"github.com/golang/glog"
)

// Solve the maze
func Solve(maze *m.Maze, solution string, debug bool) {
	switch solution {
	case "alwaysLeft":
		meta := alwaysLeft(maze, debug)
		printSolution(maze, meta)
		break
	}
}

func printSolution(maze *m.Maze, meta *MazeMeta) {
	glog.Info("Solution found:")
	glog.Infof(" - Number of moves: %d", len(meta.moves))
	glog.Infof(" - Time taken: %d\n\n", 0)

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
