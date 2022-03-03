package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	var curr *LinkedList = linkedList

	for curr != nil && curr.Next != nil {
		if curr.Value == curr.Next.Value {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}

	return linkedList
}
