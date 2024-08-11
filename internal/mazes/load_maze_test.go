package mazes

import (
	"fmt"
	"testing"
)

func Test_LoadMaze(t *testing.T) {
	maze, err := LoadMaze("/Users/louna/Desktop/Work/Personal/maze-solver/puzzles/puzzle_01.png")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%v ", maze[i][j])
		}

		fmt.Println()
	}
}
