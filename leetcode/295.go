package leetcode

// 这题如果对标准库文件熟的话，用Java写会非常简单。借此题熟悉下Go的库（我只能说这写法很Go）

import "container/heap"

type MedianFinder struct {
	left  MaxHeap
	right MinHeap
}

type MaxHeap []int
type MinHeap []int

var _ heap.Interface = &MaxHeap{}
var _ heap.Interface = &MinHeap{}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{make([]int, 0), make([]int, 0)}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == 0 {
		heap.Push(&this.left, num)
		return
	}
	mid := this.left[0]
	if num > mid {
		heap.Push(&this.right, num)
	} else {
		heap.Push(&this.left, num)
	}
	if this.left.Len() > this.right.Len()+1 {
		v := heap.Pop(&this.left)
		heap.Push(&this.right, v)
	} else if this.left.Len() < this.right.Len() {
		v := heap.Pop(&this.right)
		heap.Push(&this.left, v)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return float64(this.left[0])
	} else {
		return float64(this.left[0]+this.right[0]) / 2.0
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
