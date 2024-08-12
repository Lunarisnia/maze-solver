package mazes

import (
	"container/heap"
	ds "maze-solver/internal/datastructures"
	pqueue "maze-solver/internal/priorityqueue"
)

type AStar struct {
	maze Maze
}

func NewAStar(maze *Maze) *AStar {
	return &AStar{
		maze: *maze,
	}
}

func weightWorld(maze Maze) [][]pqueue.Weight {
	start := ds.Vector2{maze.Start[0], maze.Start[1]}
	end := ds.Vector2{maze.End[0], maze.End[1]}
	weightedWorld := make([][]pqueue.Weight, len(maze.World))
	for i := range weightedWorld {
		weightedWorld[i] = make([]pqueue.Weight, len(maze.World[0]))
	}
	for i, rows := range maze.World {
		for j := range rows {
			weight := pqueue.Weight{
				Position:  ds.Vector2{i, j},
				StartCost: 0,
				EndCost:   0,
				FinalCost: 0,
				Passable:  false,
			}
			if rows[j] != 0 {
				weight.StartCost = start.MoveTo(ds.Vector2{i, j}).Length()
				weight.EndCost = end.MoveTo(ds.Vector2{i, j}).Length()
				weight.FinalCost = weight.StartCost + weight.EndCost
				weight.Passable = true
				//fmt.Printf("(%v, %v) To (%v, %v) is (%v, %v) or length of %v\n", i, j, arbitraryTarget[0], arbitraryTarget[1], moveToEnd[0], moveToEnd[1], distanceToEnd)
			}

			weightedWorld[i][j] = weight
		}
	}
	return weightedWorld
}

type Node struct {
	Parent *Node
	Value  *pqueue.Weight
}

func (a *AStar) Solve() (*Maze, error) {
	weightedWorld := weightWorld(a.maze)
	startWeight := weightedWorld[a.maze.Start[0]][a.maze.Start[1]]
	minHeap := pqueue.PriorityQueue{&startWeight}
	heap.Init(&minHeap)

	found := false
	visited := make([][]bool, len(a.maze.World))
	for i := range visited {
		visited[i] = make([]bool, len(a.maze.World[0]))
	}
	var pointer *pqueue.Weight
	var findPath func(w *pqueue.Weight)
	findPath = func(w *pqueue.Weight) {
		if found {
			return
		}
		if w.Position.Equal(a.maze.End) {
			found = true
			pointer = w
			return
		}
		visited[w.Position[0]][w.Position[1]] = true
		for _, dir := range directions {
			row := w.Position[0] + dir[0]
			col := w.Position[1] + dir[1]

			if isValidPoint(col, row, a.maze.World) && !visited[row][col] {
				nW := weightedWorld[row][col]
				if nW.Passable {
					nW.Parent = w
					heap.Push(&minHeap, &nW)
				}
			}
		}
		if len(weightedWorld) > len(weightedWorld)-1 {
			findPath(heap.Pop(&minHeap).(*pqueue.Weight))
		}
	}
	findPath(heap.Pop(&minHeap).(*pqueue.Weight))

	safePath := make([]ds.Vector2, 0)
	var getPath func(p *pqueue.Weight)
	getPath = func(p *pqueue.Weight) {
		if p != nil {
			safePath = append(safePath, p.Position)
			getPath(p.Parent)
		}
		return
	}
	getPath(pointer)
	solution := Maze{
		Start: a.maze.Start,
		End:   a.maze.End,
		World: make([][]int, len(a.maze.World)),
	}
	for i := range solution.World {
		solution.World[i] = make([]int, len(a.maze.World[0]))
	}
	for i := len(safePath) - 1; i >= 0; i-- {
		move := safePath[i]
		solution.World[move[0]][move[1]] = 4
	}

	return &solution, nil
}
