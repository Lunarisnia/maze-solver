package mazes

import (
	"fmt"
	"testing"
)

func Test_Save(t *testing.T) {
	maze, err := LoadMaze("/Users/louna/Desktop/Work/Personal/maze-solver/puzzles/puzzle_01.png")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	floodFill := NewFloodFill(maze)
	solved, err := Solve(floodFill)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	saved := Save(maze, solved)

	for i := range saved.Bounds().Max.Y {
		for j := range saved.Bounds().Max.X {
			r, g, b, a := saved.At(j, i).RGBA()
			r /= 256
			g /= 256
			b /= 256
			a /= 256

			r /= 255
			g /= 255
			b /= 255
			a /= 255

			cellColor := 0
			if r == 1 && g == 0 && b == 0 {
				cellColor = 2
			} else if r == 1 && g == 1 && b == 1 {
				cellColor = 1
			} else if r == 0 && g == 1 && b == 0 {
				cellColor = 3
			} else if r == 0 && g == 0 && b == 1 {
				cellColor = 4
			}
			fmt.Printf("%v ", cellColor)
		}

		fmt.Println()
	}
}
