package stringsandarrays

//It is assumed that the string contains only small letters between a to z (26 Letters)
func HasUniqueChars(s string)(bool){
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