package main

import "fmt"

//Efficient String concatenation
func concat(length int,s1 string,s2 string){
   fmt.Println("Hello concat")
   size := len(s2) + len(s2)
   cArr := make([]rune,size)
   
   for i,v := range s1{
     cArr[i] = v
   }
   
  for i,v := range s2{
     cArr[len(s2) + i] = v
  }

  fmt.Println("Concatenated String is :"+ string(cArr))
}