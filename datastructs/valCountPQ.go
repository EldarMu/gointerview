package datastructs

import "container/heap"

// this is almost verbatim from the godoc examples page
// https://golang.org/pkg/container/heap/
// with minor adjustments for data type
type ValCount struct {
	Value int
	Count int
	Index int
}

type PriorityQueue []*ValCount

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Count >= pq[j].Count }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*ValCount)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *ValCount, count int) {
	item.Count = count
	heap.Fix(pq, item.Index)
}
