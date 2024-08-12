package mazes

import (
	"fmt"
	"testing"
)

func Test_AStar(t *testing.T) {
	maze, err := LoadMaze("/Users/louna/Desktop/Work/Personal/maze-solver/puzzles/puzzle_01.png")
	if err != nil {
		panic(err)
	}
	aStar := NewAStar(maze)
	solved, err := Solve(aStar)
	if err != nil {
		panic(err)
	}

	for _, rows := range solved.World {
		for _, col := range rows {
			fmt.Printf("%v ", col)
		}
		fmt.Println()
	}
}
