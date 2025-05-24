package analysis_test

import (
	"testing"

	"github.com/pieroshka/gf/internal/analysis"
	"github.com/stretchr/testify/assert"
)

func TestLexical(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		input          []byte
		expectedOutput []byte
	}{
		{"simple", []byte("abc>>>"), []byte(">>>")},
		{"only bf", []byte("++--"), []byte("++--")},
		{"empty", []byte(""), []byte("")},
		{"non-bf", []byte("xyz"), []byte("")},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := analysis.NewLexical()

			out := analyzer.Run([]byte(tt.input))
			assert.Equal(t, tt.expectedOutput, out)
		})
	}
}

func FuzzLexical(f *testing.F) {
	f.Add(">>>abc>>>")
	f.Add("+-hello++[world],.")
	f.Add("no bf chars")
	f.Add("")

	bfChars := map[rune]bool{
		'>': true, '<': true, '+': true, '-': true,
		'[': true, ']': true, '.': true, ',': true,
	}

	f.Fuzz(func(t *testing.T, input string) {
		analyzer := analysis.NewLexical()

		expected := make([]rune, 0, len(input))
		for _, ch := range input {
			if bfChars[ch] {
				expected = append(expected, ch)
			}
		}

		actual := analyzer.Run([]byte(input))
		assert.Equal(t, string(expected), string(actual))
	})
}
