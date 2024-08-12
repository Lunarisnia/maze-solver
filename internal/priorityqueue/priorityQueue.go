package pqueue

import (
	ds "maze-solver/internal/datastructures"
)

type Weight struct {
	Position  ds.Vector2
	StartCost float64
	EndCost   float64
	FinalCost float64
	Passable  bool
	Parent    *Weight
}

type PriorityQueue []*Weight

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].FinalCost < pq[j].FinalCost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(w any) {
	// Maybe we can check if exist here
	item := w.(*Weight)
	for _, weight := range *pq {
		if weight.Position.Equal(item.Position) {
			return
		}
	}
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	//item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
