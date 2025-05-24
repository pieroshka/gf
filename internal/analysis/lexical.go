package analysis

import (
	"slices"

	"github.com/pieroshka/gf/internal/token"
)

type lexical struct{}

func NewLexical() *lexical {
	return &lexical{}
}

func isToken(b byte) bool {
	return slices.Contains(token.Tokens, b)
}

func (l lexical) Run(source []byte) []byte {
	out := []byte{}
	for i := range source {
		if isToken(source[i]) {
			out = append(out, source[i])
		}
	}
	return out
}
