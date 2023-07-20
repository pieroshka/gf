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
	BRACKET_OPEN    = '['
	BRACKET_CLOSE   = ']'
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
		BRACKET_OPEN,
		BRACKET_CLOSE,
	}
)

func isToken(b byte) bool {
	for i := range LEGAL_CHARS {
		if b == LEGAL_CHARS[i] {
			return true
		}
	}
	return false
}

func lexical(bfcode []byte) []byte {
	var out []byte
	for i := 0; i < len(bfcode); i++ {
		if isToken(bfcode[i]) {
			out = append(out, bfcode[i])
		}
	}

	return out
}

func syntactical(bfcode []byte) (int, bool) {
	var stack []rune

	for i := range bfcode {
		if bfcode[i] == BRACKET_OPEN {
			stack = append(stack, BRACKET_OPEN)
		}

		if bfcode[i] == BRACKET_CLOSE {
			if len(stack) < 1 || stack[len(stack)-1] != BRACKET_OPEN {
				return i, false
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return len(bfcode) - 1, false
	}

	return -1, true
}

func countMemoryNeededNaive(bfcode []byte) uint {
	var min, max int
	for i := range bfcode {
		switch bfcode[i] {
		case PTR_MOVE_LEFT:
			min--
		case PTR_MOVE_RIGHT:
			max++
		default:
			continue
		}
	}
	fmt.Println("min:", min, "max:", max)
	return uint(max - min)
}

func main() {
	bfcode, err := os.ReadFile("demo/main.bf")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(lexical(bfcode)))
	fmt.Println(syntactical(bfcode))
	fmt.Println(countMemoryNeededNaive(bfcode))
}
