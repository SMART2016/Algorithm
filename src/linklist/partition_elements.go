package linklist
import "fmt"

/**
  Approach:
      - Navigate to the node around the partition needs to happen.
      - Isolate the Pivot node by creating another standalone node with the pivot data.
      - consider the first half of the pivot as one Linklist and the second half as another.
      - compare the elements of the two link list individually with the new pivot node and based on whether they are larger or smaller than pivot either append at the end (For larger than pivot) or prepend (for smaller than pivot)
**/
func (lst *List) Partition(k int) {
	if lst != nil {
		pivot := lst.Start
		var p2 *Node

		//Find the partition element in the Link list
		for pivot.data != k {
			pivot = pivot.next
		}

		//Once the partition element is found set a pointer to the next node of the partition and a pointer to the start of the List
		p2 = pivot.next
		p1 := lst.Start

		//Isolate pivot by disconnect the nodes before Pivot and after pivot

		pivotNode := &Node{
			data: pivot.data,
			next: nil,
		}

		s := pivotNode
		end := pivotNode

    fmt.Printf("Pivot is: %d \n" ,pivot.data)

		for p1 != pivot {
      tmp := p1
      p1 = p1.next
      fmt.Printf("Processing: %+v \n",tmp)
      if tmp.data < pivot.data{
        tmp.next = s
        s = tmp
      }else{
          end.next = tmp
          end = tmp
          
          end.next = nil
      }
     
		}
    
    for p2 != nil{
      tmp := p2
      p2 = p2.next
       if tmp.data < pivot.data{
        tmp.next = s
        s = tmp
      }else{
          end.next = tmp
          end = tmp
          end.next = nil
      }
     
    }

    t := s
    for (t != nil){
      fmt.Print("-->", t.data)
      t = t.next
    }
	}
}
