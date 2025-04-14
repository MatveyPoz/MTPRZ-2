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
}
