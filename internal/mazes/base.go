package mazes

type Algorithm interface {
	Solve() (*Maze, error)
}

