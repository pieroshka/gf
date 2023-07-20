package main

import (
	"fmt"
	"os"
)

const (
	PTR_MOVE_RIGHT  = '>'
	PTR_MOVE_LEFT   = '<'
	INCR_MEM_CELL   = '+'
	DECR_MEM_CELL   = '-'
	OUTPUT_MEM_CELL = '.'
	INPUT_MEM_CELL  = ','
	LOOP_OPEN       = '['
	LOOP_CLOSE      = ']'
)

var (
	MEMORY_SIZE uint64
	LEGAL_CHARS = []byte{
		PTR_MOVE_RIGHT,
		PTR_MOVE_LEFT,
		INCR_MEM_CELL,
		DECR_MEM_CELL,
		OUTPUT_MEM_CELL,
		INPUT_MEM_CELL,
		LOOP_OPEN,
		LOOP_CLOSE,
	}
)

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

func lexical(bfcode []byte) []byte {
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

	fmt.Println(string(lexical(bfcode)))
}
