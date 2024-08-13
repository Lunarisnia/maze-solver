package mazes

type Algorithm interface {
	Solve() (*Maze, error)
}

type UnsolvableError struct {
}

func (e UnsolvableError) Error() string {
	return "maze is unsolvable"
}
