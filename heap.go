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

func (hp *heap) sink(rootIndex int) {
	if rootIndex >= hp.l {
		return
	}
	if 2 * rootIndex < hp.l {
		if hp.compare(hp.Arr[2 * rootIndex + 1], hp.arr[rootIndex]) {
			hp.swap(rootIndex, 2 * rootIndex + 1)
			hp.sink(2 * rootIndex + 1)
			return
		}
	} else if 2 * rootIndex + 1 < hp.l {
		if hp.compare(hp.Arr[2 * rootIndex + 2], hp.arr[rootIndex]) {
			hp.sink(rootIndex, 2 * rootIndex + 2)
			hp.sink(2 * rootIndex + 2)
			return
		}
	}
}

func (hp *heap) sift(childIndex int) {
	if childIndex <=0 {
		return
	}
	parentIndex := childIndex - 1 / 2
	if parentIndex >= 0 {
		if hp.compare(hp.arr[childIndex], hp.arr[parentIndex]) {
			hp.swap(childIndex, parentIndex)
			hp.sift(parentIndex)
			return
		}
	}
}

func(hp *heap) insert(val int) {
	if hp.l == hp.cap {
		fmt.Println("capacity exceeded, copying")
		hp.cap += 1
	}
	hp.arr = append(hp.arr, val)
	hp.l += 1
	hp.sift(hp.l - 1)
}

func(hp *heap) pop(index int) {
	hp.swap(0, l-1)
	hp.arr = hp.arr[:l-1]
	hp.sink(0)
}