package main

import (
	"fmt"
)

type Node struct { // New Node
	value int
	next  *Node
}
type Stack struct { // New stack
	head *Node
	len  int
}

func (s Stack) Display() { // Display function
	if s.head == nil { // if head is nil
		fmt.Println("No stack is found")
	}
	temp := s.head    // temp variable for looping
	for temp != nil { // Traversing to first to last
		fmt.Print(" <-- ", temp.value) // Printing Value
		temp = temp.next
	}
}

func (s *Stack) Pop() {
	if s.head == nil {
		fmt.Println("Stack underflow")
		return
	}
	s.head = s.head.next

}
func (s *Stack) push(val int) { // Push opration value
	newNode := &Node{value: val} // Create a New Node using by Given value
	if s.head == nil {           // if head is nill
		s.head = newNode // Set as head
	} else {
		newNode.next = s.head // Else Head is there
		s.head = newNode      // Value is Adding
	}
}
func main() {
	new := &Stack{}
	new.push(1)
	new.push(2)
	new.push(3)
	new.push(4)
	new.Display()
	fmt.Println("\n-------------------")
	new.Pop()
	new.Display()
}
