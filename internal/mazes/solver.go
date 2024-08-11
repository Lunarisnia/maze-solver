package mazes

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

func Solve(maze [][]int) ([][]int, error) {
	visited := make([][]int, len(maze))
	for i := range visited {
		visited[i] = make([]int, len(maze[i]))
	}

	solved := false

	var floodFill func(x int, y int)
	floodFill = func(x int, y int) {
		if solved {
			return
		}

		visited[y][x] = 4
		if isSolved(maze[y][x]) {
			solved = true
			return
		}
		for _, dir := range directions {
			row := y + dir[0]
			col := x + dir[1]
			if isValidPoint(col, row, maze) && visited[row][col] != 4 && maze[row][col] != 0 {
				floodFill(col, row)
			}
		}
	}

	for y := range maze {
		for x := range maze[y] {
			if solved {
				return visited, nil
			}
			if maze[y][x] != 0 && visited[y][x] != 4 {
				floodFill(x, y)
			}
		}
	}

	return [][]int{}, nil
}
