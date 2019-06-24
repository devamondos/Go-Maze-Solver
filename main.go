package main

import (
	"fmt"

	"github.com/devamondos/maze_solver/mazesolvers"
	"github.com/devamondos/maze_solver/mazeutils"
)

func main() {
	fmt.Println("\nReading image...")
	maze := mazeutils.Read("medium.png")
	fmt.Println("\nAnalysing maze...")
	mazeutils.Analyse(maze, true)
	fmt.Println("\nSolving maze...")
	mazesolvers.Solve(maze, "alwaysLeft")
	// TODO:
	// - Create map
	// - Solve map using simple algorithm (always turn left etc.)
	// - Print out route (maybe ASCII or directly ontop of image)

	// IMPROVEMENTS:
	// - Pass maze image location as a envrionment variable
}
