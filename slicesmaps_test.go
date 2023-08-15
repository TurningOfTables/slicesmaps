package slicesmaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneSlice(t *testing.T) {
	var s = SliceEx{1, 2, 3}
	c := s.CloneSlice()
	assert.Equal(t, s, c)
}

func TestCompaceSlice(t *testing.T) {
	var s = SliceEx{1, 1, 1, 2, 3, 4, 4}
	c := s.CompactSlice()
	assert.Equal(t, SliceEx{1, 2, 3, 4}, c)
}

func TestCompareSlice(t *testing.T) {
	var s1 = SliceEx{1, 2, 3, 4}
	var s2 = SliceEx{1, 2, 3, 4}
	var s3 = SliceEx{5, 6, 7, 8}
	var s4 = SliceEx{1, 2, 3}

	// Comparison should return 0 for identical slices

	c1 := s1.CompareSlice(s2)
	assert.Equal(t, 0, c1)

	// Comparison should return -1 where the first slice is less than the second
	c2 := s1.CompareSlice(s3)
	assert.Equal(t, -1, c2)

	// Comparison should return 1 where the first slice is longer than the second
	c3 := s1.CompareSlice(s4)
	assert.Equal(t, 1, c3)
}

func TestContainsSlice(t *testing.T) {
	var s1 = SliceEx{5, 6, 7}
	c1 := s1.ContainsSlice(5)
	assert.True(t, c1)

	var s2 = SliceEx{5, 6, 7}
	c2 := s2.ContainsSlice(1)
	assert.False(t, c2)
}

func TestDeleteSlice(t *testing.T) {
	var s = SliceEx{1, 2, 3, 4, 5}
	d := s.DeleteSlice(0, 2)

	assert.Equal(t, SliceEx{3, 4, 5}, d)
}

func TestIndexSlice(t *testing.T) {
	// Item exists
	var s = SliceEx{1, 2, 3, 4, 5}
	i := s.IndexSlice(5)

	assert.Equal(t, 4, i)

	// Item does not exist
	i2 := s.IndexSlice(10)

	assert.Equal(t, -1, i2)
}

func TestInsertSlice(t *testing.T) {
	var s = SliceEx{1, 2, 4, 5}
	i := s.InsertSlice(2, 3) // Insert 3 at index 2

	assert.Equal(t, SliceEx{1, 2, 3, 4, 5}, i)
}

func TestIsSortedSlice(t *testing.T) {
	var s = SliceEx{1, 2, 3, 4, 5} // ascending order sorted slice
	i := s.IsSortedSlice()
	assert.True(t, i)

	var s2 = SliceEx{5, 4, 3, 2, 1} // not ascending order sorted slice
	i2 := s2.IsSortedSlice()
	assert.False(t, i2)
}

func TestMaxSlice(t *testing.T) {
	var s = SliceEx{1, 5, 10, 20}
	m := s.MaxSlice()
	assert.Equal(t, 20, m)
}

func TestMinSlice(t *testing.T) {
	var s = SliceEx{1, 5, 10, 20}
	m := s.MinSlice()
	assert.Equal(t, 1, m)
}

func TestReplaceSlice(t *testing.T) {
	var s = SliceEx{1, 2, 3, 4, 5}
	r := s.ReplaceSlice(0, 2, 9) // replace s[0:2] with value 9
	assert.Equal(t, SliceEx{9, 3, 4, 5}, r)
}

func TestReverseSlice(t *testing.T) {
	var s = SliceEx{1, 2, 3, 4, 5}
	s.ReverseSlice()
	assert.Equal(t, SliceEx{5, 4, 3, 2, 1}, s)
}

func TestSortSlice(t *testing.T) {
	var s = SliceEx{4, 2, 1, 3, 5}
	s.SortSlice()

	assert.Equal(t, SliceEx{1, 2, 3, 4, 5}, s)
}

func TestCloneMap(t *testing.T) {
	var m = MapEx{"foo": "bar"}
	c := m.CloneMap()

	assert.Equal(t, m, c)
}

func TestCopyMap(t *testing.T) {
	var m = MapEx{"foo": "bar"}
	var m1 = MapEx{"baz": "bob"}

	m.CopyMap(m1)
	assert.Equal(t, MapEx{"baz": "bob", "foo": "bar"}, m)

	var m2 = MapEx{"foo": "bar"}
	var m3 = MapEx{"foo": "alice"}
	m2.CopyMap(m3)
	assert.Equal(t, MapEx{"foo": "alice"}, m2)
}

func TestEqualMap(t *testing.T) {
	var m = MapEx{"foo": "bar"}
	var m1 = MapEx{"foo": "bar"}

	e := m.EqualMap(m1)
	assert.True(t, e)

	var m2 = MapEx{"baz": "bob"}
	e1 := m.EqualMap(m2)
	assert.False(t, e1)
}

func TestDeleteValueFromMap(t *testing.T) {
	var m = MapEx{"foo": "bar", "baz": "bob"}
	m.DeleteValueFromMap("bob")

	assert.Equal(t, MapEx{"foo": "bar"}, m)
}

func TestDeleteKeyFromMap(t *testing.T) {
	var m = MapEx{"foo": "bar", "baz": "bob"}
	m.DeleteKeyFromMap("baz")

	assert.Equal(t, MapEx{"foo": "bar"}, m)
}
