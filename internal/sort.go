package internal

func Unique(l []int64) []int64 {
	result := make([]int64, 0, len(l))

	lSet := make(map[int64]struct{}, len(l))
	for _, el := range l {
		if _, ok := lSet[el]; !ok {
			result = append(result, el)
		}

		lSet[el] = struct{}{}
	}

	return result
}

func IsEqual(l1, l2 []int64) bool {
	if len(l1) != len(l2) {
		return false
	}

	l1Set := make(map[int64]struct{}, len(l1))
	for _, el := range l1 {
		l1Set[el] = struct{}{}
	}

	for _, el := range l2 {
		if _, ok := l1Set[el]; !ok {
			return false
		}
	}

	return true
}

type Updater func([]int64) []int64

type Slice struct {
	SliceNums []int64
}

func (m *Slice) UpdateSlice(updater Updater) (hasChange bool) {
	updated := Unique(updater(m.SliceNums))
	if !IsEqual(updated, m.SliceNums) {
		hasChange = true
	}

	m.SliceNums = updated

	return
}

func isArraysIntersects(a1, a2 []int64) bool {
	h := make(map[int64]bool)
	for _, e := range a1 {
		h[e] = true
	}

	for _, e := range a2 {
		if h[e] {
			return true
		}
	}

	return false
}
