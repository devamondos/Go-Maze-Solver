package main

import (
	"flag"
	"os"

	mb "github.com/devamondos/maze_solver/mazebuilder"
	ms "github.com/devamondos/maze_solver/mazesolver"
)

const debug bool = false
const solution string = "alwaysLeft"

// const image string = "tiny.png"
const image string = "medium.png"

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	var logLevel string
	if debug {
		logLevel = "2"
	} else {
		logLevel = "1"
	}
	flag.Usage = usage
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Set("v", logLevel)
	flag.Parse()
}

func main() {
	maze := mb.Build(image, debug)
	ms.Solve(maze, solution, debug)
}
