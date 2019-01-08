package datastructs

import "container/heap"

// this is almost verbatim from the godoc examples page
// https://golang.org/pkg/container/heap/
// with minor adjustments for data type

// ValCount is a container for int values and their corresponding counts in some sort of array
type ValCount struct {
	Value int
	Count int
	Index int
}

// PriorityQueue is a data structure for working with a maxheap
type PriorityQueue []*ValCount

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Count >= pq[j].Count }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

//Push adds a new ValCount onto the maxheap
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*ValCount)
	item.Index = n
	*pq = append(*pq, item)
}

//Pop removes the element with the highest count
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

//Update adjusts the order in the heap without having to remove and reinsert the element
func (pq *PriorityQueue) Update(item *ValCount, count int) {
	item.Count = count
	heap.Fix(pq, item.Index)
}
