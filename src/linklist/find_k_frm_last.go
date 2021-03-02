package linklist
import (
  "fmt"
)

/**
  Approach 1: 
      If we Know the length of the Link List we can identify that Kth Element from the end is (Length - K)th element from the start of the Link List.

  Approach 2:
     Runner Technique:  Fast and Slow Pointer
     - Start a fast pointer and loop the fast pointer till it reahes the end of the list.
     - Increment a temp variable to count when the fast pointer reaches K.
     - When fast pointer reaches K , initialize the slow pointer.
     - run the slow pointer through the list while looping the fast pointer till the end of the list.
     - Print the element pointed to by the slow pointer when the fast pointer reaches the end and that will be the Kth element from the end of the list.

**/
func (lst *List) FindKLast(k int){
   fast := lst.Start
   var slow *Node
   k_running := 0
   for fast != nil{
     if k_running == k{
         slow = lst.Start
       }
      if ( slow != nil){
        slow = slow.next
      }
     fast = fast.next
     k_running ++
   }
   fmt.Printf("Kth = %d, Elemnt from end of List is : %d\n",k,slow.data)
}