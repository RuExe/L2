package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	slice := []string{
		"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слиток",
	}
	sort.Strings(slice)

	res := Anagrams(&slice)
	fmt.Println(slice)
	fmt.Println(*res)
}

//func Anagrams(slice []string) *map[string][]string {
//	table := map[string][]string{}
//	for i := 0; i < len(slice); i++ {
//		for j := i + 1; j < len(slice); j++ {
//			iWord := strings.ToLower(slice[i])
//			jWord := strings.ToLower(slice[j])
//
//			iSorted := SortStringByCharacter(iWord)
//			jSorted := SortStringByCharacter(jWord)
//			if iSorted == jSorted {
//				table[iSorted] = append(table[iSorted], jWord)
//			}
//		}
//	}
//
//	for k := range table {
//		if len(table[k]) < 2 {
//			delete(table, k)
//		}
//	}
//	return &table
//}

func Anagrams(slice *[]string) *map[string][]string {
	keys := make(map[string]struct{})
	table := make(map[string][]string)
	for _, v := range *slice {
		if _, ok := keys[v]; ok {
			continue
		}
		keys[v] = struct{}{}

		word := strings.ToLower(v)
		sorted := SortStringByCharacter(word)
		table[sorted] = append(table[sorted], word)
	}

	res := make(map[string][]string)
	for k, v := range table {
		if len(table[k]) > 2 {
			res[v[0]] = v[1:]
		}
	}
	return &res
}

func SortStringByCharacter(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
