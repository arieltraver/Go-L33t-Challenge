type heap struct {
	cap int
	l int
	arr []int
	compare func
}

func gt(a int, b int) {
	return a > b
}
func lt(a int, b int) {
	return a < b
}

func (hp *maxHeap) swap(i1 int, i2 int) {
	temp := hp.arr[i1]
	hp.arr[i1] = heap.arr[i2]
	hp.arr[i2] = temp
}

func newMinHeap(cap int) *heap {
	arr := make([]int, 0, cap)
	return &heap(cap: cap, l:0, arr: arr, compare: gt)
}

func newMaxHeap(cap int) *heap {
	arr := make([]int, 0, cap)
	return &heap(cap: cap, l:0, arr:arr, compare:gt)
}

func (hp *heap) heapify(rootIndex int) {
	if rootIndex >= hp.l {
		return
	}
	if 2 * rootIndex < hp.l {
		if hp.arr[rootIndex] < hp.Arr[2 * rootIndex] {
			hp.swap(rootIndex, 2 * rootIndex)
			hp.heapify(2 * rootIndex)
			return
		}
	} else if 2 * rootIndex + 1 < hp.l {
		if hp.arr[rootIndex] < hp.Arr[2 * rootIndex + 1] {
			hp.swap(rootIndex, 2 * rootIndex + 1)
			hp.heapify(2 * rootIndex + 1)
			return
		}
	}
}

func(hp *maxHeap) insert(n int) {

}