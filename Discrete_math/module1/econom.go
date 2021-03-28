package main

import (
	"bytes"
	"fmt"
)

func main() {
	var (
		polandLetter []byte
		result int
		closeIndex int
	)
	fmt.Scan(&polandLetter)
	closeIndex = bytes.IndexAny(polandLetter, ")")
	for closeIndex > 0 {
		closeIndex -= 4
		result++
		polandLetter = bytes.Replace(polandLetter, polandLetter[closeIndex: closeIndex + 5], ([]byte)(string(0)), -1)
		closeIndex = bytes.IndexAny(polandLetter, ")")
	}
	fmt.Println(result)
}
