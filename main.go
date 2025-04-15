package main

import "fmt"

// ArrayList structure which realizes list of runes
type ArrayList struct {
	elements []rune
}

// Length return amount of elements
func (a *ArrayList) Length() int {
	return len(a.elements)
}

// Append adds element to the end of list
func (a *ArrayList) Append(element rune) {
	a.elements = append(a.elements, element)
}

// Insert insert element to speific position
func (a *ArrayList) Insert(element rune, index int) {
	if index < 0 || index > a.Length() {
		panic("invalid index for insert")
	}
	a.elements = append(a.elements[:index], append([]rune{element}, a.elements[index:]...)...)
}

// Delete delete element by index and return it
func (a *ArrayList) Delete(index int) rune {
	if index < 0 || index >= a.Length() {
		panic("invalid index for delete")
	}
	deleted := a.elements[index]
	a.elements = append(a.elements[:index], a.elements[index+1:]...)
	return deleted
}

// DeleteAll delete all occurrences of element
func (a *ArrayList) DeleteAll(element rune) {
	newElements := make([]rune, 0, a.Length())
	for _, e := range a.elements {
		if e != element {
			newElements = append(newElements, e)
		}
	}
	a.elements = newElements
}

// Get return element by index
func (a *ArrayList) Get(index int) rune {
	if index < 0 || index >= a.Length() {
		panic("invalid index for get")
	}
	return a.elements[index]
}

// Clone create copy of list
func (a *ArrayList) Clone() *ArrayList {
	newElements := make([]rune, len(a.elements))
	copy(newElements, a.elements)
	return &ArrayList{elements: newElements}
}

// Reverse reverse order of elements in list
func (a *ArrayList) Reverse() {
	for i := 0; i < len(a.elements)/2; i++ {
		j := len(a.elements) - i - 1
		a.elements[i], a.elements[j] = a.elements[j], a.elements[i]
	}
}

// FindFirst return first element oqqurence in list index or -1
func (a *ArrayList) FindFirst(element rune) int {
	for i, e := range a.elements {
		if e == element {
			return i
		}
	}
	return -1
}

// FindLast return last element occurrence in list index or -1
func (a *ArrayList) FindLast(element rune) int {
	for i := len(a.elements) - 1; i >= 0; i-- {
		if a.elements[i] == element {
			//uncomment to pass tests
			//return i
		}
	}
	return -1
}

// Clear clean list
func (a *ArrayList) Clear() {
	a.elements = nil
}

// Extend add elements from another list
func (a *ArrayList) Extend(other *ArrayList) {
	a.elements = append(a.elements, other.elements...)
}

func main() {

	list := &ArrayList{}

	// Append
	list.Append('a')
	list.Append('b')
	fmt.Println("After Append('a', 'b'):", list.elements)

	// Length
	fmt.Println("Length:", list.Length())

	// Insert
	list.Insert('c', 1)
	fmt.Println("After Insert('c', 1):", list.elements)
	// Get
	fmt.Println("Get(1):", string(list.Get(1)))

	// Delete
	deleted := list.Delete(1)
	fmt.Println("Deleted element at index 1:", string(deleted))
	fmt.Println("After Delete(1):", list.elements)

	// DeleteAll
	list.Append('a')
	list.Append('a')
	fmt.Println("Before DeleteAll('a'):", list.elements)
	list.DeleteAll('a')
	fmt.Println("After DeleteAll('a'):", list.elements)

	// Clone
	cloned := list.Clone()
	cloned.Append('d')
	list.Append('d')
	fmt.Println("Original list:", list.elements)
	fmt.Println("Cloned list:", cloned.elements)

	// Reverse
	list.Reverse()
	fmt.Println("After Reverse:", list.elements)

	// FindFirst и FindLast
	list.Append('x')
	list.Append('y')
	list.Append('x')
	fmt.Println("List:", list.elements)
	fmt.Println("FindFirst('x'):", list.FindFirst('x'))
	fmt.Println("FindLast('x'):", list.FindLast('x'))

	// Extend
	otherList := &ArrayList{}
	otherList.Append('z')
	otherList.Append('w')
	list.Extend(otherList)
	fmt.Println("After Extend:", list.elements)

	// Clear
	list.Clear()
	fmt.Println("After Clear:", list.elements)

	// Panic handler when wrong index
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	list.Append('a')
	list.Insert('b', 5)
}
