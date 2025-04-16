package main

import "testing"

func TestArrayList(t *testing.T) {
	t.Run("Length", func(t *testing.T) {
		list := &ArrayList{}
		if list.Length() != 0 {
			t.Error("Empty list should have length 0")
		}

		list.Append('a')
		if list.Length() != 1 {
			t.Error("Length should be 1 after append")
		}
	})

	t.Run("Append", func(t *testing.T) {
		list := &ArrayList{}
		list.Append('a')
		if list.Get(0) != 'a' || list.Length() != 1 {
			t.Error("Append failed")
		}
	})

	t.Run("Insert", func(t *testing.T) {
		list := &ArrayList{}
		list.Insert('a', 0) // Valid
		list.Insert('b', 1) // Valid
		list.Insert('c', 1) // Middle

		if list.Get(1) != 'c' || list.Length() != 3 {
			t.Error("Insert failed")
		}

		// Test panic
		assertPanic(t, func() { list.Insert('d', -1) })
	})
	t.Run("Delete", func(t *testing.T) {
		list := &ArrayList{}
		list.Append('a')
		list.Append('b')

		deleted := list.Delete(0)
		if deleted != 'a' || list.Length() != 1 {
			t.Error("Delete failed")
		}

		assertPanic(t, func() { list.Delete(2) })
	})

	t.Run("DeleteAll", func(t *testing.T) {
		list := &ArrayList{}
		list.Append('a')
		list.Append('b')
		list.Append('a')

		list.DeleteAll('a')
		if list.Length() != 1 || list.Get(0) != 'b' {
			t.Error("DeleteAll failed")
		}

		list.DeleteAll('x') // No elements
		if list.Length() != 1 {
			t.Error("DeleteAll should do nothing")
		}
	})

	t.Run("Get", func(t *testing.T) {
		list := &ArrayList{}
		list.Append('a')

		if list.Get(0) != 'a' {
			t.Error("Get failed")
		}

		assertPanic(t, func() { list.Get(1) })
	})

	t.Run("Clone", func(t *testing.T) {
		list := &ArrayList{}
		list.Append('a')
		cloned := list.Clone()

		cloned.Append('b')
		if list.Length() != 1 || cloned.Length() != 2 {
			t.Error("Clone failed")
		}
	})

	t.Run("Reverse", func(t *testing.T) {
		list := &ArrayList{}
		list.Append('a')
		list.Append('b')
		list.Append('c')

		list.Reverse()
		if list.Get(0) != 'c' || list.Get(2) != 'a' {
			t.Error("Reverse failed")
		}
	})

	t.Run("FindFirst/FindLast", func(t *testing.T) {
		list := &ArrayList{}
		list.Append('a')
		list.Append('b')
		list.Append('a')

		if list.FindFirst('a') != 0 || list.FindLast('a') != 2 {
			t.Error("Find methods failed")
		}

		if list.FindFirst('x') != -1 {
			t.Error("Find non-existing failed")
		}
	})

	t.Run("Clear", func(t *testing.T) {
		list := &ArrayList{}
		list.Append('a')
		list.Clear()

		if list.Length() != 0 {
			t.Error("Clear failed")
		}
	})

	t.Run("Extend", func(t *testing.T) {
		list1 := &ArrayList{}
		list2 := &ArrayList{}
		list1.Append('a')
		list2.Append('b')

		list1.Extend(list2)
		if list1.Length() != 2 || list1.Get(1) != 'b' {
			t.Error("Extend failed")
		}

		list2.Append('c')
		if list1.Length() != 2 { // Should be independent
			t.Error("Extend should create copy")
		}
	})
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic didn't occur")
		}
	}()
	f()
}
