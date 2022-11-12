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
	minH := NewHeap(func(a, b float64) bool {
		return a < b
	})
	maxH := NewHeap(func(a, b float64) bool {
		return a > b
	})
	return MedianFinder{
		minH: minH,
		maxH: maxH,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.minH, float64(num))
	heap.Push(this.maxH, heap.Pop(this.minH))
	if this.maxH.Len() > this.minH.Len()+1 {
		heap.Push(this.minH, heap.Pop(this.maxH))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minH.Len() < this.maxH.Len() {
		return this.maxH.Peek()
	}
	return (this.minH.Peek() + this.maxH.Peek()) / 2
}

type Heap struct {
	Values   []float64
	LessFunc func(float64, float64) bool
}

func NewHeap(less func(float64, float64) bool) *Heap {
	return &Heap{LessFunc: less}
}

func (h Heap) Len() int {
	return len(h.Values)
}
func (h Heap) Less(i, j int) bool {
	return h.LessFunc(h.Values[i], h.Values[j])
}
func (h Heap) Swap(i, j int) {
	h.Values[i], h.Values[j] = h.Values[j], h.Values[i]
}

func (h *Heap) Push(x interface{}) {
	h.Values = append(h.Values, x.(float64))
}

func (h *Heap) Peek() float64 {
	return h.Values[0]
}

func (h *Heap) Pop() (x interface{}) {
	h.Values, x = h.Values[:h.Len()-1], h.Values[h.Len()-1]
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
