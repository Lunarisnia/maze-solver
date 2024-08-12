package mazes

func Solve(algo Algorithm) (*Maze, error) {
	solution, err := algo.Solve()
	if err != nil {
		return nil, err
	}

	return solution, nil
}
