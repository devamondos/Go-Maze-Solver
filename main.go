package main

import (
	"flag"
	"os"

	"github.com/devamondos/maze_solver/mazesolvers"
	"github.com/devamondos/maze_solver/mazeutils"
	"github.com/golang/glog"
)

const debug bool = false

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
	glog.Info("Reading image...")
	maze := mazeutils.Read("medium.png")
	glog.Info("Analysing maze...")
	mazeutils.Analyse(maze, debug)
	glog.Info("Solving maze...")
	mazesolvers.Solve(maze, "alwaysLeft", debug)
}
