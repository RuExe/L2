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
func UnpackString(s string) string {
	str := []rune(s)
	if len(str) == 0 || unicode.IsDigit(str[0]) {
		return ""
	}

	var b strings.Builder
	for i := 0; i < len(str); {
		char := str[i]
		i++
		countStr := ""
		for ; i < len(str) && unicode.IsDigit(str[i]); i++ {
			countStr += string(str[i])
		}

		count, err := strconv.Atoi(countStr)
		if err != nil {
			count = 1
		}
		fmt.Fprint(&b, strings.Repeat(string(char), count))
	}
	return b.String()
}
