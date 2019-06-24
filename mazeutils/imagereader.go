package mazeutils

import (
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

// Read image and return pixel array
func Read(filePath string) *Maze {
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(filePath)

	if err != nil {
		log.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	maze, err := getMaze(file)

	if err != nil {
		log.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	return maze
}

// Get the bi-dimensional pixel array
func getMaze(file io.Reader) (*Maze, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var rowPixels []Pixel
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rowPixels = append(rowPixels, rgbaToPixel(r, g, b, a, y, x))
		}
		pixels = append(pixels, rowPixels)
	}

	maze := &Maze{
		pixels: pixels,
	}

	return maze, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32, row int, rowPos int) Pixel {
	return Pixel{
		rgba: Rgba{
			int(r / 257),
			int(g / 257),
			int(b / 257),
			int(a / 257),
		},
		Row:    row,
		RowPos: rowPos,
	}
}
