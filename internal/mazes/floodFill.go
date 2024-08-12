package mazes

type FloodFill struct {
	maze Maze
}

func NewFloodFill(maze *Maze) *FloodFill {
	return &FloodFill{
		maze: *maze,
	}
}

var directions = [][]int{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func isValidPoint(x int, y int, maze [][]int) bool {
	rows := len(maze)
	cols := len(maze[0])
	return x >= 0 && x < cols && y >= 0 && y < rows
}

func isSolved(cellValue int) bool {
	return cellValue == 3
}

func (f *FloodFill) Solve() (*Maze, error) {
	visited := make([][]int, len(f.maze.World))
	for i := range visited {
		visited[i] = make([]int, len(f.maze.World[i]))
	}

	solved := false

	var floodFill func(x int, y int)
	floodFill = func(x int, y int) {
		if solved {
			return
		}

		visited[y][x] = 4

		if isSolved(f.maze.World[y][x]) {
			solved = true
			return
		}
		for _, dir := range directions {
			row := y + dir[0]
			col := x + dir[1]
			if isValidPoint(col, row, f.maze.World) && visited[row][col] != 4 && f.maze.World[row][col] != 0 {
				floodFill(col, row)
			}
		}
	}

	solution := Maze{
		Start: f.maze.Start,
		End:   f.maze.End,
		World: visited,
	}

	for y := range f.maze.World {
		if solved {
			break
		}
		for x := range f.maze.World[y] {
			if f.maze.World[y][x] != 0 && visited[y][x] != 4 {
				floodFill(x, y)
			}
		}
	}

	return &solution, nil
}
