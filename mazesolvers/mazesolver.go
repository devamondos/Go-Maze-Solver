package mazesolvers

import (
	"fmt"
	"log"

	m "github.com/devamondos/maze_solver/mazeutils"
)

// Solve the maze
func Solve(maze *m.Maze, solution string) {
	switch solution {
	case "alwaysLeft":
		meta := alwaysLeft(maze)
		printSolution(maze, meta)
		break
	}
}

func printSolution(maze *m.Maze, meta *MazeMeta) {
	fmt.Printf("\n\n - Number of moves: %d\n\n", len(meta.moves))
	// fmt.Printf(" - Time taken: %d\n\n", )

	for row := 0; row < maze.Rows; row++ {
		for rowPos := 0; rowPos < maze.RowLength; rowPos++ {
			if meta.containsMove(row, rowPos) {
				fmt.Print("~")
			} else {
				pixel, err := maze.GetPixel(row, rowPos)
				if err != nil {
					log.Fatal(err)
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
