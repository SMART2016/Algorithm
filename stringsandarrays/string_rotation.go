package main
import (
"fmt"
"strings"
) 

func isStrRoatated( s1 string, s2 string){
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