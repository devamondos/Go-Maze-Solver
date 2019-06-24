package main

import (
	"log"

	"github.com/devamondos/maze_solver/mazesolvers"
	"github.com/devamondos/maze_solver/mazeutils"
)

func main() {
	debug := true
	log.Println("Reading image...")
	maze := mazeutils.Read("medium.png")
	log.Println("Analysing maze...")
	mazeutils.Analyse(maze, true)
	log.Println("Solving maze...")
	mazesolvers.Solve(maze, "alwaysLeft", debug)
	// TODO:
	// - Create map
	// - Solve map using simple algorithm (always turn left etc.)
	// - Print out route (maybe ASCII or directly ontop of image)

	// IMPROVEMENTS:
	// - Pass maze image location as a envrionment variable
}
