package stringsandarrays

import "sort"

func IsAnagram (s1 string,s2 string) bool{
  if len(s1) != len(s2){
    return false
  }
  sortedS1 := SortStringByCharacter(s1)
  sortedS2 := SortStringByCharacter(s2)

  if sortedS1 == sortedS2{
    return true
  }else{
    return false
  }

}

func SortStringByCharacter(s string) string {
      r := []rune(s)
      sort.Slice(r, func(i, j int) bool {
              return r[i] < r[j]
      })
      return string(r)
}