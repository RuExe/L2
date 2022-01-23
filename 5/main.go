package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	var after = flag.Int("A", 0, "печатать +N строк после совпадения")
	var before = flag.Int("B", 0, "печатать +N строк до совпадения")
	var context = flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	var count = flag.Bool("c", false, "количество строк")
	var ignore = flag.Bool("i", false, "игнорировать регистр")
	var invert = flag.Bool("v", false, "вместо совпадения, исключать")
	var fixed = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	var linenum = flag.Bool("n", false, "напечатать номер строки")

	log.SetFlags(0)
	flag.Parse()

	args := flag.Args()

	pt := args[0]
	if *fixed {
		pt = `\Q` + pt + `\E`
	}

	if *ignore {
		pt = "(?i)" + pt
	}
	r, _ := regexp.Compile(pt)

	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	strings := bytes.Split(content, []byte("\n"))
	temp := findLines(strings, r, *invert)
	if *count {
		fmt.Print(len(temp))
	} else {
		printRes(strings, temp, *linenum, *after, *before, *context)
	}
}

func printRes(lines [][]byte, res []int, linenum bool, after, before, context int) {
	b := maxInt(before, context)
	a := maxInt(after, context)
	fmt.Println(b)
	fmt.Println(a)
	for _, v := range res {
		be := maxInt(v-b, 0)
		af := minInt(1+v+a, len(lines))
		for i := be; i < af; i++ {
			if linenum {
				fmt.Printf("Результат %v: %s\n", v, lines[i])
			} else {
				fmt.Printf("%q\n", lines[i])
			}
		}
	}
}

func findLines(lines [][]byte, r *regexp.Regexp, invert bool) []int {
	res := make([]int, 0)
	for i := range lines {
		if r.Match(lines[i]) != invert {
			res = append(res, i)
		}
	}
	return res
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
