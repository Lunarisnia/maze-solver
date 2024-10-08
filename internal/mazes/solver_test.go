package mazes

import (
	"fmt"
	"testing"
)

func Test_Solver(t *testing.T) {
	maze, err := LoadMaze("/Users/louna/Desktop/Work/Personal/maze-solver/puzzles/puzzle_01.png")
	if err != nil {
		t.Fail()
	}
	floodFill := NewFloodFill(maze)
	solved, err := Solve(floodFill)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for _, rows := range solved.World {
		for _, col := range rows {
			fmt.Printf("%v ", col)
		}

		fmt.Println()
	}
}
