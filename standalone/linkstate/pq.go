package main

//
//type PriorityQueue []*LinkState
//
//func (pq PriorityQueue) Len() int { return len(pq) }
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	return pq[i].Cost < pq[j].Cost
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//}
//
//func (pq *PriorityQueue) Push(x interface{}) {
//	item := x.(*LinkState)
//	*pq = append(*pq, item)
//}
//
//func (pq *PriorityQueue) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	item := old[n-1]
//	old[n-1] = nil // avoid memory leak
//	*pq = old[0 : n-1]
//	return item
//}
