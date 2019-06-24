
# Golang Maze Solver

*Disclaimer: Inspired by Computerphiles video about maze solving - https://www.youtube.com/watch?v=rop0W4QDOUI*

I wanted to learn Go so thought that this would be a nice and simple starting point.

This project comes with two mazes in the form of PNG images. These are generated using [this generator](https://keesiemeijer.github.io/maze-generator/). However, this should work with any image providing that:

- The image is in black and transparent (PNG). Black being the walls.
- The maze entrance is at the top and the maze exit is at the bottom.
- The wall and path widths are one pixel wide

Currently this application will only solve the maze using an 'always left' strategy (I'm learning Go, not maze solving strategies) but it can be extended easily.

To run the application - ```go run main.go```

To test the application - (Unfinished)
