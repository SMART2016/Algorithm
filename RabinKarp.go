package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	prime int = 101
)

//Input:  txt[] = "THIS IS A TEST TEXT"
//        pat[] = "TEST"
//Output: Pattern found at index 10

//Input:  txt[] =  "AABAACAADAABAABA"
//        pat[] =  "AABA"
//Output: Pattern found at index 0
//        Pattern found at index 9
//        Pattern found at index 12
func compare(pattern string, text string) {
	m := len([]rune(pattern))
	n := len([]rune(text))
	indexMatched := make([]int, 0)
	if m <= 0 || n <= 0 || m > n {
		fmt.Println("Invalid Pattern , Pattern must be smaller than or equal to the actual Text")
		return
	}
	fmt.Println("Text is: ", text)
	fmt.Println("Pattern is:", pattern)
	hashPattern := hash(pattern)
	runeText := []rune(text)

	hashInitialText := hash(string(runeText[0:3]))
	for i := 0; i < n-m+1; i++ {
		if hashInitialText == hashPattern {
			result := strings.Compare(pattern, string(runeText[i:i+3]))
			if result == 0 {
				indexMatched = append(indexMatched, i)
				continue
			}
		}

		if i != n-m {
			hashInitialText = recalculateHash(hashInitialText, int(runeText[i]), int(runeText[i+3]), len(pattern))
		}
	}
	fmt.Println(indexMatched)
}

//The recalculate hash for rolling string:
//(oldHash for the previous string - the first char of the old string set removed)/prime + (new character added * (prime ^ length of pattern -1 ))
func recalculateHash(oldHash, leadingCharAscii, trailingCharAscii, patternLength int) int {
	return int(math.Ceil((float64(oldHash-leadingCharAscii) / float64(prime)))) + trailingCharAscii*int(math.Pow(float64(prime), float64(patternLength-1)))
}

//Returns the hash of the provided string based on below formulae:
// hash = sum (ith Char of String * prime ^ i)  where i = len(str)
func hash(str string) int {
	var hashPattern int = 0
	for i, char := range []rune(str) {
		hashPattern += int(char) * int(math.Pow(float64(prime), float64(i)))
	}
	return hashPattern
}
