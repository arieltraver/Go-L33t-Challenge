//slice operations (from https://github.com/golang/go/wiki/SliceTricks) on int slices.
//to help with data structure implementations
//because apparently good Gogrammers use slices for everything
//and honestly
//yeah

func appendSlices(a []int, b []int) {
  a = append(a, b...)
}

func copySlice(a []int) []int {
  b := make([]int, len(a))
  copy(b, a)
  return b
}

func cutSlice(a []int, start int, end int) {
  a = append(a[:i], a[j:]...)
}

func delete(a []int, i int) {
  a = append(a[:i], a[i+1:]...)
}

func popBack(a []int) int {
  n := a[-1]
  a = a[:len(a) - 1]
  return n
}

func pushBack(a []int, val int) {
  a.append(val)
}

func fastDelete(a []int, i int) {
  a[i] = a[-1]
  a = a[:a.len - 1]
}
