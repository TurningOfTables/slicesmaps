package slicesmaps

import (
	"maps"
	"slices"
)

type SliceEx []int
type MapEx map[string]string

func (s SliceEx) CloneSlice() SliceEx {
	clonedSlice := slices.Clone(s)
	return clonedSlice
}

func (s SliceEx) CompactSlice() SliceEx {
	compactSlice := slices.Compact(s)
	return compactSlice
}

func (s1 SliceEx) CompareSlice(s2 SliceEx) int {
	comparison := slices.Compare(s1, s2)
	return comparison
}

func (s SliceEx) ContainsSlice(n int) bool {
	contains := slices.Contains(s, n)
	return contains
}

func (s SliceEx) DeleteSlice(a, b int) SliceEx {
	deleted := slices.Delete(s, a, b)
	return deleted
}

func (s SliceEx) IndexSlice(i int) int {
	index := slices.Index(s, i)
	return index
}

func (s SliceEx) InsertSlice(i, v int) SliceEx {
	newSlice := slices.Insert(s, i, v)
	return newSlice
}

func (s SliceEx) IsSortedSlice() bool {
	isSorted := slices.IsSorted(s)
	return isSorted
}

func (s SliceEx) MaxSlice() int {
	max := slices.Max(s)
	return max
}

func (s SliceEx) MinSlice() int {
	min := slices.Min(s)
	return min
}

func (s SliceEx) ReplaceSlice(i, j, v int) SliceEx {
	newSlice := slices.Replace(s, i, j, v)
	return newSlice
}

func (s SliceEx) ReverseSlice() {
	slices.Reverse(s)
}

func (s SliceEx) SortSlice() {
	slices.Sort(s)
}

func (m MapEx) CloneMap() MapEx {
	c := maps.Clone(m)
	return c
}

func (m MapEx) CopyMap(m1 MapEx) {
	maps.Copy(m, m1)
}

func (m MapEx) DeleteValueFromMap(n string) {

	maps.DeleteFunc(m, func(k, v string) bool {
		return v == n
	})
}

func (m MapEx) DeleteKeyFromMap(n string) {
	maps.DeleteFunc(m, func(k, v string) bool {
		return k == n
	})
}

func (m MapEx) EqualMap(m1 MapEx) bool {
	e := maps.Equal(m, m1)
	return e
}
