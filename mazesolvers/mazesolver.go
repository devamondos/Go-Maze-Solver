package mazesolvers

import m "github.com/devamondos/maze_solver/mazeutils"

// Solve the maze
func Solve(maze *m.Maze, solution string) {
	switch solution {
	case "alwaysLeft":
		alwaysLeft(maze)
		break
	}
}
