package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var fields = flag.Int("f", -1, "выбрать поля (колонки)")
	var delimiter = flag.String("d", "\t", "использовать другой разделитель")
	var separated = flag.Bool("s", false, "только строки с разделителем")

	log.SetFlags(0)
	flag.Parse()

	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	delimiterByte := []byte(*delimiter)
	lines := bytes.Split(content, []byte("\n"))

	if *separated {
		temp := make([][]byte, 0)
		for i := range lines {
			if bytes.Contains(lines[i], delimiterByte) {
				temp = append(temp, lines[i])
			}
		}
		lines = temp
	}

	res := make([][][]byte, 0)
	for i := range lines {
		temp := bytes.Split(lines[i], delimiterByte)
		if field := *fields; field > -1 {
			if len(temp) > field {
				temp = [][]byte{temp[field]}
			} else {
				temp = [][]byte{}
			}
		}
		res = append(res, temp)
	}

	fmt.Printf("%q", res)
}
