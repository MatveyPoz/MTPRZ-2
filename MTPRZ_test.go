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
		if list.elements[0] != 'a' || list.Length() != 1 {
			t.Error("Append failed")
		}
	})

	t.Run("Insert", func(t *testing.T) {
		list := &ArrayList{}
		list.Insert('a', 0) // Valid
		list.Insert('b', 1) // Valid
		list.Insert('c', 1) // Middle

		if list.elements[1] != 'c' || list.Length() != 3 {
			t.Error("Insert failed")
		}

		// Test panic
		assertPanic(t, func() { list.Insert('d', -1) })
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
