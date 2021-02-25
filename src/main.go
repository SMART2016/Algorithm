package main

import (
	"fmt"
	lst "linklist"
	"math/rand"
	st "stringsandarrays"
)

func main() {
	fmt.Println("Arrays And String")
	fmt.Printf("\nString has All Unique characters: %t \n", st.HasUniqueChars("iamforu"))

	fmt.Printf("Strings are anagram: %t \n", st.IsAnagram("dog", "god123"))

	s1 := "Hi There How are You"
	r := make([]rune, 29)
	fmt.Println("Space encoded String: " + st.EncodeSpaceString(s1, 20, r))

	a := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	st.MarkWithZero(a, 3, 4)

	st.IsStrRoatated("erbottlewat", "waterbottle")

	list := createList()
	list.Display()
}

func createList() *lst.List {
	items := &lst.List{}
	size := 10
	//rand_number := make([]int, size, size)
	for i := 0; i < size; i++ {
		data := rand.Intn(100)
		if data == 0 {
			i = i - 1
			continue

		}
		items.InsertNode(data)
	}
	return items
}
