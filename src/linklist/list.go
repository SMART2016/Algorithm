package linklist

import "fmt"

type List struct {
	start  *Node
	length int
}

func (l *List) Display() {
	fmt.Print("Linklist is :")
	node := l.start
	for node != nil {
		fmt.Printf("%v ->", node.data)
		node = node.next
	}
	fmt.Println()
}

func (l *List) InsertNode(value int) {
	newNode := &Node{
		data: value,
	}
	if l.length == 0 {
		l.start = newNode
	} else {
		currentNode := l.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}

	l.length++
}
