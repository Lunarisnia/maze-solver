package main

import (
	"maze-solver/internal/mazes"
)

func main() {
	maze, err := mazes.LoadMaze("/Users/louna/Desktop/Work/Personal/maze-solver/puzzles/puzzle_01.png")
	if err != nil {
		panic(err)
	}
	solved, err := mazes.Solve(maze)
	if err != nil {
		panic(err)
	}
	mazes.Save(maze, solved)
}
