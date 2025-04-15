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

// DeleteAll delete all oqqurencies of element
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

}
