package main

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
		l.head.next = l.head
		l.head = newNode
		l.length++
	}else {
		l.head = newNode

	}

}