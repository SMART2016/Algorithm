package linklist

import "fmt"

type List struct {
	Start  *Node
	length int
}

func (l *List) Display() {
	fmt.Print("Linklist is :")
	node := l.Start
	for node != nil {
		fmt.Printf("%v ->", node.data)
		node = node.next
	}
	fmt.Println()
}

//Allows insert of duplicate elements
func (l *List) InsertNode(value int) {
	newNode := &Node{
		data: value,
	}
	if l.length == 0 {
		l.Start = newNode
	} else {
		currentNode := l.Start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}

	l.length++
}

//Removes all occurences of the elemnt from the link list
func (l *List) RemoveNode(node *Node ,value int){
  if l.length == 0 || node == nil{
    return
  }
  if node.data == value{
      tmp := node
      node = node.next
      if l.Start == tmp{
        l.Start = node
      }
      tmp = nil
      l.length --
   }else{
      current := node.next
      if current != nil{
        if current.data == value{
           node.next = current.next
           node = current.next
           current = nil
           l.length--
        }else{
          l.RemoveNode(current,value)
        }
      }
      node = current
   }
   l.RemoveNode(node,value)
}
