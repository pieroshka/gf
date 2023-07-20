package main

import (
	"fmt"
	"gofuck/memory"
	"log"
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
	MEMORY_SIZE uint = 4096
	TOKENS           = []byte{
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
	for i := range TOKENS {
		if b == TOKENS[i] {
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

	return len(bfcode), len(stack) == 0
}

func getCorrespondingClosingBracketIdx(idx uint64, bfcode []byte) uint64 {
	if idx == uint64(len(bfcode)-1) {
		return 0 // should never happen
	}
	counter := 1
	for i := idx + 1; i < uint64(len(bfcode)); i++ {
		if bfcode[i] == BRACKET_OPEN {
			counter++
		}

		if bfcode[i] == BRACKET_CLOSE {
			counter--
			if counter == 0 {
				return i
			}
		}
	}
	return 404 // should never occur for a file that passed the syntactical analysis
}

func getCorrespondingOpeningBracketIdx(idx uint64, bfcode []byte) uint64 {
	if idx == 0 {
		return 0 // should never happen
	}

	counter := 1
	for i := idx - 1; i > 0; i-- {
		if bfcode[i] == BRACKET_CLOSE {
			counter++
		}

		if bfcode[i] == BRACKET_OPEN {
			counter--
			if counter == 0 {
				return i
			}
		}
	}
	return 0 // should never occur for a file that passed the syntactical analysis
}

func interpret(bfcode []byte) {
	for i := uint64(0); i < uint64(len(bfcode)); i++ {
		switch c := bfcode[i]; c {
		case PTR_MOVE_LEFT:
			memory.PointerMoveLeft()

		case PTR_MOVE_RIGHT:
			memory.PointerMoveRight()

		case INCR_MEM_CELL:
			memory.Incr()

		case DECR_MEM_CELL:
			memory.Decr()

		case OUTPUT_MEM_CELL:
			fmt.Printf("%c", memory.Get())

		case INPUT_MEM_CELL:
			var temp byte
			fmt.Scanf("%c", &temp)
			memory.Set(temp)

		case BRACKET_OPEN:
			if memory.Get() == 0 {
				i = getCorrespondingClosingBracketIdx(i, bfcode)
			}

		case BRACKET_CLOSE:
			if memory.Get() != 0 {
				i = getCorrespondingOpeningBracketIdx(i, bfcode) - 1
			}
		}
	}
}

func main() {
	bfcode, err := os.ReadFile("./input.bf")
	if err != nil {
		panic(err)
	}

	bfcode = lexical(bfcode)
	i, ok := syntactical(bfcode)
	if !ok {
		log.Fatalln("syntax error at position:", i)
	}

	interpret(bfcode)
}
