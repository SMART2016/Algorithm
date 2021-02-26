package linklist

/**
  Approach 1: 
      - Create a map
      - Iterate over the LL
          - for every element check if it is present in the map.
            - If it is present in the map delete it from the List 
                - Note: Keep a Prev and Current pointer to delete an element.
            - If the element is not present in the map then add the element to the map with value as true and continue to next element.
       - Problem is this algorithm needs an extra space O(n) and time complexity is O(n) , n is the number of elements.     

  Approach 2:
        - With two pointers:
            - one that will be pointing each element of the LL
            - other that will be going through each element to compare them with the element pointed by pointer one.
        - This is O(n^2) complexity algo, n is the number of elements and space complexity if O(1) 
**/


// Remove Duplicates from a singly LL without extra space
func (lst * List) RemoveDuplicates(){
  tmp1 := lst.Start

  for tmp1 != nil{
    tmp2 := tmp1.next
    prev := tmp1
    for tmp2 != nil{
      if tmp2.data == tmp1.data{
        prev.next = tmp2.next
      }else{
        prev = tmp2
      }
      tmp2 = tmp2.next
    }
    tmp1 = tmp1.next
  }
}