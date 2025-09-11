package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
)

var replaceMap = make(map[rune]rune)

func main() {
	d, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(bytes.NewReader(d))
	r.Comma = ','

	for {
		record, err := r.Read()
		if err != nil {
			break
		}
		if len(record) != 2 {
			panic("invalid record")
		}
		kk := []rune(record[0])
		vv := []rune(record[1])
		for i, k := range kk {
			fmt.Println(string(k))
			replaceMap[k] = vv[i]
		}
	}

	f, err := os.Create("replacemap.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("package xstrings\n\n")
	f.WriteString("var nameRuneReplaceMap = map[rune]rune{\n")
	for k, v := range replaceMap {
		f.WriteString(fmt.Sprintf("'%c': '%c',\n", k, v))
	}
	f.WriteString("}\n")
}
