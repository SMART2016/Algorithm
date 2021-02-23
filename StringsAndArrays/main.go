package main
import (
  "fmt"
  "sort"
  )


func main(){
   fmt.Println("Arrays And String")
   fmt.Println(HasUniqueCharsV2("iamforu"))
   fmt.Println(isAnagramV2("dog","god123"))
}

//It is assumed that the string contains only small letters between a to z (26 Letters)
func HasUniqueCharsV2(s string)(bool){
   var checker int32 =  0

   for _,v := range s{
      index := v - 'a'

      if checker & (1 << index ) > 0 {
        return false
      }
      checker = (checker | (1 << index))
   }
  
  return true
}



func isAnagramV2 (s1 string,s2 string) bool{
  if len(s1) != len(s2){
    return false
  }
  sortedS1 := SortStringByCharacterV2(s1)
  sortedS2 := SortStringByCharacterV2(s2)

  if sortedS1 == sortedS2{
    return true
  }else{
    return false
  }

}

func SortStringByCharacterV2(s string) string {
      r := []rune(s)
      sort.Slice(r, func(i, j int) bool {
              return r[i] < r[j]
      })
      return string(r)
}