// Lab 12 -- slices and functions
//
// Programmer name: Nigel S Adams
// Date completed:  11/12/19

package main

import (
	"strings"
)

// StringsOfLength returns a slice containing all strings
// in a given list matching a given length.
func StringsOfLength(list []string, length int) []string {
	var listStrg []string
	for _, strg := range list {
		if len(strg) == length {
			listStrg = append(listStrg, strg)
		}

	}
	return listStrg
}

// StringsLongerThan returns a slice containing all strings
// in a given list longer than a given length.
func StringsLongerThan(list []string, length int) []string {
	var listStrg []string
	for _, strg := range list {
		if len(strg) > length {
			listStrg = append(listStrg, strg)
		}
	}
	return listStrg
}

// StringsShorterThan returns a slice containing all strings
// in a given list shorter than a given length.
func StringsShorterThan(list []string, length int) []string {
	var listStrg []string
	for _, strg := range list {
		if len(strg) < length {
			listStrg = append(listStrg, strg)
		}
	}
	return listStrg
}

// StringsBeginningWith returns a slice containing all strings
// in a given list that begin with a given prefix. Case insensitive.
func StringsBeginningWith(list []string, prefix string) []string {
	var listStrg []string
	for _, strg := range list {
		if strings.HasPrefix(strg, prefix) == true {
			listStrg = append(listStrg, strg)
		}
	}
	return listStrg
}

// ContainsString returns true if a given slice contains a target
// string. Case insensitive.
func ContainsString(list []string, target string) bool {
	var obj bool
	for _, strg := range list {
		if strings.Contains(strg, target) == true {
			obj = true
		}
	}

	return obj
}

// GetUnique returns a slice containing only non-duplicate strings
// from a given list. (HINT: use ContainsString as a helper)
func GetUnique(list []string) []string {
	var special []string

	for _, word := range list {
		if !ContainsString(special, word) {
			special = append(special, word)
		}

	}

	return special
}

// LengthOfShortest returns the length of the shortest string in a
// given list of strings.
func LengthOfShortest(list []string) int {
	var short int
	short = len(list[0])

	for _, obj := range list {
		if len(obj) < short {
			short = len(obj)
		}
	}

	return short
}

// LengthOfLongest returns the length of the longest string in a
// given list of strings.
func LengthOfLongest(list []string) int {
	var long int
	long = len(list[0])

	for _, obj := range list {
		if len(obj) > long {
			long = len(obj)
		}
	}

	return long

}

// ListOfShortest returns a slice containing all strings in a given
// list whose length matches the shortest string in the list.
//  (HINT: use LengthOfShortest as a helper)
func ListOfShortest(list []string) []string {
	var shortest int = LengthOfShortest(list)
	var ShortList []string
	for _, obj := range list {
		if len(obj) <= shortest {
			ShortList = append(ShortList, obj)
		}
	}
	return ShortList
}

// ListOfLongest returns a slice containing all strings in a given
// list whose length matches the longest string in the list.
//  (HINT: use LengthOfLongest as a helper)
func ListOfLongest(list []string) []string {
	var longest int = LengthOfLongest(list)
	var LongList []string
	for _, obj := range list {
		if len(obj) >= longest {
			LongList = append(LongList, obj)
		}
	}
	return LongList
}

func main() {
}
