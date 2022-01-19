package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func usage() {
	log.Printf("Usage: sort [-k column] [-n sortByNumber] [-r reverse] [-u unique]")
	flag.PrintDefaults()
}

func main() {
	var column = flag.Int("k", 0, "Колонки для сортировки")
	//var sortByNumber = flag.Bool("n", false, "Сортировать по числовому значению")
	var reverse = flag.Bool("r", false, "Сортировать в обратном порядке")
	var unique = flag.Bool("u", false, "Не выводить повторяющиеся строки")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	slice := make([]string, 0)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		slice = append(slice, sc.Text())
	}

	if *unique {
		slice = removeDuplicates(slice)
	}

	sortData(slice, *column, *reverse)
	fmt.Printf("%+q", slice)
}

func removeDuplicates(s []string) []string {
	newSlice := make([]string, 0)
	hash := make(map[string]struct{})
	for _, v := range s {
		if _, ok := hash[v]; !ok {
			newSlice = append(newSlice, v)
			hash[v] = struct{}{}
		}
	}
	return newSlice
}

func sortData(slice []string, column int, reverse bool) {
	sort.Slice(slice, func(i, j int) bool {
		res := func() bool {
			first := strings.Fields(slice[i])
			second := strings.Fields(slice[j])
			iHaveKey := len(first) > column
			jHaveKey := len(second) > column

			if !iHaveKey && jHaveKey {
				return true
			}
			if iHaveKey && !jHaveKey {
				return false
			}

			if iHaveKey && jHaveKey {
				if first[column] < second[column] {
					return true
				}
				if first[column] > second[column] {
					return false
				}
			}
			return len(first) < len(second)
		}()
		return res != reverse
	})
}
