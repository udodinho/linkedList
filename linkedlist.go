package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkedList struct {
	head *node
	length int
}

func (l *linkedList) prepend(val int)  {
	newNode := &node{data: val}
	if l.head != nil {
		newNode.next = l.head
		l.head = newNode
		l.length++
	}else {
		l.head = newNode
		l.length++
	}
	return
}

func (l *linkedList) print()  {
	if l.head == nil {
		panic("Empty List")
	}

	current := l.head
	for current != nil {
	fmt.Println(current.data)
	current = current.next
	}

}

func (l *linkedList) deleteWithValue(val int)  {
	if l.length == 0 {
		return
	}
	if l.head.data == val {
		l.head = l.head.next
		l.length--
	}

	previousDel := l.head
	for previousDel.next.data != val {
		if previousDel.next.next == nil {

		}
	}

}



func main() {
	v := &linkedList{}
	v.prepend(33)
	v.prepend(80)
	v.prepend(90)

	v.print()
}