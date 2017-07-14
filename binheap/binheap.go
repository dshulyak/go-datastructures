package binheap

import "fmt"

func down(heap []int, i int) {
	size := len(heap) - 1
	for i*2 <= size {
		min := min(heap, i)
		if heap[min] < heap[i] {
			fmt.Println("vals", heap[min], heap[i])
			heap[i], heap[min] = heap[min], heap[i]
		}
		i = min
		fmt.Println(i, heap)
	}
}

func up(heap []int, size int) {
	for size/2 > 0 {
		if heap[size] < heap[size/2] {
			heap[size], heap[size/2] = heap[size/2], heap[size]
		}
		size = size / 2
	}
}

func insert(heap []int, val int) []int {
	heap = append(heap, val)
	size := len(heap) - 1
	up(heap, size)
	return heap
}

func pop(heap []int) ([]int, int) {
	ret := heap[1]
	lastInx := len(heap) - 1
	heap[1] = heap[lastInx]
	heap = append(heap[:lastInx], []int{}...)
	down(heap, 1)
	return heap, ret
}

func min(heap []int, i int) int {
	doable := i * 2
	lth := len(heap) - 1
	if doable+1 > lth {
		return doable
	}
	if heap[doable] > heap[doable+1] {
		fmt.Println("min vals", heap[doable], heap[doable+1])
		return doable + 1
	}
	return doable

}

func heapify(heap []int) []int {
	i := len(heap) / 2
	heap = append(heap, 0)
	heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
	for i > 0 {
		down(heap, i)
		i--
	}
	return heap
}
