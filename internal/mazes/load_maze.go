package mazes

import (
	"image"
	_ "image/png"
	"os"
)

func LoadMaze(mazePath string) ([][]int, error) {
	file, err := os.Open(mazePath)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	maze := make([][]int, img.Bounds().Max.Y)
	for i := range maze {
		maze[i] = make([]int, img.Bounds().Max.X)
	}
	for y := range img.Bounds().Max.Y {
		for x := range img.Bounds().Max.X {
			r, g, b, _ := img.At(x, y).RGBA()
			r /= 256
			g /= 256
			b /= 256

			r /= 255
			g /= 255
			b /= 255

			cellColor := 0
			if r == 1 && g == 0 && b == 0 {
				cellColor = 2
			} else if r == 1 && g == 1 && b == 1 {
				cellColor = 1
			} else if r == 0 && g == 1 && b == 0 {
				cellColor = 3
			}
			maze[y][x] = cellColor
		}
	}

	return maze, nil
}
