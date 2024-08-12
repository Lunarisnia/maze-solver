package mazes

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const maxUint16 = 256*256 - 1

func colorCell(maze *Maze, dest *image.RGBA, additive bool) {
	for y, rows := range maze.World {
		for x, col := range rows {
			if additive && ((maze.Start[0] == y && maze.Start[1] == x) || (maze.End[0] == y && maze.End[1] == x)) {
				continue
			}
			if additive && maze.World[y][x] != 4 {
				continue
			}
			cellColor := color.RGBA64{
				R: 0,
				G: 0,
				B: 0,
				A: maxUint16,
			}

			if col == 2 {
				cellColor.R = maxUint16
				cellColor.G = 0
				cellColor.B = 0
			} else if col == 1 {
				cellColor.R = maxUint16
				cellColor.G = maxUint16
				cellColor.B = maxUint16
			} else if col == 3 {
				cellColor.R = 0
				cellColor.G = maxUint16
				cellColor.B = 0
			} else if col == 4 {
				cellColor.R = 0
				cellColor.G = 0
				cellColor.B = maxUint16
			}

			dest.Set(x, y, cellColor)
		}
	}

}

func Save(originalMaze *Maze, solution *Maze) *image.RGBA {
	maze := image.NewRGBA(image.Rect(0, 0, len(solution.World[0]), len(solution.World)))

	colorCell(originalMaze, maze, false)
	colorCell(solution, maze, true)

	f, err := os.Create("/Users/louna/Desktop/Work/Personal/maze-solver/solved/maze.png")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, maze)
	if err != nil {
		panic(err)
	}

	return maze
}
