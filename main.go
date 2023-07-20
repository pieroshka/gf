package main

import (
	"fmt"
	"os"
)

var LEGAL_CHARS = []byte{'>', '<', '+', '-', '.', ',', '[', ']'}

const MEMORY_SIZE = 300000

func isLegal(b byte) bool {
	for c := range LEGAL_CHARS {
		if b == LEGAL_CHARS[c] {
			return true
		}
	}
	return false
}

func remove(slice []byte, s int) []byte {
	return append(slice[:s], slice[s+1:]...)
}

func preprocess(bfcode []byte) []byte {
	for c := 0; c < len(bfcode); c++ {
		if !isLegal(bfcode[c]) {
			bfcode = remove(bfcode, c)
			c--
		}
	}

	return bfcode
}

func main() {
	bfcode, err := os.ReadFile("./main.bf")
	if err != nil {
		panic(err)
	}

	bfcode = preprocess(bfcode)

	fmt.Println(string(bfcode))
}
