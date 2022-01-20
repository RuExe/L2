package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	slice := []string{
		"пятак", "листок", "слиток", "столик", "слиток", "тяпка", "пятка",
	}

	res := Anagrams(&slice)
	fmt.Println(slice)
	fmt.Println(*res)
}

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
			temp := v[1:]
			sort.Strings(temp)
			res[v[0]] = temp
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
