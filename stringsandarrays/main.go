package main
import (
  "fmt"
  "sort"
  "strings"
  )


func main(){
   fmt.Println("Arrays And String")
   fmt.Printf("\nString has All Unique characters: %t \n" ,HasUniqueCharsV2("iamforu"))
  
   fmt.Printf("Strings are anagram: %t \n",isAnagramV2("dog","god123"))

   s1 := "Hi There How are You"
   r := make([]rune,29)
   fmt.Println("Space encoded String: "+encodeSpaceStringV2(s1,20, r))

   a := [3][4]int{  
   {0, 1, 2, 3} ,  
   {4, 5, 6, 7} ,  
   {8, 9, 10, 11},  
   }
   markWithZero(a)

   isStrRoatatedV2("erbottlewat","waterbottle")
}

func isStrRoatatedV2( s1 string, s2 string){
   if len(s1) != len(s2){
     fmt.Printf("%s is not Rotation of %s",s2,s1)
     return
   }

   s := s1 + s1

    if strings.Contains(s,s2){
      fmt.Printf("%s is Rotation of %s",s2,s1)
    }else{
      fmt.Printf("%s is not Rotation of %s",s2,s1)
    }

}

func markWithZero(matrix [3][4]int){
   rows  := new([3]int)
   cols  := new([4]int)

   for  i := 0; i < 3; i++ {
      for j := 0; j < 4; j++ {
         if rows[i] != 1 && cols[j] != 1{
             fmt.Printf("Processing %d : %d\n",i,j)
             if matrix[i][j] == 0{
                setRowAndColsToZero(&matrix,i,j)
                rows[i] = 1
                cols[j] = 1
             }
         } 

      }
   }

   fmt.Println("Transformed Matrix : %v",matrix)
}

func setRowAndColsToZero(matrix *[3][4]int , row int ,col int){
  for i := 0; i < 4;i++{
      matrix[row][i] = 0
  }  

  for i:= 0;i< 3;i++{
     matrix[i][col] = 0
  }
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

//Encodes spaces in a string with %20 , without using any extra space
func encodeSpaceStringV2(s string, length int,r1 []rune)string{
  var spaceCount int
  for _,v := range s{
    if v == ' '{
        spaceCount++
    }
  }

  newLength := length + spaceCount * 2
  r := []rune(s)
  for i := length - 1; i >= 0; i-- {
     if r[i] == ' '{
        r1[newLength - 2]= '%'
        r1[newLength - 1]= '2'
        r1[newLength - 0]= '0'
        newLength = newLength - 3
     } else{
       r1[newLength] = r[i]
       newLength = newLength - 1
     }
  }

  return string(r1)
}

