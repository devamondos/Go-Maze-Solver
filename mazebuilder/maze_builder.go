package mazeutils

// Build maze from image and return pixel array
func Build(filePath string, debug bool) *Maze {
	maze := read(filePath, debug)
	analyse(maze, debug)
	return maze
}
