package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	str := "a4bc2d5e"
	fmt.Println(UnpackString(str))

	str = "abcd"
	fmt.Println(UnpackString(str))

	str = "45"
	fmt.Println(UnpackString(str))

	str = ""
	fmt.Println(UnpackString(str))

	str = "a10bc3d15"
	fmt.Println(UnpackString(str))
}

// UnpackString unpack string
func UnpackString(str string) string {
	if len(str) == 0 || unicode.IsDigit(rune(str[0])) {
		return ""
	}

	var b strings.Builder
	for i := 0; i < len(str); {
		res := string(str[i])
		i++
		countStr := ""
		for ; i < len(str) && unicode.IsDigit(rune(str[i])); i++ {
			countStr += string(str[i])
		}

		count, err := strconv.Atoi(countStr)
		if err != nil {
			count = 1
		}
		fmt.Fprint(&b, strings.Repeat(res, count))
	}
	return b.String()
}
