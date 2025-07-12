package analysis

import (
	"errors"

	"github.com/pieroshka/gf/internal/token"
)

var (
	ErrUnexpectedClosingBracket = errors.New("unexpected closing bracket")
	ErrClosingBracketNotFound   = errors.New("closing bracket not found")
)

type syntactical struct{}

func NewSyntactical() *syntactical {
	return &syntactical{}
}

func (s syntactical) Run(source []byte) (int, bool, error) {
	var stack []byte
	var openingBracketIdx []int

	for i := range source {
		if source[i] == token.BracketOpen {
			openingBracketIdx = append(openingBracketIdx, i)
			stack = append(stack, token.BracketOpen)
		}

		if source[i] == token.BracketClose {
			if len(stack) < 1 || stack[len(stack)-1] != token.BracketOpen {
				return i, false, ErrUnexpectedClosingBracket
			}

			stack = stack[:len(stack)-1]
			openingBracketIdx = openingBracketIdx[:len(openingBracketIdx)-1]
		}
	}

	var errIdx int
	var err error
	if len(stack) != 0 {
		err = ErrClosingBracketNotFound
		errIdx = openingBracketIdx[len(openingBracketIdx)-1]
	}

	return errIdx, len(stack) == 0, err
}
