package binheap

import (
	"fmt"
	"testing"
)

func TestBinHeap(t *testing.T) {
	a := []int{9, 5, 6, 2, 3}
	a = heapify(a)
	fmt.Println(a)

}
