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


}