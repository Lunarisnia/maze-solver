package models

import "maze-solver/internal/mazes"

type SolverAlgo string

const (
	FloodFill SolverAlgo = "FLOOD_FILL"
	AStar     SolverAlgo = "A_STAR"
)

type UserMaze struct {
	Start     []int      `json:"start"`
	End       []int      `json:"end"`
	World     [][]int    `json:"world"`
	Algorithm SolverAlgo `json:"algorithm"`
}

func (u UserMaze) ConvertToMaze() *mazes.Maze {
	maze := mazes.Maze{
		Start: u.Start,
		End:   u.End,
		World: u.World,
	}

	return &maze
}
