package main

import (
	"fmt"
)

// Node tructure
type Node struct {
	value rune
	next  *Node
}

// ArrayList structure which realizes list of Node
type ArrayList struct {
	tail   *Node
	length int
}

// Length return amount of elements
func (a *ArrayList) Length() int {
	return a.length
}

// Append adds element to the end of list
func (a *ArrayList) Append(element rune) {
	newNode := &Node{value: element}

	if a.tail == nil {
		newNode.next = newNode
		a.tail = newNode
	} else {
		newNode.next = a.tail.next
		a.tail.next = newNode
		a.tail = newNode
	}
	a.length++
}

// Insert insert element to speific position
func (a *ArrayList) Insert(element rune, index int) {
	if index < 0 || index > a.length {
		panic("invalid index for insert")
	}

	if index == a.length {
		a.Append(element)
		return
	}

	newNode := &Node{value: element}
	if index == 0 {
		newNode.next = a.tail.next
		a.tail.next = newNode
	} else {
		prev := a.getNode(index - 1)
		newNode.next = prev.next
		prev.next = newNode
	}
	a.length++
}

// getNode return *Node by index
func (a *ArrayList) getNode(index int) *Node {
	if index < 0 || index >= a.length {
		panic("invalid index")
	}

	current := a.tail.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current
}

// Delete delete node by index and return it
func (a *ArrayList) Delete(index int) rune {
	if index < 0 || index >= a.length {
		panic("invalid index for delete")
	}

	var value rune
	if a.length == 1 {
		value = a.tail.value
		a.tail = nil
	} else {
		var prev *Node
		current := a.tail.next

		if index == 0 {
			prev = a.tail
		} else {
			prev = current
			for i := 0; i < index-1; i++ {
				prev = prev.next
			}
		}

		if prev.next == a.tail {
			a.tail = prev
		}

		value = prev.next.value
		prev.next = prev.next.next
	}
	a.length--
	return value
}

// DeleteAll delete all occurrences of *Node.value
func (a *ArrayList) DeleteAll(element rune) {
	if a.length == 0 {
		return
	}

	current := a.tail.next
	prev := a.tail
	count := 0

	for i := 0; i < a.length; i++ {
		next := current.next
		if current.value == element {
			prev.next = next
			if current == a.tail {
				a.tail = prev
			}
			count++
		} else {
			prev = current
		}
		current = next
	}
	a.length -= count
}

// Get return *Node.value by index
func (a *ArrayList) Get(index int) rune {

	return a.getNode(index).value
}

func (a *ArrayList) Clone() *ArrayList {
	clone := &ArrayList{}
	if a.length == 0 {
		return clone
	}

	current := a.tail.next
	for i := 0; i < a.length; i++ {
		clone.Append(current.value)
		current = current.next
	}
	return clone
}

// Reverse reverse order of *Node in list
func (a *ArrayList) Reverse() {
	if a.length < 2 {
		return
	}

	var prev, current, next *Node
	current = a.tail.next
	tailNode := current

	for i := 0; i < a.length; i++ {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	a.tail = tailNode
	a.tail.next = prev
}

// FindFirst return first *Node.vsalue occurence in list index or -1
func (a *ArrayList) FindFirst(element rune) int {
	current := a.tail.next
	for i := 0; i < a.length; i++ {
		if current.value == element {
			return i
		}
		current = current.next
	}
	return -1
}

// FindLast return last element occurrence in list index or -1
func (a *ArrayList) FindLast(element rune) int {
	lastIndex := -1
	current := a.tail.next
	for i := 0; i < a.length; i++ {
		if current.value == element {
			lastIndex = i
		}
		current = current.next
	}
	return lastIndex
}

// Clear clean linked list
func (a *ArrayList) Clear() {
	a.tail = nil
	a.length = 0
}

// Extend add nodes from another linked list
func (a *ArrayList) Extend(elements *ArrayList) {
	if elements.length == 0 {
		return
	}

	current := elements.tail.next
	for i := 0; i < elements.length; i++ {
		a.Append(current.value)
		current = current.next
	}
}

func (a *ArrayList) GetList() string {
	var result string
	for i := 0; i < a.length; i++ {
		result += string(a.getNode(i).value)
	}
	return result
}

func main() {

	list := &ArrayList{}

	// Append
	list.Append('a')
	list.Append('b')
	fmt.Println("After Append('a', 'b'):", list.GetList())

	// Length
	fmt.Println("Length:", list.Length())

	// Insert
	list.Insert('c', 1)
	fmt.Println("After Insert('c', 1):", list.GetList())
	// Get
	fmt.Println("Get(1):", string(list.Get(1)))

	// Delete
	deleted := list.Delete(1)
	fmt.Println("Deleted element at index 1:", string(deleted))
	fmt.Println("After Delete(1):", list.GetList())

	// DeleteAll
	list.Append('a')
	list.Append('a')
	fmt.Println("Before DeleteAll('a'):", list.GetList())
	list.DeleteAll('a')
	fmt.Println("After DeleteAll('a'):", list.GetList())

	// Clone
	cloned := list.Clone()
	cloned.Append('d')
	list.Append('d')
	fmt.Println("Original list:", list.GetList())
	fmt.Println("Cloned list:", cloned.GetList())

	// Reverse
	list.Reverse()
	fmt.Println("After Reverse:", list.GetList())

	// FindFirst и FindLast
	list.Append('x')
	list.Append('y')
	list.Append('x')
	fmt.Println("List:", list.GetList())
	fmt.Println("FindFirst('x'):", list.FindFirst('x'))
	fmt.Println("FindLast('x'):", list.FindLast('x'))

	// Extend
	otherList := &ArrayList{}
	otherList.Append('z')
	otherList.Append('w')
	list.Extend(otherList)
	fmt.Println("After Extend:", list.GetList())

	// Clear
	list.Clear()
	fmt.Println("After Clear:", list.GetList())

	// Panic handler when wrong index
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	list.Append('a')
	list.Insert('b', 5)
}
