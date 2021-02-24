package main


//Encodes spaces in a string with %20 , without using any extra space
func encodeSpaceString(s string, length int,r1 []rune)string{
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