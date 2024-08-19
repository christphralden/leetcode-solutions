// Time : 2ms (74.41%) | Memory: 2.65 (97.51)
package main

import (
	"fmt"
)


type ListNode struct {
	Val  int
	Next *ListNode
}

func createNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func generator(head *ListNode) *ListNode {
	data := []int{1, 2, 3, 4, 5}
	current := head

	for _, x := range data {
		newNode := createNode(x)
		current.Next = newNode
		current = newNode
	}
	return head.Next
}

func reverseList(head *ListNode) *ListNode {
  
  var fast *ListNode = head
  var slow *ListNode = nil
  var temp *ListNode = nil

  for fast != nil {
    temp = fast.Next
    fast.Next = slow
    slow = fast
    fast = temp

  }
  
  return slow
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func main() {
	// Initialize a dummy head node
	dummyHead := &ListNode{}
	
	// Generate the list starting from the dummy head
	head := generator(dummyHead)
	
	fmt.Println("Original list:")
	printList(head)
	
	// Reverse the list
	head = reverseList(head)
	
	fmt.Println("Reversed list:")
	printList(head)
}
