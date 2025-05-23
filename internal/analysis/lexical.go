package analysis

import (
	"slices"

	"github.com/pieroshka/gf/internal/token"
)

type Lexical struct{}

func NewLexical() *Lexical {
	return &Lexical{}
}

func isToken(b byte) bool {
	return slices.Contains(token.Tokens, b)
}
func (l Lexical) Run(source []byte) []byte {
	var out []byte
	for i := range source {
		if isToken(source[i]) {
			out = append(out, source[i])
		}
	}
	return out
}
