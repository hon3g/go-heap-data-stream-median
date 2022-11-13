package main

import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	minH *Heap
	maxH *Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minH: new(Heap),
		maxH: new(Heap),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minH.Len() < this.maxH.Len() {
		heap.Push(this.minH, num)
	} else {
		heap.Push(this.maxH, -num)
	}
	if this.maxH.Len() > 0 && this.minH.Len() > 0 &&
		-(*this.maxH)[0] > (*this.minH)[0] {
		mi, mx := heap.Pop(this.minH), heap.Pop(this.maxH)
		heap.Push(this.maxH, -mi.(int))
		heap.Push(this.minH, -mx.(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minH.Len() < this.maxH.Len() {
		return -float64((*this.maxH)[0])
	}
	return (float64((*this.minH)[0]) - float64((*this.maxH)[0])) / 2
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(3)
	fmt.Println(mf.FindMedian())
}
